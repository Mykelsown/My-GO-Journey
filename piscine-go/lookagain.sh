#!/bin/bash

find . -type f -name '*.sh' -printf '%f\n' | sort -r | cut -d'.' -f1