#!/bin/bash

_main() {
    local file=${1}
    local content="$(cat ${file})"

    cat <<EOF
{
    "file": "${file}",
    "content": "${content}"
}
EOF
}

_main "$@"
