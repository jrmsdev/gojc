#!/bin/sh
set -eu
ARGS=$@
if test "X" = "X${ARGS}"; then
	ARGS='./...'
fi
mkdir -p ./_test
./test.sh -coverprofile ./_test/coverage.out ${ARGS}
go tool cover -html ./_test/coverage.out -o ./_test/coverage.html
exit 0
