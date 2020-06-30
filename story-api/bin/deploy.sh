#!/bin/bash
pid=`ps -ef|grep storyapi_server.bin |grep -v grep | awk '{printf $2}'`
if [ -n $pid ]; then
 kill $Pid
fi
chmod +x storyapi_server.bin
nohup ./storyapi_server.bin &
