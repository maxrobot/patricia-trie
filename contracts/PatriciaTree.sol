pragma solidity ^0.4.23;
pragma experimental ABIEncoderV2;

import {D} from "./Data.sol";
import {Utils} from "./Utils.sol";
import "./RLP.sol";

contract PatriciaTree {
    // Mapping of hash of key to value
    mapping (bytes32 => bytes) values;

    // Particia tree nodes (hash to decoded contents)
    mapping (bytes32 => D.Node) nodes;
    // The current root hash, keccak256(node(path_M('')), path_M(''))
    bytes32 public root;
    address public owner;
    D.Edge rootEdge;

    event showkey(bytes key);
    event showvalue(bytes value);
    event showEmptyTrie(uint256 trie);
    event showroot(bytes32 root);
    event showLabel(bytes32 data, uint256 length);
    event hashData(bytes32 node, uint256 length, bytes32 data);
    // event test(RLP.RLPItem item);

    function PatriciaTree() {
        owner = msg.sender;
    }

    function getNode(bytes32 hash) constant returns (uint, bytes32, bytes32, uint, bytes32, bytes32) {
        var n = nodes[hash];
        return (
            n.children[0].label.length, n.children[0].label.data, n.children[0].node,
            n.children[1].label.length, n.children[1].label.data, n.children[1].node
        );
    }

    function getRootEdge() constant returns (uint, bytes32, bytes32) {
        return (rootEdge.label.length, rootEdge.label.data, rootEdge.node);
    }
    
    function edgeHash(D.Edge e) internal returns (bytes32) {
        return keccak256(e.node, e.label.length, e.label.data);
    }
    
    function testHash(D.Edge e) public returns (bytes32) {
        emit hashData(e.node, e.label.length, e.label.data);

        // RLP.RLPItem memory item = RLP.toRLPItem("0x00");

        // RLP.toList(item);

        // return keccak256(item);
        return keccak256(0xc98320646f8476657262);
        // return keccak256(e.node, e.label.data);
    }
    
    // Returns the hash of the encoding of a node.
    function hash(D.Node memory n) internal returns (bytes32) {
        return keccak256(edgeHash(n.children[0]), edgeHash(n.children[1]));
    }
    
    // Returns the Merkle-proof for the given key
    // Proof format should be:
    //  - uint branchMask - bitmask with high bits at the positions in the key
    //                    where we have branch nodes (bit in key denotes direction)
    //  - bytes32[] hashes - hashes of sibling edges
    function getProof(bytes key) constant returns (uint branchMask, bytes32[] _siblings) {
        D.Label memory k = D.Label(keccak256(key), 256);
        D.Edge memory e = rootEdge;
        bytes32[256] memory siblings;
        uint length;
        uint numSiblings;
        while (true) {
            var (prefix, suffix) = Utils.splitCommonPrefix(k, e.label);
            require(prefix.length == e.label.length);
            if (suffix.length == 0) {
                // Found it
                break;
            }
            length += prefix.length;
            branchMask |= uint(1) << (255 - length);
            length += 1;
            var (head, tail) = Utils.chopFirstBit(suffix);
            siblings[numSiblings++] = edgeHash(nodes[e.node].children[1 - head]);
            e = nodes[e.node].children[head];
            k = tail;
        }
        if (numSiblings > 0)
        {
            _siblings = new bytes32[](numSiblings);
            for (uint i = 0; i < numSiblings; i++)
                _siblings[i] = siblings[i];
        }
    }

    function verifyProof(bytes32 rootHash, bytes key, bytes value, uint branchMask, bytes32[] siblings) constant {
        D.Label memory k = D.Label(keccak256(key), 256);
        D.Edge memory e;
        e.node = keccak256(value);
        for (uint i = 0; branchMask != 0; i++) {
            uint bitSet = Utils.lowestBitSet(branchMask);
            branchMask &= ~(uint(1) << bitSet);
            (k, e.label) = Utils.splitAt(k, 255 - bitSet);
            uint bit;
            (bit, e.label) = Utils.chopFirstBit(e.label);
            bytes32[2] memory edgeHashes;
            edgeHashes[bit] = edgeHash(e);
            edgeHashes[1 - bit] = siblings[siblings.length - i - 1];
            e.node = keccak256(edgeHashes);
        }
        e.label = k;
        require(rootHash == edgeHash(e));
    }
    
    // TODO also return the proof
    function insert(bytes key, bytes value) {
        emit showkey(key);
        emit showvalue(value);
        D.Label memory k = D.Label(keccak256(key), 256);
        emit showLabel(k.data, k.length);
        bytes32 valueHash = keccak256(value);
        values[valueHash] = value;
        // keys.push(key);
        D.Edge memory e;
        if (rootEdge.node == 0 && rootEdge.label.length == 0)
        {
            emit showEmptyTrie(0);
            // Empty Trie
            e.label = k;
            e.node = valueHash;
        }
        else
        {
            emit showEmptyTrie(1);
            e = insertAtEdge(rootEdge, k, valueHash);
        }
        root = edgeHash(e);
        emit showroot(root);
        emit showroot(testHash(e));
        rootEdge = e;
    }
    
    function insertAtNode(bytes32 nodeHash, D.Label key, bytes32 value) internal returns (bytes32) {
        require(key.length > 1);
        D.Node memory n = nodes[nodeHash];
        var (head, tail) = Utils.chopFirstBit(key);
        n.children[head] = insertAtEdge(n.children[head], tail, value);
        return replaceNode(nodeHash, n);
    }
    
    function insertAtEdge(D.Edge e, D.Label key, bytes32 value) internal returns (D.Edge) {
        require(key.length >= e.label.length);
        var (prefix, suffix) = Utils.splitCommonPrefix(key, e.label);
        bytes32 newNodeHash;
        if (suffix.length == 0) {
            // Full match with the key, update operation
            newNodeHash = value;
        } else if (prefix.length >= e.label.length) {
            // Partial match, just follow the path
            newNodeHash = insertAtNode(e.node, suffix, value);
        } else {
            // Mismatch, so let us create a new branch node.
            var (head, tail) = Utils.chopFirstBit(suffix);
            D.Node memory branchNode;
            branchNode.children[head] = D.Edge(value, tail);
            branchNode.children[1 - head] = D.Edge(e.node, Utils.removePrefix(e.label, prefix.length + 1));
            newNodeHash = insertNode(branchNode);
        }
        return D.Edge(newNodeHash, prefix);
    }
    function insertNode(D.Node memory n) internal returns (bytes32 newHash) {
        bytes32 h = hash(n);
        nodes[h].children[0] = n.children[0];
        nodes[h].children[1] = n.children[1];
        return h;
    }
    function replaceNode(bytes32 oldHash, D.Node memory n) internal returns (bytes32 newHash) {
        delete nodes[oldHash];
        return insertNode(n);
    }
}
