#!/bin/bash

_data() {
    local file=${1}
    local content=${2}

    cat <<EOF
{
    "file": "${file}",
    "content": "${content}"
}
EOF
}

_main() {
    local file=${1}
    local content="${2}"

    echo "${content}" > ${file}
    echo "$(_data ${file} \"${content}\")"
}

_main "$@"
