# Modus MoonBit SDK

This is an experimental [Modus SDK] for the [MoonBit] programming language.

[Modus SDK]: https://github.com/gmlewis/modus
[MoonBit]: https://www.moonbitlang.com/

## Local Development

To work on the MoonBit SDK locally, first clone the repo and then run the following commands:

```bash
$ REPODIR=$(pwd)
$ cd cli
$ ./build.sh
$ cd ${REPODIR}/sdk/moonbit/examples/hello-world
$ ./update.sh
$ ./build.sh
$ ./modus-local.sh sdk install moonbit
$ ./modus-local.sh dev  # Visit the API Explorer and try it out
# Now update the local Modus Runtime to support more examples:
$ cd ${REPODIR}/runtime
$ ./build.sh
# Now update the local Modus MoonBit SDK to support more examples:
$ cd ${REPODIR}/sdk/go/tools/modus-moonbit-build
$ ./install.sh
```

## Status

The code has been updated to support compiler:

```bash
$ moon version --all
moon 0.1.20250113 (a18570d 2025-01-13) ~/.moon/bin/moon
moonc v0.1.20250113+2f846af8e ~/.moon/bin/moonc
moonrun 0.1.20250113 (a18570d 2025-01-13) ~/.moon/bin/moonrun
```
