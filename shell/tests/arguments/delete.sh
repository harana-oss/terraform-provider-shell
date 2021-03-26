#!/bin/bash

_data() {
    echo "{}"
}

_main() {
    local file=${1}
    if [ -f "${file}" ]; then
        rm -f ${file}
    fi

    echo "DELETE SCRIPT"
    echo "$(_data)"
}

_main "$@"
