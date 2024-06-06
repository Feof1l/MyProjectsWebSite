.PHONY:
.SILIENT:

build:
	go build -o main.go
run: build
	./main.go



