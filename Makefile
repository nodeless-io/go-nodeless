.PHONY: test
test:
	go test -v .

.PHONY: test-integration
test-integration:
	go test -v --tags=integration .
