root=$(shell pwd)

all:
	make restart || make init

init:
	# 使用后台启动supervisor来管理多个进程
	supervisord -c supervisor/supervisord.conf

restart:
	# 重启所有进程，需要注意进程间的依赖关系，rpc需要先于api启动
	supervisorctl -c supervisor/supervisord.conf restart userrpc && \
    supervisorctl -c supervisor/supervisord.conf restart orderapi

shutdown:
	# 停止supervisor
	supervisorctl -c supervisor/supervisord.conf shutdown