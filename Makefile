build:
  @go build  -o bin/ecom  cmd/mani.go

test:
  @go test ./...

run: 
  @./bin/ecom 