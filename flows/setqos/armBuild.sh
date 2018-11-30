#!/bin/bash

# Wipe out previous build
rm -rf ./armsetQoSapp

# Create project
flogo create -f set_qo_s.json armsetQoSapp

#Build the app for ARM32
cd armsetQoSapp
GOOS=linux GOARCH=arm GOARM=7 flogo build -e




