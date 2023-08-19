.PHONY: build

# invoke make with VERSION=n.n.n
# WARN: any vars here can be overridden by same-named env vars. DO NOT USE THESE FOR LOCAL DEV.
TAG := -$(VERSION)
ifeq ($(VERSION),)
TAG :=
endif

setup:
		chmod +x ./devtools/setup.sh
		./devtools/setup.sh

# Run a local test pg instance. Assumes you have docker and docker daemon is running
pgup:
		docker compose --env-file ./.env -f ./devtools/compose.yml up -d
pgdown:
		docker compose --env-file ./.env -f ./devtools/compose.yml down


# building and testing
build:
	  go build -o ./build/nnmg$(TAG) nnmg.go
test:
		make clean
		make build
		go test

# see list of migrations
ls:
		ls -al $(HOME)/nn_datavol/migrations


clean:
		rm ./build/*
uninstall:
	  rm $(GOPATH)/bin/nano-migrate
