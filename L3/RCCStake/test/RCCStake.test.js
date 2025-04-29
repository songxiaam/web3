const { expect } = require("chai");
const { ethers } = require("hardhat");
const { ZeroAddress, parseEther } = require("ethers");
const { anyValue } = require("@nomicfoundation/hardhat-chai-matchers/withArgs");


describe("RCCStake", function () {
  let RCCStake, rccStake, TestToken, testToken, owner, user1, user2, currentBlock;

  beforeEach(async function () {
    [owner, user1, user2] = await ethers.getSigners();
    RCCStake = await ethers.getContractFactory("RCCStake");
    rccStake = await RCCStake.deploy();
    await rccStake.waitForDeployment();
    // const rccStakeAddress = await rccStake.getAddress();
    // console.log("rccStakeAddress:", rccStakeAddress);

    // 再加一个代币pool
    TestToken = await ethers.getContractFactory("TestToken");
    testToken = await TestToken.deploy("Test Token", "TST", 18);
    await testToken.waitForDeployment();
    console.log("RCCStake deployed at:", rccStake.target);

    // 初始化
    currentBlock = await ethers.provider.getBlockNumber();
    const startBlock = currentBlock + 5;
    const endBlock = startBlock + 500;
    const RCCPerBlock = parseEther("1");
    await rccStake.initialize(testToken.target, startBlock, endBlock, RCCPerBlock);

    // 授权owner拥有Admin权限
    await rccStake.grantRole(await rccStake.ADMIN_ROLE(), owner.address);
  });

  describe("RCCStake Admin Function", function () {
    it("create new pool", async function () {
      await expect(
        rccStake.connect(owner).addPool(
          ZeroAddress,
          1,
          parseEther("0.1"),
          10,
          true
        )).to.emit(rccStake, "AddPool")
        .withArgs(0, ZeroAddress, 1, anyValue, parseEther("0.1"), 10);//测试参数
      // 检查pool内容
      const pool = await rccStake.pools(0);
      const l = await rccStake.pools;
      expect(pool.stTokenAddress).to.equal(ZeroAddress);
      expect(pool.poolWeight).to.equal(1);
      expect(pool.minDepositAmount).to.equal(ethers.parseEther("0.1"));
      expect(pool.unstakeLockedBlocks).to.equal(10);
    });
  })

  describe("RCCStake User Function", function () {

    beforeEach(async function () {
      console.log("RCCStake User Function before each");
      // ⭐ 统一：每次测试前创建一个 pool（假设是原生币）
      await rccStake.connect(owner).addPool(
        ZeroAddress,
        4,
        ethers.parseEther("0.1"),
        10,
        true
      );
      await rccStake.setRCC(testToken.target);
      await rccStake.resumeWithdraw();
      await rccStake.connect(owner).addPool(
        testToken.target,
        2,
        ethers.parseEther("0.5"),
        20,
        true
      );
      console.log("pools.length:", await rccStake.poolLength());
      // await testToken.connect(user1).mint(user1.getAddress(), parseEther("10")); // 给用户代币
      await testToken.transfer(user1.getAddress(), parseEther("10")); // 给用户代币
      await testToken.connect(user1).approve(rccStake.target, parseEther("10")); // 批准给合约

    });

    it("质押原生币", async function () {
      await rccStake.connect(user1).depositNativeCurrency({ value: parseEther("1.5") });
      console.log(user1.getAddress())
      // 检查质押后余额
      const user = await rccStake.users(0, user1.getAddress())
      expect(user.stAmount).to.equal(parseEther("1.5"));
    });
    it("质押代币", async function () {
      await rccStake.connect(user1).deposit(1, parseEther("2"));
      // 检查质押后余额
      const user = await rccStake.users(1, user1.getAddress())
      expect(user.stAmount).to.equal(parseEther("2"));
    });
    it("解质押代币", async function () {
      await rccStake.connect(user1).deposit(1, parseEther("2"));
      // 检查质押后余额
      var user = await rccStake.users(1, user1.getAddress())
      expect(user.stAmount).to.equal(parseEther("2"));
      await rccStake.connect(user1).unstake(1, parseEther("1.5"));
      // 检查解质押后余额
      user = await rccStake.users(1, user1.getAddress());
      expect(user.stAmount).to.equal(parseEther("0.5"));
      var unstakeRequest = await rccStake.userUnstakeRequest(1, user1.getAddress(), 0);
      expect(unstakeRequest.amount).to.equal(parseEther("1.5"));
    });

    it("代币提现", async function () {

      await rccStake.connect(user1).deposit(1, parseEther("2"));
      // 检查质押后余额
      var user = await rccStake.users(1, user1.getAddress())
      expect(user.stAmount).to.equal(parseEther("2"));

      // 质押
      var tmpStartBlock = await ethers.provider.getBlockNumber();
      console.log("tmpStartBlock:", tmpStartBlock);
      await network.provider.send("hardhat_mine", ["0x21"]); // 推进 16 个区块
      var pool = await rccStake.pools(1);
      console.log("pool.accRCCPerST:", pool.accRCCPerST);
      console.log("pool.lastRewardBlock:", pool.lastRewardBlock); 
      console.log("pool.stTokenAmount:", pool.stTokenAmount);
      // 刷新pool
      await rccStake.updatePool(1);
      pool = await rccStake.pools(1);
      // 校验奖励
      console.log("pool.accRCCPerST:", pool.accRCCPerST);
      var tmpEndBlock = await ethers.provider.getBlockNumber();
      console.log("tmpEndBlock:", tmpEndBlock);
      var RCCReward = (tmpEndBlock-tmpStartBlock);
      console.log("奖励1:", parseEther(RCCReward.toString()));
      var reward = await rccStake.getRewardAmount(tmpStartBlock, tmpEndBlock)
      console.log("奖励2:",reward)

      console.log("pool.poolWeight:", pool.poolWeight)
      console.log("pool.totalPoolWeight", await rccStake.getTotalPoolWeight());

      // uint256 RCCReward = (34 * 2) / 6/2;
      // pool.accRCCPerST += (RCCReward * 1e18) / stSupply;
      // console.log("accRCCPerST:", RCCReward*1e18/6)
      // expect(pool.accRCCPerST).to.equal(RCCReward*1e18/6);

      //解质押
      await rccStake.connect(user1).unstake(1, parseEther("1.5"));
      // 
      console.log(await ethers.provider.getBlockNumber());
      // 检查解质押后余额
      user = await rccStake.users(1, user1.getAddress());
      expect(user.stAmount).to.equal(parseEther("0.5"));
      var unstakeRequest = await rccStake.userUnstakeRequest(1, user1.getAddress(), 0);
      pool = await rccStake.pools(1);
      expect(unstakeRequest.amount).to.equal(parseEther("1.5"));
      console.log(unstakeRequest.unlockBlocks)
      //区块推进
      await network.provider.send("hardhat_mine", ["0x20"]); // 推进 16 个区块
      console.log(await ethers.provider.getBlockNumber());

      //提现
      await expect(rccStake.connect(user1).withdraw(1)).to.emit(rccStake, "Withdraw").withArgs(user1.getAddress(), 1, parseEther("1.5"), anyValue);
      expect(pool.stTokenAmount).to.equal(parseEther("0.5"));
      var unstakeRequestLength = await rccStake.userUnstakeRequestLength(1, user1.getAddress());
      expect(unstakeRequestLength).to.equal(0);

      // 提取奖励
      console.log(pool.accRccPerST)
      await rccStake.resumeClaim()
      await expect(rccStake.connect(user1).claim(1)).to.emit(rccStake, "Claim").withArgs(user1.getAddress(), 1, anyValue);

    })

  })
});