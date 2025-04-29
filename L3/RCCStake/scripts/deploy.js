const hre = require("hardhat");

async function main() {
  const RCCStake = await hre.ethers.getContractFactory("RCCStake");
  const rccStake = await RCCStake.deploy();

  await rccStake.deployed();
  console.log("RCCStake deployed to:", rccStake.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
