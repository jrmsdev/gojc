#!/bin/sh
set -eu
ARGS=$@
if test "X" = "X${ARGS}"; then
	ARGS='./...'
fi
go test ${ARGS}
exit 0
