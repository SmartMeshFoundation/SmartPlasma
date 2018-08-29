pragma solidity ^0.4.23;

library PlasmaLib {
    function generateUID(
        address account,
        address currency,
        uint depositCount
    )
        internal
        pure
        returns(bytes32)
    {
        return bytes32(
            keccak256(
                abi.encodePacked(
                    currency,
                    account,
                    depositCount
                )
            )
        );
    }
}
