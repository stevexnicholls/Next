#!/bin/bash

set -e

#rm -rf swagger.json
#json-refs resolve --filter relative ./spec/swagger.yml >> swagger.json
rm -rf restapi/
rm -rf models/
swagger generate server -A next -f spec/swagger.yaml --model-package=models --template-dir=./templates --exclude-main

#./scripts/copy_post_code_gen.sh

rm -rf client/
swagger generate client -f spec/swagger.yaml -A next
