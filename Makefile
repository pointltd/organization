NETWORK = point

create-network:
	@docker network ls | grep -w $(NETWORK) || docker network create $(NETWORK)

up:
	@docker-compose up -d