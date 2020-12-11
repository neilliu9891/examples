#!/bin/bash

echo $1

for((i=1;i<=$1;i++));
    do
       echo "Welcome $i times"
       ./subscriber $2 $3 &
 done
