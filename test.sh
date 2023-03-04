#!/bin/sh

$(docker inspect --type=image swaggo/swa &> /dev/null) || echo "OK"
