# Testing for Harmony RPC1
This repository is started to start solving  
https://github.com/harmony-one/bounties/issues/117

In this setup we will use Ganache Harmony for hosting an localnet, and we use Go (geth) to compile the contracts and 
run the tests.


## Requirements and dependencies
There are three major requirements

* Ganache Harmony
* Go
* Geth 
* Solc


You can install all requirements using `make installRequirements`

If you want to do it manually 

To run the tests you will need Go installed
```bash
wget https://go.dev/dl/go1.18.3.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.3.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

You also need Geth installed. You can find installation instructions on their website.

```bash
sudo add-apt-repository -y ppa:ethereum/ethereum
sudo apt-get install ethereum
sudo apt-get upgrade geth
sudo apt-get install solc
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

You can use `make start` to start this localnet

## Setup of Environments variables
In the repository you can find a file named `rpctester/example.env`, in it 
there is a few environment variables needed. 

It is prefilled with the Localnet defaults for genesis account. 

You can replace with whatever network and private key and mneonic you want,

Copy the file into `.env` and fill in the needed keys  


GanacheHarmony has a bug where Transactions to fund the wallets in ganache hangs.
But the Generis wallet has funds  
`0xA5241513DA9F4463F1d4874b548dFBAC29D91f34`  

The private key for this wallet is 
`1f84c95ac16e6a50f08d44c7bde7aff8742212fda6e4321fde48bf83bef266dc`

If you want to use a mainnet, you can create a `.env-prod` file and insert the real values
and private key. 

Set the environment variable `RPCTESTER_ENVIRONMENT` = "production" to use the `.env-prod`

## Compiling solidity contracts
This requires Solc installed

```bash
solc --abi contracts/Counter.sol -o build
solc --bin contracts/Counter.sol -o build
abigen --abi=./build/Counter.abi --bin=./build/Counter.bin --pkg=counter --out=rpctester/contracts/counter/Counter.go

```

## Running tests
To run the tests you will use the std Go tooling.

Make sure you have started the localnet or configure the network you want using the flags explained below.
```bash
make start
```

```bash
go test 
```

By default it will look for a `.env` file in the rpctester folder
If you want another Network or Address you can insert flags as following 

```bash
go test --address="your adr" --url="networkurl"
```


## Test results
Test results are printed in a `results.json` file inside the rpctester folder.



## Connecting ganach cli to networks
```bash
ganache-cli -f http://localhost:9500 --networkId 1666700000
ganache-cli -f https://api.s0.b.hmny.io --networkId 1666700000
ganache-cli -f https://api.s1.t.hmny.io --networkId 1666600000
```
