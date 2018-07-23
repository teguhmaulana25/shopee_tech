init: create-network create-container run-it
create-network:
	sudo docker network prune
	sudo docker network create -d bridge --subnet=172.8.0.0/16 --gateway=172.8.0.1 mysql_bridge
create-container:
	sudo docker-compose up -d
run-it:
	sudo docker exec -it mysql-dev /bin/bash
