#!/bin/sh

set -e

docker run --rm -v `pwd`:/local swaggerapi/swagger-codegen-cli-v3 generate \
    -i /local/openapi.yaml \
    -l go-server \
    -o /local/server


