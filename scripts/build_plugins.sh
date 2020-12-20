#!/bin/bash

for path in cf-templates/* ; do
    [ -d "${path}" ] || continue

    dirname="$(basename "${path}")"

    go build -buildmode=plugin -o "cf-templates/${dirname}/${dirname}.so" "cf-templates/${dirname}/${dirname}.go"
done