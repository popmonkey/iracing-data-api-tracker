#!/usr/bin/env bash

set -ex

GO_BIN=/usr/local/go/bin/go

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
KEYFILE=${1}
DOC_DIR=${2:-${SCRIPT_DIR}/../iracing-data-api-doc}

pushd $SCRIPT_DIR

${GO_BIN} run fetch.go ${KEYFILE} .creds ${DOC_DIR}

pushd ${DOC_DIR}

diffcount=$(git diff doc.json | wc -c)

if [ $diffcount -gt 0 ]; then
    tmpfile=$(mktemp /tmp/iracing-doc-changes.XXXXXX)
    git diff --stat doc.json >> $tmpfile
    git add doc.json
    git commit -F $tmpfile doc.json
    git push
    rm $tmpfile
else
    echo "no changes"
fi
