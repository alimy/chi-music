#!/bin/sh

create_volume_subfolder() {
    # Create VOLUME subfolder
    for f in /data/chimusic/data /data/chimusic/conf /data/chimusic/log; do
        if ! test -d $f; then
            mkdir -p $f
        fi
    done
}

setids() {
    PUID=${PUID:-1000}
    PGID=${PGID:-1000}
    groupmod -o -g "$PGID" chimusic
    usermod -o -u "$PUID" chimusic
}

setids
create_volume_subfolder

# Exec CMD or S6 by default if nothing present
if [ $# -gt 0 ];then
    exec "$@"
else
    exec /bin/s6-svscan /app/chimusic/docker/s6/
fi
