#!/bin/sh

(cd src &&
 rm -f libGoTest.dylib &&
 ln -s ${BUILT_PRODUCTS_DIR}/${EXECUTABLE_PATH} libGoTest.dylib &&
 /usr/local/go/bin/go build -o main &&
 rm -f libGoTest.dylib &&
 mv main ${BUILT_PRODUCTS_DIR}/${EXECUTABLE_FOLDER_PATH})
