pragma solidity ^0.4.23;

import "truffle/Assert.sol";  
import { RLPEncode } from "../libraries/RLPEncode.sol";

contract TestRLPEncode {  

    function testEncodeEmpty() {
        // Input and expected output
        bytes memory input = hex'';
        bytes memory expected = hex'80';
        uint expectedLength = 1;

        // RLP Encode the ting...
        bytes memory output = RLPEncode.encodeItem(input);

        // Assert
        Assert.equal(output.length, expected.length, "RLP Encoded Length");
        Assert.equal(output.length, expectedLength, "RLP Encoded Length just to be sure");
        for (uint256 index = 0; index < expected.length; index++) {
            Assert.equal(output[index], expected[index], "The RLP Encoded values");            
        }
    }

    function testEncodeIntegersItem() {
        // Input and expected output
        bytes memory input = hex'FFFFFF';
        bytes memory expected = hex'83FFFFFF';
        uint expectedLength = 4;

        // RLP Encode the ting...
        bytes memory output = RLPEncode.encodeItem(input);

        // Assert
        Assert.equal(output.length, expected.length, "RLP Encoded Length");
        Assert.equal(output.length, expectedLength, "RLP Encoded Length just to be sure");
        for (uint256 index = 0; index < expected.length; index++) {
            Assert.equal(output[index], expected[index], "The RLP Encoded values");            
        }
    }


    function testEncodeStringItem() {
        // Input and expected output
        bytes memory input = 'dog';
        bytes memory expected = hex'83646F67';
        uint expectedLength = 4;

        // RLP Encode the ting...
        bytes memory output = RLPEncode.encodeItem(input);

        // Assert
        Assert.equal(output.length, expected.length, "RLP Encoded Length");
        Assert.equal(output.length, expectedLength, "RLP Encoded Length just to be sure");
        for (uint256 index = 0; index < expected.length; index++) {
            Assert.equal(output[index], expected[index], "The RLP Encoded values");            
        }
    }

    function testEncodeLongStringItem() {
        // Input and expected output
        bytes memory input = 'Lorem ipsum dolor sit amet, consectetur adipisicing eli';
        bytes memory expected = hex'B74C6F72656D20697073756D20646F6C6F722073697420616D65742C20636F6E7365637465747572206164697069736963696E6720656C69';
        uint expectedLength = 56;

        // RLP Encode the ting...
        bytes memory output = RLPEncode.encodeItem(input);

        // Assert
        Assert.equal(output.length, expected.length, "RLP Encoded Length");
        Assert.equal(output.length, expectedLength, "RLP Encoded Length just to be sure");
        for (uint256 index = 0; index < expected.length; index++) {
            Assert.equal(output[index], expected[index], "The RLP Encoded values");            
        }
    }

    function testEncodeIntegersList() {
        // Input and expected output
        bytes[] memory input = new bytes[](3);
        input[0] = hex'01';
        input[1] = hex'02';
        input[2] = hex'03';

        bytes memory expected = hex'C3010203';
        uint expectedLength = 4;

        // RLP Encode the ting...
        bytes memory output = RLPEncode.encodeList(input);

        // Assert
        Assert.equal(output.length, expected.length, "RLP Encoded Length");
        Assert.equal(output.length, expectedLength, "RLP Encoded Length just to be sure");
        for (uint256 index = 0; index < expected.length; index++) {
            Assert.equal(output[index], expected[index], "The RLP Encoded values");            
        }
    }


    function testEncodeStringList() {
        // Input and expected output
        bytes[] memory input = new bytes[](15);
        // Very ugly but I don't care, suggestions please...
        input[0] = "aaa";
        input[1] = "bbb";
        input[2] = "ccc";
        input[3] = "ddd";
        input[4] = "eee";
        input[5] = "fff";
        input[6] = "ggg";
        input[7] = "hhh";
        input[8] = "iii";
        input[9] = "jjj";
        input[10] = "kkk";
        input[11] = "lll";
        input[12] = "mmm";
        input[13] = "nnn";
        input[14] = "ooo";

        bytes memory expected = hex'F83C836161618362626283636363836464648365656583666666836767678368686883696969836A6A6A836B6B6B836C6C6C836D6D6D836E6E6E836F6F6F';
        uint expectedLength = 62;

        // RLP Encode the ting...
        bytes memory output = RLPEncode.encodeList(input);

        // Assert
        Assert.equal(output.length, expected.length, "RLP Encoded Length");
        Assert.equal(output.length, expectedLength, "RLP Encoded Length just to be sure");
        for (uint256 index = 0; index < expected.length; index++) {
            Assert.equal(output[index], expected[index], "The RLP Encoded values");            
        }
    }

}