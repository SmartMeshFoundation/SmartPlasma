pragma solidity ^0.4.23;

import "./libraries/openzeppelin-solidity/contracts/token/ERC20/MintableToken.sol";

/**
 * @title ExampleToken
 */
contract ExampleToken is MintableToken {

    string public constant name = "Example Token";

    string public constant symbol = "SMT";

    uint32 public constant decimals = 8;

    bool disableApprove = false;
    bool disableTransferFrom = false;


    modifier checkApprove {
        require(!disableApprove);
        _;
    }

    modifier checkTransferFrom {
        require(!disableTransferFrom);
        _;
    }

    constructor() public {

    }

    // if changeApproveState == true then approve and increaseApproval not work
    function changeApproveState() public {
        if (disableApprove) {
            disableApprove = false;
        } else {
            disableApprove = true;
        }
    }

    // if disableTransferFrom == true then transferFrom not work
    function changeTransferFromState() public {
        if (disableTransferFrom) {
            disableTransferFrom = false;
        } else {
            disableTransferFrom = true;
        }
    }

    function approve(
        address _spender,
        uint256 _value
    ) public checkApprove returns (bool) {
        return super.approve(_spender, _value);
    }

    function increaseApproval(
        address _spender,
        uint _addedValue
    ) public checkApprove returns (bool) {
        return super.increaseApproval(_spender, _addedValue);
    }

    function transferFrom(
        address from, address to,
        uint256 value) public checkTransferFrom returns (bool) {
        return super.transferFrom(from, to, value);
    }
}
