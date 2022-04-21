#!make
include config.env
SHELL := /bin/bash
export $(shell sed 's/=.*//' config.env)

# for development, create a file called config.env (clone example.env)
# and fill in your keys
run:
	go run ./cli \
		--bookstack-url=$(BookStackEndpoint) \
		--token-id=$(BookStackAPITokenID) \
		--token-secret=$(BookStackAPITokenSecret)
build:
	go build -o ./bin/bookstack2site ./cli