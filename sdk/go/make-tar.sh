#!/bin/bash -ex
TARFILE=templates_go_v0.16.1.tar.gz
tar --no-xattrs --disable-copyfile -zcvf ${TARFILE} templates
cp sdk.json ${TARFILE} ~/Downloads
