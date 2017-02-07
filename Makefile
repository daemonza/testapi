SERVICE_NAME=testapi
VERSION=latest
QUAYVERSION = "daemonza/$(SERVICE_NAME):$(VERSION)"

GOCMD = go
GOBUILD=env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build -v

GO_FILES = ./main.go
OBJ_FILES = $(basename $(GO_FILES))

%: %.go
		$(GOBUILD) $<

build-go: $(OBJ_FILES)

test:
		go test ./...

build: clean test build-go

docker: build
		echo $(VERSION)
		docker build --no-cache -t $(SERVICE_NAME) .
		docker tag $(SERVICE_NAME) $(QUAYVERSION)

image: docker
		docker push $(QUAYVERSION)

clean:
		rm -f $(notdir $(OBJ_FILES))


