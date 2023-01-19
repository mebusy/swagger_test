#!/bin/sh

set -e

PACKAGE=api

DIST_DIR=server/${PACKAGE}
mkdir -p $DIST_DIR

OPENAPI_FILE=../openapi.yaml

# for i in [types gorilla spec] 
for item in types gorilla spec
do
    oapi-codegen -package ${PACKAGE} -generate $item  ${OPENAPI_FILE} > $DIST_DIR/$item.gen.go
done







