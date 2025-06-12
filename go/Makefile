GOCMD   =  go
GOBUILD =  $(GOCMD)  build
GORUN = $(GOCMD) run
TINYGOCMD   =  tinygo
TINYGOBUILD = $(TINYGOCMD) build

create-go:
	$(GOCMD) mod init tutorialone

create-tinygo-project:
	$(TINYGOBUILD) mod init tiny-tutorialone

compile-wasm:export GOOS=js
compile-wasm:export GOARCH=wasm
compile-wasm:
	$(GOBUILD) -o main.wasm

compile-tiny-wasm:export GOOS=js
compile-tiny-wasm:export GOARCH=wasm
compile-tiny-wasm:
	$(TINYGOBUILD) -o main.wasm
	
# Clean command
clean:
	rm -f $(OUT)

.PHONY: build clean

run-server:
	cd html && $(GORUN) server.go	

install-go:
	sudo apt install go

install-go-tiny:
	sudo apt install tinygo

update-go:
	sudo apt upgrade go

update-go-tiny:
	sudo apt upgrade tinygo

copy-wasm:
	cp main.wasm html/main.wasm

remove-wasm:
	cp rm main.wasm

compile-wasm-linux:
	$(MAKE) compile-wasm 
	$(MAKE) copy-wasm
	$(MAKE) run-server

compile-tiny-wasm-linux:
	$(MAKE) compile-tiny-wasm 
	$(MAKE) copy-wasm
	$(MAKE) run-server