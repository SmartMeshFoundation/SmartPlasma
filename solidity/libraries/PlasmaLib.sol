pragma solidity ^0.4.18;

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
                currency,
                account,
                depositCount
            )
        );
    }
}
