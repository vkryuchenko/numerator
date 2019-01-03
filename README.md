# numerator
Incremental builds numbers as a service

# Usage
From source:
 * clone repo
 * go get -d -v ./...
 * go build -o numerator src/numerator.go
 * ./numerator --help

From Docker:
 * docker run -p :8888 slavyan85/numerator

# Methods
 * GET /\<name> -- increment \<name> value and return it
 * POST /\<name> -- set value for \<name>
 * DELETE /\<name> -- drop \<name>
