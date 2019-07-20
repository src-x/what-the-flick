#!/bin/bash

protoc recommendpb/recommend.proto --go_out=plugins=grpc:.