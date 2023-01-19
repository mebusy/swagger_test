#!/bin/sh

set -e

PACKAGE=api

DIST_DIR=server/${PACKAGE}
mkdir -p $DIST_DIR

OPENAPI_FILE=../openapi.yaml

for item in types spec
do
    oapi-codegen -package ${PACKAGE} -generate $item  ${OPENAPI_FILE} > $DIST_DIR/$item.gen.go
done

# srv code
oapi-codegen --config=srv.cfg.yaml ${OPENAPI_FILE}







