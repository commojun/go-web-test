#!/bin/bash
echo build domainfinder
go build -o domainfinder
cd ../synonyms
go build -o ../domainfinder/lib/synonyms
cd ../available
go build -o ../domainfinder/lib/available
cd ../sprinkle
go build -o ../domainfinder/lib/sprinkle
cd ../coolify
go build -o ../domainfinder/lib/coolify
cd ../domainify
go build -o ../domainfinder/lib/domainify

