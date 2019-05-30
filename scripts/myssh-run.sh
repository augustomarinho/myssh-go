#!/bin/bash

BASEDIR=${BASEDIR:-`pwd`}
ROOT_DIR=${ROOT_DIR:-`dirname $BASEDIR`}
SRC_DIR=${SRC_DIR:-"$ROOT_DIR/src/internal/app"}
BIN_DIR=${BIN_DIR:-"$ROOT_DIR/bin"}

GO=${GO:-`which go`}

cd $SRC_DIR

echo "Building project"
$GO install -a

echo "Running"
cd $BIN_DIR
./app