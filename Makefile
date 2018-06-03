all: build-docker

multus-cni-bin:
	git submodule init
	git submodule update
	cd multus-cni; ./build; cd ..

build-docker: multus-cni-bin
	sudo docker build -t hwchiu/docker-multus-cni:latest .
	sudo docker push hwchiu/docker-multus-cni
