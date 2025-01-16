#!/bin/bash -ex

MOONBIT_SDK_VERSION=v0.16.4

# Detect the operating system
OS="$(uname -s)"

# Set tar options based on the OS
case "$OS" in
    Darwin*)  # MacOS
        TAR_OPTIONS="--no-xattrs --disable-copyfile"
        ;;
    Linux*)   # Linux
        TAR_OPTIONS=""
        ;;
    CYGWIN*|MINGW*|MSYS*)  # Windows (Git Bash, MSYS, etc.)
        TAR_OPTIONS=""
        ;;
    *)
        echo "Unsupported OS: $OS"
        exit 1
        ;;
esac

# Define the combinations of GOOS and GOARCH
platforms=(
  "darwin amd64"
  "darwin arm64"
  "linux amd64"
  "linux arm64"
  "windows amd64"
  "windows arm64"
)

# Loop through each platform
for platform in "${platforms[@]}"; do
  # Split the platform into GOOS and GOARCH
  read -r GOOS GOARCH <<< "$platform"

  # Export GOOS and GOARCH
  export GOOS GOARCH

  # Build the binary
  go build

  # Define the TAR file name
  TARFILE="runtime_${MOONBIT_SDK_VERSION}_${GOOS}_${GOARCH}.tar.gz"

  # Move the runtime binary
  mv runtime ../modus_runtime

  # Create the tarball and copy it to ${HOME}/Downloads
  pushd .. > /dev/null && tar ${TAR_OPTIONS} -zcvf "${TARFILE}" LICENSE README.md modus_runtime && cp "${TARFILE}" ${HOME}/Downloads && popd > /dev/null
done

mkdir -p ${HOME}/.modus/runtime/${MOONBIT_SDK_VERSION}
pushd ${HOME}/.modus/runtime/${MOONBIT_SDK_VERSION} > /dev/null
export GOOS=darwin
export GOARCH=arm64
tar xf ${HOME}/Downloads/runtime_${MOONBIT_SDK_VERSION}_${GOOS}_${GOARCH}.tar.gz
popd > /dev/null
pushd ${HOME}/Downloads > /dev/null
sha256sum runtime_${MOONBIT_SDK_VERSION}_*.tar.gz > checksums.txt
popd > /dev/null
