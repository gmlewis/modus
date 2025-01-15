#!/bin/bash -ex
rm -rf templates/default/.mooncakes templates/default/target
TARFILE=templates_moonbit_v0.16.2.tar.gz
tar --no-xattrs --disable-copyfile -zcvf ${TARFILE} templates
cp sdk.json ${TARFILE} ~/Downloads
