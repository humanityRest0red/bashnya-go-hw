COVERAGE_FILE := coverage.out

TEST_CMD := go test work

test:
	$(TEST_CMD)

coverage:
	- $(TEST_CMD) -coverprofile=$(COVERAGE_FILE)
	go tool cover -html=$(COVERAGE_FILE)

clean:
	rm -rf $(COVERAGE_FILE)