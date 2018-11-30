#!/bin/bash

# Wipe out previous build
rm -rf ./linuxsetQoSapp

# Create project
flogo create -f set_qo_s.json linuxsetQoSapp

# Build the app for Linux64
cd linuxsetQoSapp
flogo build -e

# Execute the app
cd bin
./linuxsetQoSapp
