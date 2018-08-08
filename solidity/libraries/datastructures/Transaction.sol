pragma solidity ^0.4.23;

import "../RLP.sol";
import "../ECRecovery.sol";

library Transaction {
    using RLP for bytes;
    using RLP for bytes[];
    using RLP for RLP.RLPItem;
    using RLP for RLP.Iterator;

    uint constant TRANSACTION_LENGTH = 6;

    struct Tx {
        uint prevBlock;
        uint uid;
        uint amount;
        address newOwner;
        uint nonce;
        address signer;
        bytes32 hash;
    }

    function createTx(bytes memory txBytes)
        internal
        constant
        returns (Tx memory)
    {
        RLP.RLPItem[] memory txList = txBytes.toRLPItem().toList(TRANSACTION_LENGTH);
        return Tx({
            prevBlock: txList[0].toUint(),
            uid: txList[1].toUint(),
            amount: txList[2].toUint(),
            newOwner: txList[3].toAddress(),
            nonce: txList[4].toUint(),
            signer: _getSigner(txList),
            hash: _txHash(txList)
        });
    }

    function _txHash(RLP.RLPItem[] memory txList)
        private
        view
        returns (bytes32)
    {
        bytes[] memory unsignedTxList = new bytes[](5);
        for (uint i = 0; i < 5; i++) {
            unsignedTxList[i] = txList[i].toBytes();
        }
        bytes memory unsignedTx = unsignedTxList.encodeList();
        return keccak256(unsignedTx);
    }

    function _getSigner(RLP.RLPItem[] memory txList)
        private
        view
        returns (address)
    {
        bytes[] memory unsignedTxList = new bytes[](5);
        for (uint i = 0; i < 5; i++) {
            unsignedTxList[i] = txList[i].toBytes();
        }
        bytes memory unsignedTx = unsignedTxList.encodeList();
        bytes memory sig = txList[5].toData();
        return ECRecovery.recover(keccak256(unsignedTx), sig);
    }
}
