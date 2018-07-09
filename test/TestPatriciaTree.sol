pragma solidity ^0.4.23;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/PatriciaTree.sol";
import {D} from "../contracts/Data.sol";
import {Utils} from "../contracts/Utils.sol";

// contract TestPatriciaTree is Patricia {
contract TestPatriciaTree {

    // PatriciaTree patricia;
    

    function testInsert() public {
        PatriciaTree patricia = PatriciaTree(DeployedAddresses.PatriciaTree());
        
        patricia.insert("val", "VAL");
        Assert.equal(patricia.root(), "TPAIN", "should equal expected");
    }

    // function test() {
    //     //testInsert();
    //     testProofs();
    // }
    // function testInsert() internal {
    //     insert("one", "ONE");
    //     insert("two", "ONE");
    //     insert("three", "ONE");
    //     insert("four", "ONE");
    //     insert("five", "ONE");
    //     insert("six", "ONE");
    //     insert("seven", "ONE");
    //     // update
    //     insert("one", "TWO");
    // }
    // function testProofs() internal {
    //     insert("one", "ONE");
    //     var (branchMask, siblings) = getProof("one");
    //     verifyProof(root, "one", "ONE", branchMask, siblings);
    //     insert("two", "TWO");
    //     (branchMask, siblings) = getProof("one");
    //     verifyProof(root, "one", "ONE", branchMask, siblings);
    //     (branchMask, siblings) = getProof("two");
    //     verifyProof(root, "two", "TWO", branchMask, siblings);
    // }
}
