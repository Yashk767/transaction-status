# Transaction-status
We will be able to see how much time is required to get the transaction status moved from pending state to status as `1`. 
We also log a time required for state change

### Steps to run
**_NOTE:_** Golang stable version should be installed
1. Clone the repository
2. Run command
```
go run main.go
```

Reveal_Mock Contract deployed at `0x13696984Ba47f0a0257D379a5b622D0670628e14` on SKALE staging network

```
// SPDX-License-Identifier: MIT

pragma solidity ^0.8.4;

import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.7.3/contracts/utils/cryptography/ECDSA.sol";

contract Reveal {

    struct AssignedAsset {
        uint16 leafId;
        uint256 value;
    }

    struct MerkleTree {
        AssignedAsset[] values;
        bytes32[][] proofs;
        bytes32 root;
    }

    uint256 public toAssign = 3;
    uint16 public numActiveCollections = 1; 
    uint256 public depth = 0;

    function reveal(
        uint32 epoch,
        MerkleTree memory tree,
        bytes memory signature
    ) external {
        require(tree.values.length == toAssign, "values length mismatch");

        bytes32 messageHash = keccak256(abi.encodePacked(msg.sender, epoch, block.chainid, "razororacle"));
        require(ECDSA.recover(ECDSA.toEthSignedMessageHash(messageHash), signature) == msg.sender, "invalid signature");
        bytes32 secret = keccak256(signature);

        //bytes32 seed = keccak256(abi.encode(salt, secret));
        // require(keccak256(abi.encode(tree.root, seed)) == commitmentHash, "incorrect secret/value");
        //below line also avoid double reveal attack since once revealed, commitment has will be set to 0x0

        uint16 max = numActiveCollections;
        for (uint16 i = 0; i < tree.values.length; i++) {
        //    require(_isAssetAllotedToStaker(seed, i, max, tree.values[i].leafId), "Revealed asset not alloted");
            require(tree.values[i].value != 0, "0 vote for assigned coll");
            // reason to ignore : its internal lib not a external call
            // slither-disable-next-line calls-loop
            require(
                MerklePosAware.verify(
                    tree.proofs[i],
                    tree.root,
                    keccak256(abi.encode(tree.values[i].value)),
                    tree.values[i].leafId,
                    depth,
                    numActiveCollections
                ),
                "invalid merkle proof"
            );
        }
    }

    function _isAssetAllotedToStaker(
        bytes32 seed,
        uint16 iterationOfLoop,
        uint16 max,
        uint16 leafId
    ) internal view returns (bool) {
        // max= numAssets, prng_seed = seed+ iteration of for loop
        if (_prng(keccak256(abi.encode(seed, iterationOfLoop)), max) == leafId) return true;
        return false;
    }

    /**
     * @dev an internal function used by _isAssetAllotedToStaker to check for allocation
     * @param prngSeed hash of seed and exact position in sequence
     * @param max total number of active collections
     */
    function _prng(bytes32 prngSeed, uint256 max) internal pure returns (uint256) {
        uint256 sum = uint256(prngSeed);
        return (sum % max);
    }
}

/**
 * @dev These functions deal with verification of Merkle trees (hash trees),
 */
library MerklePosAware {
    function verifyMultiple(
        bytes32[][] memory proofs,
        bytes32 root,
        bytes32[] memory leaves,
        uint16[] memory leafId,
        uint256 depth,
        uint16 maxAssets
    ) internal pure returns (bool) {
        for (uint256 i = 0; i < proofs.length; i++) {
            if (!verify(proofs[i], root, leaves[i], leafId[i], depth, maxAssets)) return false;
        }
        return true;
    }

    /**
     * @dev Returns true if a `leaf` can be proved to be a part of a Merkle tree
     * defined by `root`. For this, a `proof` must be provided, containing
     * sibling hashes on the branch from the leaf to the root of the tree. Each
     * pair of leaves and each pair of pre-images are assumed to be sorted.
     */
    function verify(
        bytes32[] memory proof,
        bytes32 root,
        bytes32 leaf,
        uint16 leafId,
        uint256 depth,
        uint16 maxAssets
    ) internal pure returns (bool) {
        bytes32 computedHash = leaf;
        bytes memory seq = bytes(getSequence(leafId, depth));

        uint256 lastNode = maxAssets;
        uint256 myNode = leafId + 1;
        uint256 j = depth;
        uint256 i = 0;
        while (j > 0) {
            bytes32 proofElement = proof[i];
            j--;
            //skip proof check  if my node is  last node and number of nodes on level is odd
            if (lastNode % 2 == 1 && lastNode == myNode) {
                myNode = myNode / 2 + (myNode % 2); // (myNode % 2) always equal to 1
                lastNode = lastNode / 2 + (lastNode % 2); // (lastNode % 2) always equal to 1
                continue;
            }
            // 0x30 is 0, 0x31 is 1
            if (seq[j] == 0x30) {
                computedHash = keccak256(abi.encodePacked(computedHash, proofElement));
            } else {
                computedHash = keccak256(abi.encodePacked(proofElement, computedHash));
            }
            i++;

            myNode = myNode / 2 + (myNode % 2);
            lastNode = lastNode / 2 + (lastNode % 2);
        }

        return computedHash == root;
    }

    function getSequence(uint256 leafId, uint256 depth) internal pure returns (bytes memory) {
        bytes memory output = new bytes(depth);
        for (uint8 i = 0; i < depth; i++) {
            output[depth - 1 - i] = (leafId % 2 == 1) ? bytes1("1") : bytes1("0");
            leafId /= 2;
        }
        return output;
    }
}
```


