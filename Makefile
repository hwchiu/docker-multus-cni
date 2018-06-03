all: build-docker

multus-cni-bin:
	git submodule init
	git submodule update
	cd multus-cni; git fetch; git checkotu v2.0; ./build; cd ..

build-docker: multus-cni-bin
	sudo docker build -t hwchiu/docker-multus-cni:latest .
	sudo docker push hwchiu/docker-multus-cni
