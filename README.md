# numerator

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/f22840a3dc0b42f2810c8cc1e531faa2)](https://app.codacy.com/app/slavyan85/numerator?utm_source=github.com&utm_medium=referral&utm_content=slavyan85/numerator&utm_campaign=Badge_Grade_Dashboard)

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
