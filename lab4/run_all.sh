#!/bin/bash

args=`find dataset -type f | xargs`

echo "Serial"
time bash go/serial/run.sh $args

echo "Concurrent"
time bash go/concurrent/run.sh $args
