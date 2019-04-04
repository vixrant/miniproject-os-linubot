output_name = linu

install:
	go install

uninstall:
	rm -f $(GOPATH)/bin/$(output_name)

build:
	go build *.go -o $(output_name)
