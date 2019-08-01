#!/bin/bash

SERVER="apiserver"
BASE_DIR=$PWD
INTERVAL=2

# 命令行参数，需要手动指定
ARGS=""

function is_exist()
{
	pid=`ps -u ${UID} -ef|grep ${SERVER} | grep -v grep|awk '{print $2}'`
	#如果不存在返回1，存在返回0
	if [ -z "${pid}" ];then
		return 1
	else
		return 0
	fi
}
# 启动
function start()
{
	is_exist
	if [ $? -eq "0" ];then
		echo "\033[34m ${SERVER} is already running. pid=${pid} \033[0m"
		exit 1
	else
		nohup $BASE_DIR/$SERVER $ARGS  &>/dev/null &
		echo "\033[33m ${SERVER} is starting... \033[0m" && sleep ${INTERVAL}
	fi

	#check status
	is_exist
	if [ $? -eq "0" ];then
		echo "\033[32m ${SERVER} start success \033[0m"
		exit 1
	else
		echo "\033[31m ${SERVER} start failed \033[0m"
	fi
}
# 停止
function stop()
{
	is_exist
	if [ $? -eq "0" ];then
		kill -9 $pid
	else
		echo "${SERVER} is not running"
		exit 1
	fi

	echo "\033[33m killing pid  ${pid}... \033[0m" &&  sleep ${INTERVAL}

	is_exist
	if [ $? -eq "0" ];then
		echo "\033[31m ${SERVER} stop failed \033[0m"
		exit 1
	else
		echo "\033[32m  ${SERVER} is stopped \033[0m"
	fi
}

function status()
{
	is_exist
	if [ $? -eq "0" ];then
		echo "\033[32m  ${SERVER} is  running \033[0m"
	else
		echo "\033[31m ${SERVER} is not running \033[0m"

	fi
}



case "$1" in
	'start')
	start
	;;
	'stop')
	stop
	;;
	'status')
	status
	;;
	'restart')
	stop && start
	;;
	*)
	echo "usage: $0 {start|stop|restart|status}"
	exit 1
	;;
esac