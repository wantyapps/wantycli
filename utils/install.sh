#!/usr/bin/env bash

echo "Installing to /usr/local/bin..."
go build -o /usr/local/bin/wanty .
echo "Done."
echo "*****************************************************"
echo "            You can now execute \"wanty\"."
echo "*****************************************************"
