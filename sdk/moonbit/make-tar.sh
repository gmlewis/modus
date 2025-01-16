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
TARFILE=templates_moonbit_v0.16.2.tar.gz
tar ${TAR_OPTIONS} -zcvf ${TARFILE} templates

# Copy the files into place
cp sdk.json ${TARFILE} ${HOME}/.modus/sdk/moonbit/v0.16.2
# Make them easy to upload to GitHub
cp sdk.json ${TARFILE} ~/Downloads
