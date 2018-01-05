docker-rpm:
	@sudo docker run -v `pwd`:/test:rw centos:centos${OS_VERSION} /bin/bash /test/build.sh
