#!/bin/sh
set -eu
ARGS=${@}
if test "X" = "X${ARGS}"; then
	ARGS='-bench=.'
fi
for fn in $(ls ./internal/_bench/*.go); do
	echo "--- ${fn}"
	go test ${ARGS} ${fn}
done
exit 0
