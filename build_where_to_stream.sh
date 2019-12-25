#!/bin/sh

cd pkg/http/web
npm run build
cd ../../..
go build -o where_to_stream cmd/api/main.go