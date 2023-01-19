#!/bin/sh

set -e

DIST=server

# clean old files
rm -rf $DIST/go
rm -rf $DIST/api
rm -rf $DIST/.swagger-codegen*


docker run --rm -v `pwd`:/local swaggerapi/swagger-codegen-cli-v3 generate \
    -i /local/openapi.yaml \
    -l go-server \
    -o /local/${DIST}


