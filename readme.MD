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
* hmy cli


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
sudo apt install libgmp-dev  libssl-dev  make gcc g++
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

The first time you will need Harmony Core installed, perform following, or use `make installHarmony` followed by exports

```bash
mkdir -p $(go env GOPATH)/src/github.com/harmony-one
cd $(go env GOPATH)/src/github.com/harmony-one
git clone https://github.com/harmony-one/mcl.git
git clone https://github.com/harmony-one/bls.git
git clone https://github.com/harmony-one/harmony.git
cd harmony
go mod tidy
make
export CGO_CFLAGS="-I$GOPATH/src/github.com/harmony-one/bls/include -I$GOPATH/src/github.com/harmony-one/mcl/include -I/usr/local/opt/openssl/include"
export CGO_LDFLAGS="-L$GOPATH/src/github.com/harmony-one/bls/lib -L/usr/local/opt/openssl/lib"
export LD_LIBRARY_PATH=$GOPATH/src/github.com/harmony-one/bls/lib:$GOPATH/src/github.com/harmony-one/mcl/lib:/usr/local/opt/openssl/lib
export LIBRARY_PATH=$LD_LIBRARY_PATH
export DYLD_FALLBACK_LIBRARY_PATH=$LD_LIBRARY_PATH
```

If you have trouble installing harmony CLI try 
https://docs.harmony.one/home/network/validators/node-setup/hmy-cli-download

and 

```bash
curl -LO https://harmony.one/hmycli && mv hmycli hmy && chmod +x hmy
./hmy
```

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
Set it to = `dev` to use `.env-dev`

1. Mainnet = [https://api.s1.t.hmny.io, https://api.s0.t.hmny.io] - `.env-prod`
2. Testnet = [https://api.s0.b.hmny.io/, https://api.s0.pops.one/] - `.env-test`
3. Devnet = [https://api.s0.ps.hmny.io/] - `.env-dev`

Environment var to decide network 
`RPCTESTER_ENVIRONMENT`

Value to use
1. Mainnet = "production"
2. Testnet = "test"
3. Devnet = "dev"
4. Localhost = ""

## Compiling solidity contracts
This requires Solc installed

```bash
solc --abi contracts/DevToken.sol -o build
solc --bin contracts/DevToken.sol -o build
abigen --abi=./build/DevToken.abi --bin=./build/DevToken.bin --pkg=devtoken --out=rpctester/contracts/devtoken/DevToken.go

```

## Running tests
Make sure you have configured according to the `Setup of Environments variables` chapter.  
Also you will need BLS keys generated for signing Staking Transactions.


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

## Benchmarking / Stressing
You can use the command `stress` to stresstest endpoints.  
It will send concurrent requests to the endpoint measureing failures and duration to calculate an average.

There are two parameters to give it 
`-c` - Concurrent requests to send, increase to further stress the Endpoint
`-r` - Total amount of requests to send

Example execution
to send 100 concurrent requests and 10000 requests to each endpoint
```bash
cd rpctester
go build -o rpctester
./rpctester stress -c=100 -r=10000 
```

## Benchmarking Results
Results are outputted to a JSON file containg information about each endpoint.

```json
{
    "addressUsed": "0xfc61332CC9CdFDfCba598131595C96F8C33CE209",
    "network": "https://api.s0.ps.hmny.io/",
    "Methods": {
        "hmy_blockNumber": {
            "Average": "585.761232ms",
            "Responses": 10,
            "Failures": 0
        },
    }

```

## Connecting ganach cli to networks
```bash
ganache-cli -f http://localhost:9500 --networkId 1666700000
ganache-cli -f https://api.s0.b.hmny.io --networkId 1666700000
ganache-cli -f https://api.s1.t.hmny.io --networkId 1666600000
```


## Known Issues

### Bad data types
Some endpoints fails due to given the wrong data types, etc documentation says Number but it returns String.
### Hard coded ChainID which does not match NetworkID given
When querying for ChainID it ```1666900000``` for devnet.
But the Github Code for Harmony is hardcoded to...4?
https://github.com/harmony-one/go-sdk/blob/master/pkg/common/chain-id.go#L48  

`error(*errors.errorString) *{s: "blockchain chain id:4, given 1666900000: invalid chain id for signer"}`
When sending RawStakingTransactions I am using chainID 4 for that reason...