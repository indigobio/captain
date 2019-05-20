build:
	docker build -t indigobio/captain .

run:
	docker run indigobio/captain

b:
	go get -v -d github.com/indigobio/captain
	go install -v github.com/indigobio/captain/cmd/captain

watch:
	docker run -it --rm --name captain -v "$$PWD":/go/src/github.com/indigobio/captain -w /go/src/github.com/indigobio/captain golang:1.4 watch -n 1 make b

goconvey:
	goconvey -timeout 10s


cross:
	mkdir -p build
	gox --os windows --os linux --os darwin --arch 386 --arch amd64 github.com/indigobio/captain/cmd/captain
	mv captain_darwin_* captain_linux_* captain_windows_* build
