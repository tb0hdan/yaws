#!/usr/bin/env sh

if [ -z "$1" ]; then
    echo "User/Group required"
    exit 127
fi

# Build mocks
cd /app; mockery --all
rm -rf /app/mocks/yaws/int
mv /app/mocks/yaws/internal /app/mocks/yaws/int
chown -R $1 /app/mocks

