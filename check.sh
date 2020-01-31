#!/bin/sh
set -eu
ARGS=${@:-'./...'}
go vet ${ARGS}
exit 0
