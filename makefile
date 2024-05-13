dockerbuild:
	docker compose --env-file .env build

dockerup:
	docker compose --env-file .env up