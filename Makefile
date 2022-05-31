start_postgres : 
	cd postgresql/scripts/ && docker compose up

start_local:
	go run main.go