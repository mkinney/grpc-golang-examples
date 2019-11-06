#!/bin/bash

protoc -I src/ --go_out=plugins=grpc:src/ greeting/greetingpb/greeting.proto
protoc -I src/ --go_out=plugins=grpc:src/ calc/calcpb/calc.proto
