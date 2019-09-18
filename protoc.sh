#!/bin/bash

protoc -I kv/ kv/kv.proto --go_out=plugins=grpc:kv