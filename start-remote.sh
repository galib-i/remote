#!/bin/bash

PORT=8080

cleanup() {
    echo -e "\n[Firewall] Closing port $PORT..."
    sudo ufw delete allow $PORT/tcp > /dev/null
    echo "[Go-Remote Server] Shutting down"
    exit 0
}

# ctrl+c and killing always runs the cleanup
trap cleanup SIGINT SIGTERM SIGHUP

# Open the port
echo "[Firewall] Opening port $PORT..."
sudo ufw allow $PORT/tcp > /dev/null

echo "[Go-Remote Server] Starting on port $PORT..."
PORT=$PORT ./remote-server-linux

./remote-server-linux

cleanup