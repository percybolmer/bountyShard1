# Testing for Harmony RPC1
This repository is started to start solving  
https://github.com/harmony-one/bounties/issues/117

In this setup we will use Ganache Harmony and Truffle.


## Requirements and dependencies
You will need Truffle installed
```bash
sudo apt install nodejs npm # If you dont have nodejs and npm
npm install -g truffle
npm install --save @harmony-js/core
npm install --save @harmony-js/utils
npm install --save tslib
npm install --save dotenv
```

You will also need a local Harmony node to test against when developing the tests before running them in Production RPC
https://github.com/harmony-one/harmony-one-ganache-support to find out how to use Ganache on harmony.
Releases found here: https://github.com/harmony-one/harmony-one-ganache-support/releases

```bash
npm install -g ganache-cli

```


## Setup of Environments variables
In the repository you can find a file named `example.env`, in it 
there is placeholders `<YOUR_PRIVATE_KEY>` and `<YOUR_MNEMONIC_CODE>`.  

