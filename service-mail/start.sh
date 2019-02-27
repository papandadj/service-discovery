#!/bin/bash

file="./pid"
if [ -f "$file" ]
then 
    echo "Service running."
else
    node server.js &
    pid=$!
    echo $pid > pid
    echo "Service start."
fi

