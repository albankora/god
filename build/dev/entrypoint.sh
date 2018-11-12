#!/usr/bin/env bash
set -e

GO_WORK_DIR=${GO_WORK_DIR:-$GOPATH/src}
cd ${GO_WORK_DIR}
exec "$@"