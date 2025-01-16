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
$ ./modus-local.sh dev
```
