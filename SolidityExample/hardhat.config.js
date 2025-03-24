require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: {
    version: "0.8.28",  // 推荐 ≥0.8.24
    settings: {
      evmVersion: "cancun"  // 强制指定 Cancun
    }
  },
  networks: {
    hardhat: {
      hardfork: "cancun"  // 本地节点启用 Cancun
    }
  }
};
