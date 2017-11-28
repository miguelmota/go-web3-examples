pragma solidity ^0.4.9;

contract Greeter {
  string public name;
  uint256 public count;

  event _Greet(string name, uint256 count);

  function greet(string _name) public {
    name = _name;
    count += 1;
    _Greet(_name, count);
  }
}
