// const Util = require('ethereumjs-util');
const Web3 = require('web3');
// const Web3Utils = require('web3-utils');
// const Web3Abi = require('web3-eth-abi');
// const Web3Accounts = require('web3-eth-accounts');
// const rlp = require('rlp');

// const RLPEncode = artifacts.require("RLPEncode");

// function hexToBytes(hex) {
//     for (var bytes = [], c = 0; c < hex.length; c += 2)
//     bytes.push(parseInt(hex.substr(c, 2), 16));
//     return bytes;
// }

// function bytesToHex(bytes) {
//     for (var hex = [], i = 0; i < bytes.length; i++) {
//         hex.push((bytes[i] >>> 4).toString(16));
//         hex.push((bytes[i] & 0xF).toString(16));
//     }
//     return hex.join("");
// }

// contract.only('RLPEncode.js', (accounts) => {
//   const joinHex = arr => '0x' + arr.map(el => el.slice(2)).join('');

//   const watchEvent = (eventObj) => new Promise((resolve,reject) => eventObj.watch((error,event) => error ? reject(error) : resolve(event)));

//   it('Test: encode()', async () => {
//     const rlp = await RLPEncode.new();
//     let input = hexToBytes("0127");

//     console.log(input)

//     const receiptRLP = await rlp.encode(input);
//     // console.log(receiptRLP)
//     console.log(receiptRLP.logs[0].args['length'])
//     console.log(receiptRLP.logs[1].args['input'])
//   })

// });