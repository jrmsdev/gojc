#!/bin/sh
set -eu
ARGS=${@}
if test "X" = "X${ARGS}"; then
	ARGS='.'
fi
for fn in $(ls ./internal/_bench/*.go); do
	echo "--- ${fn}"
	go test -bench=${ARGS} ${fn}
done
exit 0
