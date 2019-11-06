#!/bin/bash

protoc -I src/ --go_out=plugins=grpc:src/ greeting/greetingpb/greeting.proto
