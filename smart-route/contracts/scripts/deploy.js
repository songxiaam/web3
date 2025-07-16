const hre = require("hardhat");

async function main() {
  console.log("开始部署 Smart Route 可升级合约...");

  // 获取部署账户
  const [deployer] = await ethers.getSigners();
  console.log("部署账户:", deployer.address);

  // 1. 部署实现合约
  console.log("部署实现合约...");
  const SmartRoute = await ethers.getContractFactory("SmartRoute");
  const smartRouteImpl = await SmartRoute.deploy();
  await smartRouteImpl.waitForDeployment();
  console.log("实现合约已部署到:", await smartRouteImpl.getAddress());

  // 2. 部署代理管理员合约
  console.log("部署代理管理员合约...");
  const SmartRouteProxyAdmin = await ethers.getContractFactory("SmartRouteProxyAdmin");
  const proxyAdmin = await SmartRouteProxyAdmin.deploy();
  await proxyAdmin.waitForDeployment();
  console.log("代理管理员合约已部署到:", await proxyAdmin.getAddress());

  // 3. 部署代理合约
  console.log("部署代理合约...");
  const SmartRouteProxy = await ethers.getContractFactory("SmartRouteProxy");
  
  // 准备初始化数据
  const initData = smartRouteImpl.interface.encodeFunctionData("initialize", [
    deployer.address, // owner
    5 // protocolFee (0.05%)
  ]);
  
  const proxy = await SmartRouteProxy.deploy(
    await smartRouteImpl.getAddress(), // logic
    await proxyAdmin.getAddress(), // admin
    initData // data
  );
  await proxy.waitForDeployment();
  console.log("代理合约已部署到:", await proxy.getAddress());

  // 4. 验证合约
  if (hre.network.name !== "hardhat" && hre.network.name !== "localhost") {
    console.log("等待区块确认...");
    await proxy.deploymentTransaction().wait(6);
    
    console.log("验证合约...");
    try {
      await hre.run("verify:verify", {
        address: await smartRouteImpl.getAddress(),
        constructorArguments: [],
      });
      console.log("实现合约验证成功!");
      
      await hre.run("verify:verify", {
        address: await proxyAdmin.getAddress(),
        constructorArguments: [],
      });
      console.log("代理管理员合约验证成功!");
      
      await hre.run("verify:verify", {
        address: await proxy.getAddress(),
        constructorArguments: [
          await smartRouteImpl.getAddress(),
          await proxyAdmin.getAddress(),
          initData
        ],
      });
      console.log("代理合约验证成功!");
    } catch (error) {
      console.log("合约验证失败:", error.message);
    }
  }

  console.log("\n=== 部署完成 ===");
  console.log("实现合约:", await smartRouteImpl.getAddress());
  console.log("代理管理员合约:", await proxyAdmin.getAddress());
  console.log("代理合约:", await proxy.getAddress());
  console.log("使用代理合约地址与合约交互");
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  }); 