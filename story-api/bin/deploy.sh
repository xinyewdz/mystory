#!/bin/bash
pid=`ps -ef|grep storyapi_server |grep -v grep | awk '{printf $2}'`
if [ -n $pid ]; then
 kill $pid
fi
nohup go run storyapi_server.go &
