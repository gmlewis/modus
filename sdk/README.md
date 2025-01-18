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

The following MoonBit examples work:

- [x] `sdk/moonbit/examples/hello-option-empty-string`
- [x] `sdk/moonbit/examples/hello-option-none`
- [x] `sdk/moonbit/examples/hello-option-some-string`
- [x] `sdk/moonbit/examples/hello-world`
