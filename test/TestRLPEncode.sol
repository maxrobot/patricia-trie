pragma solidity ^0.4.23;

import "truffle/Assert.sol";  
import { RLPEncode } from "../libraries/RLPEncode.sol";

contract TestRLPEncode {  

    function testEncodeEmpty() {
        // Input and expected output
        bytes memory input = hex'';
        bytes memory expected = hex'80';
        uint expectedLength = 1;

        // Act
        bytes memory output;
        output = RLPEncode.encodeItem(input);

        // Assert
        Assert.equal(output.length, expected.length, "RLP Encoded Length");
        Assert.equal(output.length, expectedLength, "RLP Encoded Length just to be sure");
        for (var index = 0; index < expected.length; index++) {
            Assert.equal(output[index], expected[index], "The RLP Encoded values");            
        }
    }

    function testEncodeIntegersItem() {
        // Input and expected output
        bytes memory input = hex'FFFFFF';
        bytes memory expected = hex'83FFFFFF';
        uint expectedLength = 4;

        // Act
        bytes memory output;
        output = RLPEncode.encodeItem(input);

        // Assert
        Assert.equal(output.length, expected.length, "RLP Encoded Length");
        Assert.equal(output.length, expectedLength, "RLP Encoded Length just to be sure");
        for (var index = 0; index < expected.length; index++) {
            Assert.equal(output[index], expected[index], "The RLP Encoded values");            
        }
    }


    function testEncodeStringItem() {
        // Input and expected output
        bytes memory input = 'dog';
        bytes memory expected = hex'83646F67';
        uint expectedLength = 4;

        // Act
        bytes memory output;
        output = RLPEncode.encodeItem(input);

        // Assert
        Assert.equal(output.length, expected.length, "RLP Encoded Length");
        Assert.equal(output.length, expectedLength, "RLP Encoded Length just to be sure");
        for (var index = 0; index < expected.length; index++) {
            Assert.equal(output[index], expected[index], "The RLP Encoded values");            
        }
    }

    function testEncodeLongStringItem() {
        // Input and expected output
        bytes memory input = 'Lorem ipsum dolor sit amet, consectetur adipisicing eli';
        bytes memory expected = hex'B74C6F72656D20697073756D20646F6C6F722073697420616D65742C20636F6E7365637465747572206164697069736963696E6720656C69';
        uint expectedLength = 56;

        // Act
        bytes memory output;
        output = RLPEncode.encodeItem(input);

        // Assert
        Assert.equal(output.length, expected.length, "RLP Encoded Length");
        Assert.equal(output.length, expectedLength, "RLP Encoded Length just to be sure");
        for (var index = 0; index < expected.length; index++) {
            Assert.equal(output[index], expected[index], "The RLP Encoded values");            
        }
    }

}