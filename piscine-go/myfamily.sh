#!/bin/bash
curl 'https://acad.learn2earn.ng/assets/superhero/all.json' | jq --arg id "$HERO_ID" ' .[] | select(.id == ($id | tonumber)) | .connections.relatives' |
sed 's/"//g'