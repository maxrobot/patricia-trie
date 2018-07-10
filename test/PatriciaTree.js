// // const Util = require('ethereumjs-util');
// const Web3 = require('web3');
// // const Web3Utils = require('web3-utils');
// // const Web3Abi = require('web3-eth-abi');
// // const Web3Accounts = require('web3-eth-accounts');
// // const rlp = require('rlp');

// const Patricia = artifacts.require("PatriciaTree");

// const web3 = new Web3();

// web3.setProvider(new web3.providers.HttpProvider('http://localhost:8501'));

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

// contract('patriciaTree.js', (accounts) => {
//   const joinHex = arr => '0x' + arr.map(el => el.slice(2)).join('');

//   const watchEvent = (eventObj) => new Promise((resolve,reject) => eventObj.watch((error,event) => error ? reject(error) : resolve(event)));

//   const blockNum = 1;

//   it('Test: Insert()', async () => {
//     const patricia = await Patricia.new();
//     const accounts = web3.eth.accounts;

//     let insertReceipt = await patricia.insert("do", "verb");
//     const recoveredKey = insertReceipt.logs[0].args['key'];
//     const recoveredValue = insertReceipt.logs[1].args['value'];
    
//     console.log(recoveredKey)
//     console.log(recoveredValue)
//     console.log("Key: ", insertReceipt.logs[2].args['data'])
//     console.log(insertReceipt.logs[2].args['length'])
//     console.log(insertReceipt.logs[3].args['trie'])
//     console.log("Root: ",insertReceipt.logs[4].args['root'])
//     console.log(insertReceipt.logs[5].args['node'])
//     console.log(insertReceipt.logs[5].args['length'])
//     console.log(insertReceipt.logs[5].args['data'])
//     console.log("Test Root: ", insertReceipt.logs[6].args['root'])

//     // Get the root
//     const rootReceipt = await patricia.root()
//     const bytesRoot = hexToBytes(rootReceipt)
    
//     const expectedRoot = "0x014f07ed95e2e028804d915e0dbd4ed451e394e1acfd29e463c11a060b2ddef7"
//     console.log("Received root: ", rootReceipt)
//     console.log("Expected root: ", expectedRoot)
//     // assert.equal(expectedRoot, bytesRoot)
//   })

// });