// https://docs.harmony.one/home/developers/api/methods/account-methods

const util = require("util")
const mlog = require('mocha-logger');
// sendRpc is a help util for making web3.currentProvider.send a async without callback needed
const sendRpc = util.promisify((web3.currentProvider).send)
    .bind(web3.currentProvider);


contract("account_methods_v2", (accounts) => {
  it("getBalance_v2 no Address Param", async () => {
      var start = Date.now();
      const response = await sendRpc({
        method: "hmyv2_getBalance",
        params: [],
        jsonrpc: "2.0",
        id: new Date().getTime()
      });
      var end = Date.now();
      mlog.log(`hmyv2_getBalance: ${end - start} ms`);
      assert.equal(response.error.code, -32602, "Should return -32602 if no param is set");
  });
  it("getBalance_v2", async () => {
    var start = Date.now();
    const response = await sendRpc({
      method: "hmyv2_getBalance",
      params: [ accounts[0]],
      jsonrpc: "2.0",
      id: new Date().getTime()
    });
    var end = Date.now();
    mlog.log(`hmyv2_getBalance: ${end - start} ms`);
    assert.equal(response.error, undefined, "Should have undefined error");
    assert.typeOf(response.result, 'number');

    //balance = (response.result).toFixed(20) / 1e18;
  });
  it("getBalanceByBlock_v2", async () => {
    var start = Date.now();
    const response = await sendRpc({
      method: "hmyv2_getBalanceByBlockNumber",
      params: [ accounts[0], 1],
      jsonrpc: "2.0",
      id: new Date().getTime()
    });
    var end = Date.now();
    mlog.log(`hmyv2_getBalance: ${end - start} ms`);
    assert.equal(response.error, undefined, "Should have undefined error");
    assert.typeOf(response.result, 'number');
    //balance = (response.result).toFixed(20) / 1e18;
  });
  
});

contract("account_methods_v1", (accounts) => {
  it("getBalance_v1 no Address Param", async () => {
      const response = await sendRpc({
        method: "hmy_getBalance",
        params: [],
        jsonrpc: "2.0",
        id: new Date().getTime()
      });
      assert.equal(response.error.code, -32602, "Should return -32602 if no param is set");
  });
  it("getBalance_v1 no block", async () => {
    const response = await sendRpc({
      method: "hmy_getBalance",
      params: [ accounts[0]],
      jsonrpc: "2.0",
      id: new Date().getTime()
    });
    assert.equal(response.error.code, -32602, "Should return -32602 if no param is set");
  });
  it("getBalance_v1 with Block 1", async () => {
    const response = await sendRpc({
      method: "hmy_getBalance",
      params: [ accounts[0], "0x1"],
      jsonrpc: "2.0",
      id: new Date().getTime()
    });
    assert.equal(response.error, undefined, "Should have undefined error");
    assert.typeOf(response.result, 'string');
  });
});