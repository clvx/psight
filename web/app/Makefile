build:
	docker run --name postgres -e POSTGRES_PASSWORD=noadmin -d -p 5432:5432 postgres

populate:
	psql -U postgres -p noadmin -h localhost
	CREATE DATABASE lss;
	\c lss
	create table users( id int primary key not null, email varchar null, password varchar null, firstname varchar null, lastname varchar null, lastlogin date null );
	insert into users values (1, 'test@gmail.com', 'password', 'Luis', 'Ibarra', '2020-04-06');
