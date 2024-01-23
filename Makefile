postgres:
	docker run --name postgres --network pharmago-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=Hoanglong2502 -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root $(name)

new_migration:
	migrate create -ext sql -dir $(service)/db/migration -seq $(name)

new_service:
	cd ..