up:
	docker-compose up

down:
	docker-compose down
	
up-build:
	docker-compose up --build
	
create-db:
	$(DOCKER_COMPOSE) exec postgres psql -U postgres -c "CREATE DATABASE postgres;"

create-schema:
	$(DOCKER_COMPOSE) exec postgres psql -U postgres -d postgres -f ./app/schema.sql

seeds:
	$(DOCKER_COMPOSE) exec postgres psql -U postgres -d postgres -f ./app/seeds.sql