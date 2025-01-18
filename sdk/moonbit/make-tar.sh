#!/bin/bash -ex

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

# Remove the unwanted directories before creating tar file
rm -rf templates/default/.mooncakes templates/default/target
# Create the tar file
MOONBIT_SDK_VERSION=v0.16.5
TARFILE=templates_moonbit_${MOONBIT_SDK_VERSION}.tar.gz
tar ${TAR_OPTIONS} -zcvf ${TARFILE} templates

# Copy the files into place
mkdir -p ${HOME}/.modus/sdk/moonbit/${MOONBIT_SDK_VERSION}
cp sdk.json ${TARFILE} ${HOME}/.modus/sdk/moonbit/${MOONBIT_SDK_VERSION}
# Make them easy to upload to GitHub
cp sdk.json ${TARFILE} ~/Downloads
