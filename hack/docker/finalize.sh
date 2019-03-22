#!/bin/sh
set -x
set -e

# Create user for Chi
addgroup -S chimusic
adduser -G chimusic -H -D -g 'chimusic User' chimusic -h /data/chimusic -s /bin/bash && usermod -p '*' chimusic && passwd -u chimusic
echo "export CHIMUSIC_CUSTOM=${CHIMUSIC_CUSTOM}" >> /etc/profile

# Final cleaning
rm /app/chimusic/docker/finalize.sh
rm /app/chimusic/docker/nsswitch.conf
