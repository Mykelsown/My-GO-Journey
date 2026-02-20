#!/bin/bash

# Fetch the JSON and extract name, power, and gender for superhero with id 170
curl -s https://acad.learn2earn.ng/assets/superhero/all.json \
  | jq -r '.[] | select(.id == 170) | "\(.name)\n\(.powerstats.power)\n\(.appearance.gender)"'
