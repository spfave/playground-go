# ---------------------------------------------------------------------------------------
# CHECKS

## Audit checks
audit:
	go mod tidy -diff
	go mod verify
	go vet ./...
	lint

## Format code
format:
	gofmt -w -s .

## Run Tests
test:
	go test -v ./...