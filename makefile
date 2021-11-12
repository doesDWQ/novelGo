root=$(shell pwd)

dev: userRpc orderApi

orderApi:
	cd $(root)/order/api/ \
	&& goctl api go -api order.api -dir . \
	&& nohup go run order.go -f etc/order.yaml > $(root)/logs/orderApi.log &

userRpc:
	cd $(root)/user/rpc \
	&& goctl rpc proto -src user.proto -dir . \
	&& go run user.go -f etc/user.yaml > $(root)/logs/userRpc.log &