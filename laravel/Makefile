up:
	cd laradock && docker-compose up -d nginx mysql workspace
build:
	cd laradock && docker-compose up --build -d nginx mysql workspace
remake:
	@make destroy
	@make up
stop:
	cd laradock && docker-compose stop
down:
	cd laradock && docker-compose down
destroy:
	cd laradock && docker-compose down --rmi all --volumes
ps:
	cd laradock && docker-compose ps
logs:
	cd laradock && docker-compose logs
logs-watch:
	cd laradock && docker-compose logs --follow
app:
	cd laradock && docker-compose exec workspace bash
