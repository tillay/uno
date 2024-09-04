#!/bin/bash

read -p "Play Singleplayer (1) or Against Computer (2): " mode

if [[ "$mode" == "1" ]]; then
    replacement="game()"
elif [[ "$mode" == "2" ]]; then
    replacement="bot()"
else
    echo "Invalid option. Defaulting to 1"
fi
read -p "Number of cards to start with: " deal
sed -i'' -e '$s/.*/'"$replacement"'/' main.py
sed -i'' -e '/^deal =/c\deal = '"$deal"'    # num cards to start with' main.py

echo "In the future, run python3 main.py to play!"

sleep 1
python3 main.py

