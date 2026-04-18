default:
    @just --list

# build arvanstatus-cli binary
build:
    @echo '{{ BOLD + CYAN }}Building arvanstatus-cli!{{ NORMAL }}'
    go build -o arvanstatus-cli .

# run arvanstatus-cli
run: build
    ./arvanstatus-cli

# update go packages
update:
    go get -u ./...
    go mod tidy

# run tests
test:
    go test -v ./... -covermode=atomic -coverprofile=coverage.out

# run golangci-lint
lint:
    golangci-lint run

# dry-run a release locally (no publish)
release-snapshot:
    goreleaser release --snapshot --clean
