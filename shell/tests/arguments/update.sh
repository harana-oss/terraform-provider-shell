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
    local content=${2}

    if [ -f "${file}" ]; then
        echo "${content}" > ${file}
    fi

    echo "$(_data ${file} \"${content}\")"
}

_main "$@"
