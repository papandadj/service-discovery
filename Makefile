-PHONY: build, clean, run, down

build:
	docker build --tag service-balance:1.0.0 --target service-balance .
	docker build --tag service-registry:1.0.0 --target service-registry .
	docker build --tag service-userlist:1.0.0 --target service-userlist .
	docker build --tag service-mail:1.0.0 --target service-mail .
	docker build --tag service-etcd:1.0.0 --target service-etcd .
	docker build --tag service-mysql:1.0.0 --target service-mysql .

run:
	docker-compose up

down:
	docker-compose down

cleanimage: 
	docker rmi -f service-balance:1.0.0
	docker rmi -f service-registry:1.0.0
	docker rmi -f service-userlist:1.0.0
	docker rmi -f service-mail:1.0.0 
	docker rmi -f service-etcd:1.0.0
	docker rmi -f service-mysql:1.0.0

