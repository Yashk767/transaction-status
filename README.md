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
Every set transaction increments `value` by 1.
You will be able to see output like below.  
```
Value Before setting:  24
Sending Set transaction...
Set Txn Hash:  0x74fd8cb34e6759023481b80cb37e7b05efcd317d2d52794d4a42afef077e7e1a
Checking if transaction is mined....
Checking state change, value  24
Checking if transaction is mined....
TIME TO GET TRANSACTION STATUS:  3.15400525s
Transaction mined successfully
Checking state change, value  25
Value has been updated, state has changed
TIME TO GET STATE CHANGE:  3.21752325s
Value After setting:  25

```
Get-Set Contract deployed at `0xC081a517Dff9Ee85acF8029ed2892E448B17b2D2` on SKALE staging network

```
// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract Value {
    uint256 public testValue;

    event ValueChanged(uint256 _newValue);

    function set(uint256 _testValue) public {
        testValue = _testValue;
        emit ValueChanged(testValue);
    }

    function get() public view returns(uint256) {
        return testValue;
    }
}
```


