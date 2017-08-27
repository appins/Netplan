#!/bin/bash

curl "https://raw.githubusercontent.com/AppIns/Netplan/master/VERSION" -o CURRENT_VERSION 2> /dev/null

echo

if diff VERSION CURRENT_VERSION > /dev/null; then
  echo "Version file matches"
  echo "The repository is up to date"
else
  echo "The server version does not match the repository online!"
  echo "The server is likley outdated."
fi

VSIZE=$(stat --printf="%s" VERSION)
CVSIZE=$(stat --printf="%s" CURRENT_VERSION)
if [[ $VSIZE != $CVSIZE ]]; then
  echo
  echo "----------WARNING---------------"
  echo "The VERSION file online and the local one have different sizes"
  echo "You might be out of date by a major version"
  echo "---------------------------------"
fi

echo
