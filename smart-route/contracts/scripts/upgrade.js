const hre = require("hardhat");

async function main() {
  console.log("开始升级 Smart Route 合约...");

  // 获取部署账户
  const [deployer] = await ethers.getSigners();
  console.log("升级账户:", deployer.address);

  // 合约地址（需要根据实际部署的地址修改）
  const PROXY_ADDRESS = "0x..."; // 代理合约地址
  const PROXY_ADMIN_ADDRESS = "0x..."; // 代理管理员地址

  // 1. 部署新的实现合约
  console.log("部署新的实现合约...");
  const SmartRouteV2 = await ethers.getContractFactory("SmartRouteV2");
  const smartRouteV2Impl = await SmartRouteV2.deploy();
  await smartRouteV2Impl.waitForDeployment();
  console.log("新实现合约已部署到:", await smartRouteV2Impl.getAddress());

  // 2. 获取代理管理员合约实例
  const proxyAdmin = await ethers.getContractAt("SmartRouteProxyAdmin", PROXY_ADMIN_ADDRESS);

  // 3. 升级代理合约
  console.log("升级代理合约...");
  const upgradeTx = await proxyAdmin.upgrade(
    PROXY_ADDRESS,
    await smartRouteV2Impl.getAddress()
  );
  await upgradeTx.wait();
  console.log("代理合约已升级!");

  // 4. 验证升级
  const proxy = await ethers.getContractAt("SmartRouteV2", PROXY_ADDRESS);
  console.log("升级后的合约地址:", await proxy.getAddress());

  // 5. 验证合约
  if (hre.network.name !== "hardhat" && hre.network.name !== "localhost") {
    console.log("等待区块确认...");
    await upgradeTx.wait(6);
    
    console.log("验证新实现合约...");
    try {
      await hre.run("verify:verify", {
        address: await smartRouteV2Impl.getAddress(),
        constructorArguments: [],
      });
      console.log("新实现合约验证成功!");
    } catch (error) {
      console.log("合约验证失败:", error.message);
    }
  }

  console.log("\n=== 升级完成 ===");
  console.log("新实现合约:", await smartRouteV2Impl.getAddress());
  console.log("代理合约地址:", PROXY_ADDRESS);
  console.log("合约已成功升级到 V2 版本");
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  }); 