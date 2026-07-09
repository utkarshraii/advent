# Advent of Code — Go
# Usage:
#   make new DAY=1          scaffold 2025/day01 from the template
#   make new DAY=1 YEAR=2024
#   make run DAY=1          run against input.txt
#   make sample DAY=1       run against sample.txt
#   make test DAY=1         test a single day
#   make test               test everything

YEAR ?= 2025
PAD  := $(shell printf '%02d' $(DAY))
DIR  := ./$(YEAR)/day$(PAD)

.PHONY: new run sample test tidy fmt

new:
	@scripts/new.sh $(DAY) $(YEAR)

run:
	@go run $(DIR)

sample:
	@go run $(DIR) -sample

test:
ifndef DAY
	@go test ./...
else
	@go test $(DIR)
endif

tidy:
	@go mod tidy

fmt:
	@go fmt ./...
