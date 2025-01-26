# Modus SDKs

This directory contains the source code for the SDKs of each language that you can write Modus apps
in.

Typically, you will not need to download this code directly, as the SDK will automatically be
installed when you create a new Modus project. However, you may wish to explore the example projects
in each SDK.

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

- [x] `sdk/moonbit/examples/test-suite`
