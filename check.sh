#!/bin/sh
set -eu
ARGS=${@:-'./...'}
go vet -mod=vendor ${ARGS}
exit 0
