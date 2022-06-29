

start:
	ganache-cli -f http://localhost:9500 --networkId 1666700000
	
compileContracts:
	solc --abi contracts/DevToken.sol -o build --overwrite
	solc --bin contracts/DevToken.sol -o build --overwrite
	abigen --abi=./build/DevToken.abi --bin=./build/DevToken.bin --pkg=devtoken --out=rpctester/contracts/devtoken/DevToken.go
	
installRequirements:
	sudo npm install -g ganache-cli
	wget https://github.com/harmony-one/harmony-one-ganache-support/releases/download/ganache-harmony-one-2.6.0-beta.3/ganache-2.6.0-beta.3-linux-x86_64.AppImage
	sudo chmod +x ganache-2.6.0-beta.3-linux-x86_64.AppImage
	mv ganache-2.6.0-beta.3-linux-x86_64.AppImage ganache
	git clone https://github.com/harmony-one/harmony-one-ganache-support.git
	./harmony-one-ganache-support/scripts/build-docker.sh
	rm -rf ./harmony-one-ganache-support
	wget https://go.dev/dl/go1.18.3.linux-amd64.tar.gz
	rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.3.linux-amd64.tar.gz
	export PATH=$PATH:/usr/local/go/bin
	sudo add-apt-repository -y ppa:ethereum/ethereum
	sudo apt-get install -y ethereum 
	sudo apt-get upgrade -y geth 
	sudo apt-get install solc
	cd rpctester && go get -d github.com/ethereum/go-ethereum/...
	sudo apt install libgmp-dev  libssl-dev  make gcc g++

installHarmony:
	mkdir -p $(shell go env GOPATH)/src/github.com/harmony-one
	cd $(shell go env GOPATH)/src/github.com/harmony-one && \
	git clone https://github.com/harmony-one/mcl.git && \
	git clone https://github.com/harmony-one/bls.git && \
	git clone https://github.com/harmony-one/harmony.git && \
	cd harmony && go mod tidy && make

linuxInstallHarmony:
	curl -LO https://harmony.one/hmycli && mv hmycli hmy && chmod +x hmy
