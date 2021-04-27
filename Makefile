.PHONY: clean
clean:
	@rm -rf bin coverage.txt profile.out

.PHONY: lint
lint:
	@golint ./...

.PHONY: report
report:
	@curl -XPOST 'https://goreportcard.com/checks' --data 'repo=github.com/blackhorseya/learn-design-patterns'

.PHONY: test-unit
test-unit:
	@sh $(shell pwd)/scripts/go.test.sh
