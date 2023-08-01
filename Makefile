serve:
	@echo "Starting Docker..."
	docker-compose up -d
	@echo "Docker containers started!"

down:
	@echo "Stopping docker containers..."
	docker-compose down
	@echo "Stopped!"
