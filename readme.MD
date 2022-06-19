# Testing for Harmony RPC1
This repository is started to start solving  
https://github.com/harmony-one/bounties/issues/117

In this setup we will use Ganache Harmony and Truffle.


## Requirements and dependencies
You can install all requirements using `make installRequirements`

If you want to do it manually 

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
# Fetch ganache fork for harmony
wget https://github.com/harmony-one/harmony-one-ganache-support/releases/download/ganache-harmony-one-2.6.0-beta.3/ganache-2.6.0-beta.3-linux-x86_64.AppImage
sudo chmod +x ganache-2.6.0-beta.3-linux-x86_64.AppImage
mv ganache-2.6.0-beta.3-linux-x86_64.AppImage ganache
# Git clone and build the docker image needed by Ganache
git clone https://github.com/harmony-one/harmony-one-ganache-support.git
./harmony-one-ganache-support/scripts/build-docker.sh
rm -rf ./harmony-one-ganache-support # Clean up
```

After installation, you will have a binary named `ganache`  
you can use to start up a localnet on Ganache.  

## Setup of Environments variables
In the repository you can find a file named `example.env`, in it 
there is placeholders `<YOUR_PRIVATE_KEY>` and `<YOUR_MNEMONIC_CODE>`.  

Copy the file into `.env` and fill in the needed keys  

When using Ganache, I recommend starting it with the -m flag  
to always get the same private keys etc so its easier for you to use  

```bash
./ganache
```

## Run tests against certain network
```bash
truffle test --network localnet
```

## Compile
```bash
truffle compile 
truffle migrate  --network local --reset
```