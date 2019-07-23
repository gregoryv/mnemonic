#!/bin/bash

mkdir -p dist
rm -rf dist/*
out=dist/mnemonic
go build -o $out ./cmd/mnemonic
upx $out
cp CHANGELOG.md dist/
# Get the version from the binary
ver=mnemonic-$(./dist/mnemonic -v)
rm -f $ver.tgz
mv dist $ver
tar -c $ver | gzip - > $ver.tgz
rm -rf $ver
tar tvfz $ver.tgz
