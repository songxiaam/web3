const {expect} = require("chai");
const hre = require("hardhat");
describe("Shipping", function(){
  let shippingContract;
  before(async()=>{
    shippingContract = await hre.ethers.deployContract("Shipping", []);
   
  });

  it("should return the status Pending", async function(){
    expect(await shippingContract.Status()).to.equal("Pending");
  });
  it("Should return the status Shipping", async()=>{
    await shippingContract.shipped();
    expect(await shippingContract.Status()).to.equal("Shipped");
  })
  it("should return correct event description", async () => { 
    await expect( shippingContract.delivered())
    .to.emit(shippingContract, "LogNewAlert") // 验证事件的参数是否符合预期 
    .withArgs("Delivered"); 
    }); 
})
