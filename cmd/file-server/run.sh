#!/bin/bash
docker build -t file-server .
docker run -p 8080:8080 -p 443:443 file-server