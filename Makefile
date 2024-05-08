
.PHONY: test-integration
test-integration:
	docker compose -f ./test-integration/docker-compose.yaml down --volumes
	docker compose -f ./test-integration/docker-compose.yaml up -d --build
	@-testerExitCode=$$(docker container wait test-integration-tester-1)
	@-docker container logs test-integration-tester-1
	@-docker compose -f ./test-integration/docker-compose.yaml stop
	@exit $(testerExitCode)
