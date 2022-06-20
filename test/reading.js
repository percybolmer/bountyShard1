const Counter = artifacts.require("Counter");

contract("Counter", (accounts) => {
  it("balanceOf Address", async () => {
    const counterInstance = await Counter.deployed();
  
    console.log(accounts);

    const balance = await web3.eth.getBalance(accounts[0]);

    
    // const balance = await counterInstance.getCount.call(accounts[0]);

    // assert.equal(balance.valueOf(), 100, "100 wasn't in the first account");
  });

});