#!/bin/sh
set -eu
rm -vrf ./_test
go clean -i -testcache ./...
exit 0
