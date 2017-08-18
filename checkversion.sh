#!/bin/bash

curl "https://raw.githubusercontent.com/AppIns/Netplan/master/VERSION" -o CURRENT_VERSION 2> /dev/null

if diff VERSION CURRENT_VERSION > /dev/null; then
  echo
  echo "Version file matches"
  echo "The repository is up to date"
else
  echo
  echo "The server version does not match the repository online!"
  echo "The server is likley outdated."
fi
