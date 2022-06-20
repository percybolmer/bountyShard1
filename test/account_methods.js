// https://docs.harmony.one/home/developers/api/methods/account-methods
contract("account_methods", (accounts) => {
  it("balanceOf Address", async () => {
  
    console.log(accounts);

    const balance = await web3.eth.getBalance(accounts[0]);
    // const balance = await counterInstance.getCount.call(accounts[0]);

    // assert.equal(balance.valueOf(), 100, "100 wasn't in the first account");
  });

});