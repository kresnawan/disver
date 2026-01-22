RUST_DIR=internal/crypto/rust
RUST_LIB=$(RUST_DIR)/target/release/libcrypto_rust.a
BINARY_NAME=disver

all: build

$(RUST_LIB): $(shell find $(RUST_DIR)/src -type f -name "*.rs")
	@echo "Building Rust library..."
	cd $(RUST_DIR) && cargo build --release

build: $(RUST_LIB)
	@echo "Building Go node..."
	go build -o $(BINARY_NAME) cmd/node/main.go

run: build
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)
	cd $(RUST_DIR) && cargo clean