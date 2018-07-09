const PatriciaTree = artifacts.require("PatriciaTree");
const RLPInterface = artifacts.require("RLPInterface");

module.exports = async (deployer) => {
  try {
    deployer.deploy(PatriciaTree)
      .then(() => PatriciaTree.deployed)
      .then(() => deployer.deploy(RLPInterface))
      .then(() => RLPInterface.deployed)
  } catch(err) {
    console.log('ERROR on deploy:',err);
  }

};
