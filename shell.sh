#!/bin/sh
set -e

(
  cd "$(dirname "$0")"
  go build -o /tmp/myshell main.go
)

exec /tmp/myshell "$@"
