#!/bin/bash
env=$1
if [ -n "$env" ]
then
  if [ $env=="prod" ]
  then
    mv conf/app.conf.prod conf/app.conf
  fi
fi;

export GOPROXY=http://goproxy.cn
GOROOT=/usr/local/go
server_name=storyapi_server
while [ true ];do
  pid=`ps -ef|grep $server_name|grep -v grep|awk '{print $2}'`
  if [ -n "$pid" ];then
    echo "kill $pid"
    kill $pid
    sleep 1
  else
    echo "$server_name already stop"
    break;
  fi
done
nohup $GOROOT/bin/go run storyapi_server.go &
echo "$server_name start success"