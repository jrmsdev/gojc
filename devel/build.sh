#!/bin/sh
set -eu
docker build --rm --network none -t jrmsdev/gojc \
	--build-arg GOJC_UID=$(id -u) \
	--build-arg GOJC_GID=$(id -g) \
	--build-arg GOJC_UMASK=$(umask) \
	./devel
exit 0
