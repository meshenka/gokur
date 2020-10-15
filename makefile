build:
  go build -o dist/local/gokur cmd/gokur/gokur.go

run: build
  dist/local/gokur