#!/bin/sh
docker rm -f $(docker ps -aq)
docker volume prune -f
rm -rf crypto-config/
rm -rf channel-artifacts/
rm system-genesis-block/genesis.block