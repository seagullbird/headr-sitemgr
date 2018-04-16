#!/usr/bin/env sh

protoc sitemgrsvc.proto --go_out=plugins=grpc:.
