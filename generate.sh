#!/bin/bash
# File: generate.sh
# Description: run to generate api

# Make sure go swagger is installed
echo -e "\033[4mEnsuring swagger generator...\033[0m"
(cd vendor/github.com/go-swagger/go-swagger && go install .)

# Get latest swagger spec from citadel
echo -e "\033[4mGetting latest swagger...\033[0m"
wget "https://api.fortifi.io/swagger.yaml"

# Generate new models & client from spec
echo -e "\033[4mGenerating fortifi API (Swagger spec has changed)...\033[0m"
(swagger generate client && go fix ./...)
