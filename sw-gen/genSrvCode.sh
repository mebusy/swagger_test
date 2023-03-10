#!/bin/sh

set -e

DIST=server

# clean old files
rm -rf $DIST/go
rm -rf $DIST/api
rm -rf $DIST/.swagger-codegen*

# get last path component of pwd
DIRNAME=`basename $PWD`

docker run --rm -v `pwd`/..:/local swaggerapi/swagger-codegen-cli-v3 generate \
    -i /local/openapi.yaml \
    -DhideGenerationTimestamp=true \
    -l go-server \
    -o /local/${DIRNAME}/${DIST}

# replace relative import paht in the generated main.go
sed -i '' 's/".\/go"/"server\/go"/g' server/main.go
