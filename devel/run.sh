#!/bin/sh
set -eu
docker run -it --rm --network none --name jrmsdev-gojc --hostname gojc -u gojc \
	-v ${PWD}:/go/src/jc jrmsdev/gojc
exit 0
