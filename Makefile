.PHONY: vendor
vendor:
	godep save ./...

.PHONY: server
server:
	export CODIES_DIR=$$(pwd); cd ./docker;docker-compose up -d

.PHONY: log
log:
	cd ./docker;docker-compose logs -f codies-server

clean_containers:
	docker rm $$(docker stop $$(docker ps -q -a))

connect_mysql:
	mysql -h 127.0.0.1 -u root -pcodies-pwd codies