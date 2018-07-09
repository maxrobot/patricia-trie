// const Util = require('ethereumjs-util');
const Web3 = require('web3');
// const Web3Utils = require('web3-utils');
// const Web3Abi = require('web3-eth-abi');
// const Web3Accounts = require('web3-eth-accounts');
// const rlp = require('rlp');

const RLPInterface = artifacts.require("RLPInterface");

function hexToBytes(hex) {
    for (var bytes = [], c = 0; c < hex.length; c += 2)
    bytes.push(parseInt(hex.substr(c, 2), 16));
    return bytes;
}

function bytesToHex(bytes) {
    for (var hex = [], i = 0; i < bytes.length; i++) {
        hex.push((bytes[i] >>> 4).toString(16));
        hex.push((bytes[i] & 0xF).toString(16));
    }
    return hex.join("");
}

contract.only('RLPInterface.js', (accounts) => {
  const joinHex = arr => '0x' + arr.map(el => el.slice(2)).join('');

  const watchEvent = (eventObj) => new Promise((resolve,reject) => eventObj.watch((error,event) => error ? reject(error) : resolve(event)));

  const blockNum = 1;

  it('Test: toRLPBytes()', async () => {
    const rlp = await RLPInterface.new();

    const receiptRLPitem = await rlp.toRLPBytes("0x0127");
    console.log(receiptRLPitem)
  })


  it('Test: toRLPData()', async () => {
    const rlp = await RLPInterface.new();

    const receiptRLPitem = await rlp.toRLPData("0x00");
    console.log(receiptRLPitem)
  })

  it('Test: toRLPAscii()', async () => {
    const rlp = await RLPInterface.new();

    const receiptRLPitem = await rlp.toRLPAscii("val");
    console.log(receiptRLPitem)
  })

});