#!/bin/bash

log() {
	local type="$1"; shift
	printf '%s [%s] [Entrypoint]: %s\n' "$(date --rfc-3339=seconds)" "$type" "$*"
}

moredoc_note() {
	log Note "$@"
}

_main() {
    cd ~
    moredoc_note "Current path is:"
    pwd
    moredoc_note "Current folder contains:"
    ls
    if [ -f app.toml ]
    then
        moredoc_note "Start server"
        exec ./workspace/moredoc serve
    else
        moredoc_note "Init server"
        ESCAPED_ORIGIN=$(printf '%s\n' "root:root@tcp(localhost:3306)/moredoc?charset=utf8mb4&loc=Local&parseTime=true" | sed -e 's/[\/&]/\\&/g')
        ESCAPED_REPLACE=$(printf '%s\n' "${MYSQL_CONNECTION}" | sed -e 's/[\/&]/\\&/g')
        sed "s/$ESCAPED_ORIGIN/$ESCAPED_REPLACE/g" ./workspace/app.example.toml > ./workspace/app.toml
        ls ./workspace
        moredoc_note "Init DB"
        cd workspace/
        ./moredoc syncdb
        moredoc_note "Start server"
        exec ./moredoc serve
    fi
}

_main "$@"