const PatriciaTree = artifacts.require("PatriciaTree");

module.exports = async (deployer) => {
  try {
    deployer.deploy(PatriciaTree)
      .then(() => PatriciaTree.deployed)
  } catch(err) {
    console.log('ERROR on deploy:',err);
  }

};
