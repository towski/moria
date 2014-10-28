#!/bin/bash
trap 'echo $(jobs -p); kill $(jobs -p)' EXIT
LD_LIBRARY_PATH=./:/usr/local/lib ./moria
