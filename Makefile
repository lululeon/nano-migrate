.PHONY: build

# invoke make with VERSION=n.n.n
# WARN: any vars here can be overridden by same-named env vars
TAG := -$(VERSION)
ifeq ($(VERSION),)
TAG :=
endif

build:
	  go build -o ./build/nnmg$(TAG) nnmg.go
clean:
		rm ./build/*
uninstall:
	  rm $(GOPATH)/bin/nano-migrate
