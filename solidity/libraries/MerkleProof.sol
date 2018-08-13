pragma solidity ^0.4.23;

library Merkle {
    /**
     * @dev Validate the merkle proof of a specifc leaf with index
     * @param leaf The specific leaf in merkle tree.
     * @param index The index of the leaf in the merkle tree.
     * @param rootHash The root of the merkle tree.
     * @param proof The proof of the leaf.
     */
    function verifyProof(bytes32 leaf, uint256 index, bytes32 rootHash, bytes proof)
        internal
        pure
        returns (bool)
    {
        bytes32 proofElement;
        bytes32 computedHash = leaf;

        for (uint256 i = 32; i <= proof.length; i += 32) {
            assembly {
                proofElement := mload(add(proof, i))
            }
            if (index % 2 == 0) {
                computedHash = keccak256(abi.encodePacked(computedHash, proofElement));
            } else {
                computedHash = keccak256(abi.encodePacked(proofElement, computedHash));
            }
            index = index / 2;
        }
        return computedHash == rootHash;
    }
}
