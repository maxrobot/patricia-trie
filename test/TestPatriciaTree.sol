pragma solidity ^0.4.23;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/PatriciaTree.sol";
import {D} from "../contracts/Data.sol";
import {Utils} from "../contracts/Utils.sol";

contract TestPatriciaTree {

    // function testSingleInsert() public {
    //     PatriciaTree patricia = PatriciaTree(DeployedAddresses.PatriciaTree());

    //     patricia.insert("A", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa");

    //     bytes32 expected = 0xd23786fb4a010da3ce639d66d5e904a11dbc02746d1ce25029e53290cabf28ab;

    //     Assert.equal(patricia.root(), expected, "should equal expected");
    // }

    // function testMultipleInsert() public {
    //     PatriciaTree patricia = PatriciaTree(DeployedAddresses.PatriciaTree());

    //     patricia.insert("doe", "reindeer");
    //     patricia.insert("dog", "puppy");
    //     patricia.insert("dogglesworth", "cat");

    //     bytes32 expected = 0x8aad789dff2f538bca5d8ea56e8abe10f4c7ba3a5dea95fea4cd6e7c3a1168d3;

    //     Assert.equal(patricia.root(), expected, "should equal expected");
    // }

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
