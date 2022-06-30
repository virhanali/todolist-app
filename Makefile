go:
	go run .
build:
	docker compose up --build

run:
	docker run --network="host" todolist_server