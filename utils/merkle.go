package utils

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"math/big"
	"razor/bindings"
)

type CommitData struct {
	AssignedCollections    map[int]bool
	SeqAllottedCollections []*big.Int
	Leaves                 []*big.Int
}

type Account struct {
	Address  string
	Password string
}

func CreateMerkle(values []*big.Int) ([][][]byte, error) {
	if len(values) == 0 {
		return [][][]byte{}, errors.New("values are nil, cannot create merkle tree")
	}
	var tree [][][]byte
	var leaves [][]byte
	for i := 0; i < len(values); i++ {
		leaves = append(leaves, solsha3.SoliditySHA3([]string{"uint256"}, []interface{}{values[i]}))
	}

	level := leaves
	var nextLevel [][]byte
	tree = append(tree, level)

	for len(level) != 1 {
		for i := 0; i < len(level); i += 2 {
			if i+1 < len(level) {
				nextLevel = append(nextLevel, solsha3.SoliditySHA3([]string{"bytes32", "bytes32"}, []interface{}{level[i], level[i+1]}))
			} else {
				nextLevel = append(nextLevel, level[i])
			}
		}
		level = nextLevel
		tree = append(tree, level)
		nextLevel = nil
	}

	// Reverse the tree
	for i, j := 0, len(tree)-1; i < j; i, j = i+1, j-1 {
		tree[i], tree[j] = tree[j], tree[i]
	}

	return tree, nil
}

func GetProofPath(tree [][][]byte, assetId uint16) [][32]byte {
	var compactProofPath [][32]byte
	for currentLevel := len(tree) - 1; currentLevel > 0; currentLevel-- {
		currentLevelNodes := tree[currentLevel]
		currentLevelCount := len(currentLevelNodes)
		if int(assetId) == currentLevelCount-1 && currentLevelCount%2 == 1 {
			assetId = assetId / 2
			continue
		}
		var node [32]byte
		if assetId%2 == 1 {
			copy(node[:], currentLevelNodes[assetId-1])
		} else {
			copy(node[:], currentLevelNodes[assetId+1])
		}
		compactProofPath = append(compactProofPath, node)
		assetId = assetId / 2
	}
	return compactProofPath
}

func GetMerkleRoot(tree [][][]byte) ([32]byte, error) {
	var root [32]byte
	if tree == nil {
		return root, errors.New("tree is nil")
	}
	copy(root[:], tree[0][0])
	if root == [32]byte{} {
		return root, errors.New("root is nil")
	}
	return root, nil
}

func GenerateTreeRevealData(merkleTree [][][]byte, commitData CommitData) bindings.RevealMerkleTree {
	if merkleTree == nil || commitData.SeqAllottedCollections == nil || commitData.Leaves == nil {
		log.Error("No data present for construction of StructsMerkleTree")
		return bindings.RevealMerkleTree{}
	}
	var (
		values []bindings.RevealAssignedAsset
		proofs [][][32]byte
	)

	for i := 0; i < len(commitData.SeqAllottedCollections); i++ {
		value := bindings.RevealAssignedAsset{
			LeafId: uint16(commitData.SeqAllottedCollections[i].Uint64()),
			Value:  big.NewInt(commitData.Leaves[commitData.SeqAllottedCollections[i].Uint64()].Int64()),
		}
		proof := GetProofPath(merkleTree, value.LeafId)
		values = append(values, value)
		proofs = append(proofs, proof)
	}

	root, err := GetMerkleRoot(merkleTree)
	if err != nil {
		log.Error("Error in getting root: ", err)
		return bindings.RevealMerkleTree{}
	}
	fmt.Sprintf("GenerateTreeRevealData: values = %+v, proofs = %+v, root = %v", values, proofs, root)
	fmt.Println()

	return bindings.RevealMerkleTree{
		Values: values,
		Proofs: proofs,
		Root:   root,
	}
}

func CalculateSecret(account Account, epoch uint32, keystorePath string, chainId *big.Int) ([]byte, []byte, error) {
	if chainId == nil {
		return nil, nil, errors.New("chainId is nil")
	}
	hash := solsha3.SoliditySHA3([]string{"address", "uint32", "uint256", "string"}, []interface{}{common.HexToAddress(account.Address), epoch, chainId, "razororacle"})
	log.Debug("CalculateSecret: Hash: ", hash)
	ethHash := SignHash(hash)
	log.Debug("Hash generated for secret")
	log.Debug("CalculateSecret: Ethereum signed hash: ", ethHash)
	signedData, err := SignData(ethHash, account, keystorePath)
	if err != nil {
		return nil, nil, errors.New("Error in signing the data: " + err.Error())
	}
	log.Debug("CalculateSecret: SignedData: ", signedData)
	fmt.Sprintf("Checking whether recovered address from Hash: %v and Signed data: %v is same as given address...", hash, signedData)
	fmt.Println()
	recoveredAddress, err := EcRecover(hash, signedData)
	if err != nil {
		return nil, nil, errors.New("Error in verifying: " + err.Error())
	}
	log.Debug("CalculateSecret: Recovered Address: ", recoveredAddress)
	if recoveredAddress != common.HexToAddress(account.Address) {
		return nil, nil, errors.New("invalid verification")
	}
	log.Debug("Address verified, generating secret....")
	if signedData[64] == 0 || signedData[64] == 1 {
		signedData[64] += 27
	}

	secret := crypto.Keccak256(signedData)
	log.Debug("Secret generated.")
	return signedData, secret, nil
}

func SignHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}

func SignData(hash []byte, account Account, defaultPath string) ([]byte, error) {
	privateKey, err := GetPrivateKey(account.Address, account.Password, defaultPath)
	if err != nil {
		return nil, err
	}
	return crypto.Sign(hash, privateKey)
}

func EcRecover(data, sig hexutil.Bytes) (common.Address, error) {
	if len(sig) != 65 {
		return common.Address{}, fmt.Errorf("signature must be 65 bytes long")
	}

	rpk, err := crypto.SigToPub(SignHash(data), sig)
	if err != nil {
		return common.Address{}, err
	}
	return crypto.PubkeyToAddress(*rpk), nil
}
