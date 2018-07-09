pragma solidity ^0.4.23;
pragma experimental ABIEncoderV2;

import "./RLP.sol";

contract RLPInterface {
	address owner;

	constructor () public {
		owner = msg.sender;
	}

	event rlpItem(bytes item);
	event test(bytes input);

	/*
	* @description	returns the validators array
	*/
	function toRLPBytes(bytes input) public view returns (bytes output) {
		RLP.RLPItem memory item = RLP.toRLPItem(input);
		output = RLP.toBytes(item);
		return;
	}

	function toRLPData(bytes input) public view returns (bytes output) {
		RLP.RLPItem memory item = RLP.toRLPItem(input);
		output = RLP.toData(item);
		return;
	}

	function toRLPAscii(bytes input) public view returns (string output) {
		RLP.RLPItem memory item = RLP.toRLPItem(input);
		output = RLP.toAscii(item);
		return;
	}

}
