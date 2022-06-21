start:
	ganache-cli -f http://localhost:9500 --networkId 1666700000
	

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



	