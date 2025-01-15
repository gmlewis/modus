# Modus MoonBit SDK

This is an experimental SDK for Modus for the [MoonBit] programming language.

[MoonBit]: https://www.moonbitlang.com/

## Local Development

To work on the MoonBit SDK locally, first clone the repo and then run the following commands:

```bash
$ cd cli
$ ./build.sh
$ cd ../sdk/moonbit/examples/simple
$ ./update.sh
$ ./build.sh
$ ./modus-local.sh install MoonBit
$ ./modus-local.sh dev
```

## Status

The code has been updated to support compiler:

```bash
$ moon version --all
moon 0.1.20250113 (a18570d 2025-01-13) ~/.moon/bin/moon
moonc v0.1.20250113+2f846af8e ~/.moon/bin/moonc
moonrun 0.1.20250113 (a18570d 2025-01-13) ~/.moon/bin/moonrun
```
