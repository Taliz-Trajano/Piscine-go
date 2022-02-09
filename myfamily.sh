#! /bin/bash
curl -s https://zone01normandie.org/assets/superhero/all.json | jq ".[] | select(.id ==1) .connections | .relatives" | tr -d "\""