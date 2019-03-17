### --------------------------------------------------------------------------------------------------------------------
### Variables
### (https://www.gnu.org/software/make/manual/html_node/Using-Variables.html#Using-Variables)
### --------------------------------------------------------------------------------------------------------------------

GO_PACKAGES := $(shell go list ./...)

# Other config
NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

tests:
	@printf "$(OK_COLOR)==> Running tests$(NO_COLOR)\n"
	@go test -race $(GO_PACKAGES)

help:
	@echo "---------------------------------------------"
	@echo "List of available targets:"
	@echo "  tests                  - Runs the unit tests"
