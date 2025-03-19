GO := go
PKGS := $(shell $(GO) list ./...)
TEST_FLAGS := -v
CONFIG_TEST_FILE := testconfig.json
INPUT_TEST_FILE := testinput.txt
INPUT_INVALID_TEST_FILE := testinvalid.txt

test: 
	$(GO) test $(TEST_FLAGS) ./...

clean:
	rm -f $(CONFIG_TEST_FILE) $(INPUT_TEST_FILE) $(INPUT_INVALID_TEST_FILE)

all: test clean