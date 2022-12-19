
db-run:
	docker run --rm -d -p 5432:5432 --name postgres -e POSTGRES_PASSWORD=password postgres:15.1
	docker run --rm -d -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=password mysql:8.0.31

db-stop:
	docker stop postgres
	docker stop mysql
