default: start

help:
	@echo ""
	@echo "clean\t\tremove all artifacts and dependencies"
	@echo "build\t\tbuild docker image"
	@echo "start\t\tstart local development setup"
	@echo "stop\t\tstop local development setup"

build: clean-dist
	@echo "\nBuilding Docker Image:\n"
	docker build -t casino-royale:latest .

start:
	@echo "\nStarting dev environment:\n"
	@docker-sync-stack start

stop: clean

clean: clean-dev clean-dist clean-deps clean-test

clean-dist:
	@rm -rf tmp/

clean-deps:
	@rm -rf vendor/

clean-dev:
	@docker-sync-stack clean

clean-test:
	@docker-compose -f docker-compose-test.yml rm -f

test: clean-test
	@docker-compose -f docker-compose-test.yml build
	@docker-compose -f docker-compose-test.yml up --abort-on-container-exit --exit-code-from server-test

# .PHONY: default help build start stop clean clean-dist clean-deps clean-dev test
.PHONY: default help build start stop clean clean-dist clean-deps clean-dev
