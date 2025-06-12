# Makefile for building Rust project

# Default target
all: build

# Build the Rust project
build:
	cargo build --release

# Run the Rust project
run:
	cargo run ..\csv-data\Crime_Data_from_2020_to_Present.csv

# Clean the build artifacts
clean:
	cargo clean

# Test the Rust project
test:
	cargo test

.PHONY: all build run clean test