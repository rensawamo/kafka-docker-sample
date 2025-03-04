

up: 
	@docker-compose up

down: 
	@docker-compose down

publisher:
	cd cmd/publisher && go run .

consumer:
	cd cmd/consumer && go run .