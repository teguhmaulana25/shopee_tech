# shopee_tech
Backend Technical Assessment on Shopee ID

### Getting Started
1. Clone This Repository
2. Put your first dump db in ``init.db`` folder, the extension should be ``.sql``
3. Open and edit ``.env`` value

### Running and import database
1. Run ``make init`` to create container and run mysql bash
2. Run ``mysql -u ${MYSQL_USER} -p -D ${MYSQL_DATABASE} < /docker-entrypoint-initdb.d/{your-database-dump-name}``

### Get MySQL IP
1. Default IP is 172.8.0.4
2. Run ``sudo docker inspect ${MYSQL_CONTAINER} | grep IPAdd``
3. Use the IP to connect from another container or application

### Get phpMyAdmin IP
1. Default IP is 172.8.0.5
2. Run ``sudo docker inspect ${PHPMYADMIN_CONTAINER} | grep IPAdd``
3. Run the IP in your browser

### Get Golang IP
1. Default IP is http://localhost:8083