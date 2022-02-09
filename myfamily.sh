#! /bin/bash
curl https://zone01normandie.org/assets/superhero/all.json | jq ' .[] | select( .id | contains('$HERO_ID'))' | jq '.connections' | jq '.relatives' | sed "s/\"//g"