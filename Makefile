RUST_DIR=internal/crypto_rust
RUST_LIB=$(RUST_DIR)/target/release/libcrypto_rust.a
BINARY_DIR=./cmd/build/disver

all: build

$(RUST_LIB): $(shell find $(RUST_DIR)/src -type f -name "*.rs")
	@echo "Building Rust library..."
	rm -f $(BINARY_DIR)
	cd $(RUST_DIR) && cargo clean
	cd $(RUST_DIR) && cargo build --release

build: $(RUST_LIB)
	@echo "Building Go binary..."
	go build -o $(BINARY_DIR) ./cmd/node

run: build
	@echo ""
	@$(BINARY_DIR)

clean:
	rm -f $(BINARY_DIR)
	cd $(RUST_DIR) && cargo clean