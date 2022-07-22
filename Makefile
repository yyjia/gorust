.PHONY: build-all
build-all: build-go build-rust

.PHONY: clean
clean: 
	rm -f ./lib/*.a 


.PHONY: build-go
build-go: 
	go build -buildmode=c-archive -o ./lib/libgo.a ./go

.PHONY: build-rust
build-rust: 
	cd rust && cargo build --release
	cp rust/target/release/librust.a ./lib

.PHONY: run
run:
	go run go/main.go