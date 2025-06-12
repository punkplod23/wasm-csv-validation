# CSV Wasm Validator (Work in Progress)

This repository houses a WebAssembly (Wasm) application designed for efficient client-side CSV validation directly in the browser. The goal is to provide a high-performance solution for validating large CSV files without needing to send them to a server.

> **Note:** This project is currently a work in progress and under active development.

---

## 🚀 Project Overview

The core idea is to leverage the near-native performance of WebAssembly to process CSV data rapidly and perform various validation checks (e.g., schema validation, data type checks, custom rules). By running the validation logic in Wasm, we aim to offload heavy computation from the main JavaScript thread and offer a responsive user experience, even with substantial datasets.

A key aspect of this project involves exploring and implementing multi-threading within WebAssembly to further enhance performance, particularly for very large files. This will utilize browser features like `SharedArrayBuffer` and require careful handling of cross-origin isolation policies.

---

## ✨ Features (Planned)

- **Fast CSV Parsing:** Efficiently parse large CSV files directly in the browser.
- **Customizable Validation Rules:** Support for defining and applying various validation rules (e.g., column count, data types, uniqueness, regex patterns).
- **Error Reporting:** Clear and informative error messages with line and column numbers.
- **Multi-threaded Processing:** Leverage WebAssembly threads for parallel validation of CSV chunks to improve performance.
- **User-friendly Interface (Future):** A simple web interface for uploading CSV files and viewing validation results.

---

## 🛠️ Technologies Under Consideration (WIP)

We are actively evaluating the best language and toolchain for this Wasm application, focusing on performance, binary size, development experience, and threading capabilities.

- **Rust:** Known for its strong performance, memory safety, and mature Wasm tooling (`wasm-bindgen`, `wasm-bindgen-rayon` for threading).
- **Go / TinyGo:** Exploring the trade-offs between standard Go's comprehensive libraries and TinyGo's focus on smaller Wasm binaries and improved performance for embedded/Wasm targets.
- **AssemblyScript:** A TypeScript-like language that compiles directly to Wasm, offering a familiar syntax for JavaScript developers.
- **C / C++ (via Emscripten):** Provides ultimate low-level control and performance, with established Wasm compilation.

---

## 🚀 Getting Started (WIP)

Detailed instructions for building and running the project will be added here as development progresses. Currently, the focus is on core Wasm module development.

---

## 🤝 Contributing (WIP)

We welcome contributions! More information on how to contribute, including coding guidelines and development setup, will be provided soon.

---

## 📜 License

[Choose a license and add it here, e.g., MIT, Apache 2.0]

---

Stay tuned for updates as this project evolves!
