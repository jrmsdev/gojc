#!/bin/sh
set -eu
ARGS=${@}
if test "X" = "X${ARGS}"; then
	ARGS='-bench=.'
	for fn in $(ls ./internal/_bench/*.go); do
		echo "--- ${fn}"
		go test -bench=. ${fn}
	done
else
	go test ${ARGS}
fi
exit 0
