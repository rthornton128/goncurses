BUILD_FLAGS=-Wl,-Bsymbolic-functions

build:
	CGO_LDFLAGS_ALLOW=${BUILD_FLAGS} go build

