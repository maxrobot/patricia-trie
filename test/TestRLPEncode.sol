pragma solidity ^0.4.23;

import "truffle/Assert.sol";  
import { RLPEncode } from "../libraries/RLPEncode.sol";

contract TestRLPEncode {  

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

}