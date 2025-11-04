![GoVyOS Logo](./assets/govyosLogo.png)

[![Go Report Card](https://goreportcard.com/badge/github.com/aevrex/govyos)](https://goreportcard.com/report/github.com/aevrex/govyos)
[![GoDoc](https://godoc.org/github.com/aevrex/govyos?status.svg)](https://godoc.org/github.com/aevrex/govyos)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/aevrex/govyos)](https://github.com/aevrex/govyos/releases)
[![License: MIT](https://img.shields.io/badge/License-Apache_2.0-yellow.svg)](https://opensource.org/license/apache-2-0)

`govyos` is a modern, actively maintained Go SDK for interacting with the VyOS router operating system's HTTP API.

This project was created as a successor to the `vyos-go` library. It will provide a complete, robust, and idiomatic Go interface for configuring, managing, and monitoring VyOS devices.

---

## ⚠️ Project Status ⚠️

This library is currently in **active development**. The API may have breaking changes until a stable `v1.0.0` release. Feedback and contributions are welcome.

---

## Intended Features

* **Full API Support:** Complete coverage for VyOS configuration (set, show, delete) and operational commands.
* **Actively Maintained:** Aims to provide a reliable, long-term solution for Go-based VyOS automation.
* **Type-Safe:** To provide type-safe helpers for common configuration paths.

---

## Installation

```sh
go get github.com/aevrex/govyos
```

## Contributing

Contributions are welcome. This project is intended to be a community-driven replacement for the old library. Please feel free to submit a Pull Request or open an Issue.

- Fork the repository.
- Create your feature branch (git checkout -b feature/my-new-feature).
- Commit your changes (git commit -am 'Add some feature').
- Push to the branch (git push origin feature/my-new-feature).
- Open a new Pull Request.

## License
This project is licensed under the Apache 2.0 License.

## Acknowledgements
- The VyOS Team for their excellent open-source router OS.
- The author of the original vyos-go for their initial work.
