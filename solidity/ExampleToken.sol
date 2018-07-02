pragma solidity ^0.4.23;

import "./libraries/openzeppelin-solidity/contracts/token/ERC20/MintableToken.sol";

/**
 * @title ExampleToken
 */
contract ExampleToken is MintableToken {

    string public constant name = "Example Token";

    string public constant symbol = "SMT";

    uint32 public constant decimals = 8;

    constructor() public {

    }
}
