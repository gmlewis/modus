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
$ cd ${REPODIR}/sdk/moonbit/examples/test-suite
$ ./update.sh
$ ./build.sh
$ ./modus-local.sh sdk install moonbit
$ ./modus-local.sh dev  # This will most likely fail until you run the following steps:
# Update the local Modus Runtime to support more examples:
$ cd ${REPODIR}/runtime
$ ./build.sh
# Update the local Modus MoonBit SDK to support more examples:
$ cd ${REPODIR}/sdk/go/tools/modus-moonbit-build
$ ./install.sh
```

## End-to-end Test Suite

To run the end-to-end test suite, cd to either the [`sdk/go`](/sdk/go) directory
or the [`sdk/go/tools/modus-moonbit-build`](/sdk/go/tools/modus-moonbit-build) directory and run:

```bash
$ ./test-all.sh
```

To add a new plugin or endpoint to the end-to-end test suite, update this file:
[end-to-end-tests.json](/sdk/go/tools/modus-moonbit-build/end-to-end-tests.json)

The `query` field is easy to fill in by opening the Chrome Devtools "Network"
tab while running the Modus API Explorer and clicking on the most recent
"graphql" request. Then under "Payload" click on "View Source" and copy
the query string to paste into the JSON file directly.

## Status

The following MoonBit examples work:

- [x] `sdk/moonbit/examples/anthropic-functions`
- [x] `sdk/moonbit/examples/hello-tuples`
- [x] `sdk/moonbit/examples/neo4j`
- [x] `sdk/moonbit/examples/simple`
- [x] `sdk/moonbit/examples/test-suite`
- [x] `sdk/moonbit/examples/time-example`
- [x] `sdk/moonbit/examples/youtube-walkthrough`

The code has been updated to support compiler:

```bash
$ moon version --all
moon 0.1.20250310 (df3bb14 2025-03-10) ~/.moon/bin/moon
moonc v0.1.20250310+a7a1e9804 ~/.moon/bin/moonc
moonrun 0.1.20250310 (df3bb14 2025-03-10) ~/.moon/bin/moonrun
```
