#!/bin/sh
set -eu
ARGS=${@:-'./...'}
go test ${ARGS}
exit 0
