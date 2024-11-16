NETWORK = point
APP_NAME = organization

YC_CR_DOMAIN = cr.yandex/crpbvv0uke53s3ief4mk

GIT_BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)
GIT_TAG ?= $(shell git rev-parse --short=10 HEAD)

create-network:
	@docker network ls | grep -w $(NETWORK) || docker network create $(NETWORK)

#docker
dkr-up:
	@docker compose --file docker-compose.local.yaml up

dkr-down:
	@docker compose --file docker-compose.local.yaml down

dkr-build:
	@docker compose --file docker-compose.local.yaml build

dkr-push:
	@docker tag organization-app cr.yandex/crp4640u3tckkugq0upa/organization-app:latest
	@docker push cr.yandex/crp4640u3tckkugq0upa/organization-app:latest

yc-auth:
	@export YC_TOKEN=$$(yc iam create-token)