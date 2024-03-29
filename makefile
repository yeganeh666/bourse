GOPATH=${HOME}/go

%:
	@true

.PHONY: fmt
fmt:
	./scripts/fmt.sh $(filter-out $@,$(MAKECMDGOALS))

.PHONY: run
run:
	./scripts/fmt.sh $(filter-out $@,$(MAKECMDGOALS))
	./scripts/run.sh $(filter-out $@,$(MAKECMDGOALS))

.PHONY: proto
proto:
	./scripts/proto.sh $(filter-out $@,$(MAKECMDGOALS))

.PHONY: doc
doc:
	./scripts/doc.sh $(filter-out $@,$(MAKECMDGOALS))

.PHONY: docker-service
docker:
	./scripts/docker-service.sh $(filter-out $@,$(MAKECMDGOALS))

.PHONY: cli
cli:
	./scripts/cli.sh $(filter-out $@,$(MAKECMDGOALS))