start:
	ganache-cli -f http://localhost:9500 --networkId 1666700000
	
compileContracts:
	solc --abi contracts/Counter.sol -o build
	solc --bin contracts/Counter.sol -o build
	abigen --abi=./build/Counter.abi --bin=./build/Counter.bin --pkg=counter --out=rpctester/contracts/counter/Counter.go
	
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



	