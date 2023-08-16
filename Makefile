.PHONY: build

# invoke make with VERSION=n.n.n
# WARN: any vars here can be overridden by same-named env vars
TAG := -$(VERSION)
ifeq ($(VERSION),)
TAG :=
endif

build:
	  go build -o nnmg$(TAG) ./main.go
clean:
		rm nnmg*
uninstall:
	  rm $(GOPATH)/bin/nano-migrate
