pragma solidity ^0.4.21;

contract Donate {
    uint public amountRaised;
    address private owner;

    uint[] public timeArray;

    constructor() public{
      owner = msg.sender;
    }

    function () payable public {
      amountRaised += msg.value;
      timeArray.push(now);
    }

    function transfer(uint256 _value) public returns (bool success) {
        require(msg.sender == owner);
        require(address(this).balance >= _value);
        owner.transfer(_value);
        amountRaised -= _value;
	    
        return true;
    }

}
