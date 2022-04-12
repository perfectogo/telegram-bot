#!/bin/bash

CURRENT_DIR=$(pwd)

#mkdir $CURRENT_DIR/genproto
#mkdir $CURRENT_DIR/genproto/catalog

protoc -I /usr/local/include \
       -I $GOPATH/src/github.com/gogo/protobuf/gogoproto \
       -I $CURRENT_DIR/protos/ \
        --gofast_out=plugins=grpc:$CURRENT_DIR/genproto/ \
        $CURRENT_DIR/protos/message.proto;

if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i "" -e "s/,omitempty//g" $CURRENT_DIR/genproto/*.go
  else
    sed -i -e "s/,omitempty//g" $CURRENT_DIR/genproto/*.go
fi