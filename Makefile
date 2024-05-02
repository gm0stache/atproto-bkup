
.PHONY: test-integration
test-integration:
	docker compose -f ./test-integration/docker-compose.yaml up --build