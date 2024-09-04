#!/bin/bash

# Function to draw a card
draw_card() {
    local card_num=$1
    local color=$2
    local spaces="  "

    case "$card_num" in
        0)
            echo -e "\e[${color}m ++++++++++ \e[0m${spaces}"
            echo -e "\e[${color}m#####+++++++\e[0m${spaces}"
            echo -e "\e[${color}m##      0 ++\e[0m${spaces}"
            echo -e "\e[${color}m##   __   ++\e[0m${spaces}"
            echo -e "\e[${color}m##  |  |  ++\e[0m${spaces}"
            echo -e "\e[${color}m##  |__|  ++\e[0m${spaces}"
            echo -e "\e[${color}m##        ++\e[0m${spaces}"
            echo -e "\e[${color}m## 0      ++\e[0m${spaces}"
            echo -e "\e[${color}m#######+++++\e[0m${spaces}"
            echo -e "\e[${color}m ########## \e[0m${spaces}"
            ;;
        1)
            echo -e "\e[${color}m ++++++++++ \e[0m${spaces}"
            echo -e "\e[${color}m#####+++++++\e[0m${spaces}"
            echo -e "\e[${color}m##      1 ++\e[0m${spaces}"
            echo -e "\e[${color}m##   /|   ++\e[0m${spaces}"
            echo -e "\e[${color}m##    |   ++\e[0m${spaces}"
            echo -e "\e[${color}m##   _|_  ++\e[0m${spaces}"
            echo -e "\e[${color}m##        ++\e[0m${spaces}"
            echo -e "\e[${color}m## 1      ++\e[0m${spaces}"
            echo -e "\e[${color}m#######+++++\e[0m${spaces}"
            echo -e "\e[${color}m ########## \e[0m${spaces}"
            ;;
        2)
            echo -e "\e[${color}m ++++++++++ \e[0m${spaces}"
            echo -e "\e[${color}m######++++++\e[0m${spaces}"
            echo -e "\e[${color}m##      2 ++\e[0m${spaces}"
            echo -e "\e[${color}m##   __   ++\e[0m${spaces}"
            echo -e "\e[${color}m##   __|  ++\e[0m${spaces}"
            echo -e "\e[${color}m##  |__   ++\e[0m${spaces}"
            echo -e "\e[${color}m##        ++\e[0m${spaces}"
            echo -e "\e[${color}m## 2      ++\e[0m${spaces}"
            echo -e "\e[${color}m#######+++++\e[0m${spaces}"
            echo -e "\e[${color}m ########## \e[0m${spaces}"
            ;;
        3)
            echo -e "\e[${color}m ++++++++++ \e[0m${spaces}"
            echo -e "\e[${color}m######++++++\e[0m${spaces}"
            echo -e "\e[${color}m##      3 ++\e[0m${spaces}"
            echo -e "\e[${color}m##   __   ++\e[0m${spaces}"
            echo -e "\e[${color}m##   __|  ++\e[0m${spaces}"
            echo -e "\e[${color}m##   __|  ++\e[0m${spaces}"
            echo -e "\e[${color}m##        ++\e[0m${spaces}"
            echo -e "\e[${color}m## 3      ++\e[0m${spaces}"
            echo -e "\e[${color}m#######+++++\e[0m${spaces}"
            echo -e "\e[${color}m ########## \e[0m${spaces}"
            ;;
        4)
            echo -e "\e[${color}m ++++++++++ \e[0m${spaces}"
            echo -e "\e[${color}m######++++++\e[0m${spaces}"
            echo -e "\e[${color}m##      4 ++\e[0m${spaces}"
            echo -e "\e[${color}m##        ++\e[0m${spaces}"
            echo -e "\e[${color}m##   |_|  ++\e[0m${spaces}"
            echo -e "\e[${color}m##     |  ++\e[0m${spaces}"
            echo -e "\e[${color}m##        ++\e[0m${spaces}"
            echo -e "\e[${color}m## 4      ++\e[0m${spaces}"
            echo -e "\e[${color}m#######+++++\e[0m${spaces}"
            echo -e "\e[${color}m ########## \e[0m${spaces}"
            ;;
        5)
            echo -e "\e[${color}m ++++++++++ \e[0m${spaces}"
            echo -e "\e[${color}m######++++++\e[0m${spaces}"
            echo -e "\e[${color}m##      5 ++\e[0m${spaces}"
            echo -e "\e[${color}m##   __   ++\e[0m${spaces}"
            echo -e "\e[${color}m##  |__   ++\e[0m${spaces}"
            echo -e "\e[${color}m##   __|  ++\e[0m${spaces}"
            echo -e "\e[${color}m##        ++\e[0m${spaces}"
            echo -e "\e[${color}m## 5      ++\e[0m${spaces}"
            echo -e "\e[${color}m#######+++++\e[0m${spaces}"
            echo -e "\e[${color}m ########## \e[0m${spaces}"
            ;;
        6)
            echo -e "\e[${color}m ++++++++++ \e[0m${spaces}"
            echo -e "\e[${color}m######++++++\e[0m${spaces}"
            echo -e "\e[${color}m##      6 ++\e[0m${spaces}"
            echo -e "\e[${color}m##   __   ++\e[0m${spaces}"
            echo -e "\e[${color}m##  |__   ++\e[0m${spaces}"
            echo -e "\e[${color}m##  |__|  ++\e[0m${spaces}"
            echo -e "\e[${color}m##        ++\e[0m${spaces}"
            echo -e "\e[${color}m## 6      ++\e[0m${spaces}"
            echo -e "\e[${color}m#######+++++\e[0m${spaces}"
            echo -e "\e[${color}m ########## \e[0m${spaces}"
            ;;
        7)
            echo -e "\e[${color}m ++++++++++ \e[0m${spaces}"
            echo -e "\e[${color}m######++++++\e[0m${spaces}"
            echo -e "\e[${color}m##      7 ++\e[0m${spaces}"
            echo -e "\e[${color}m##   __   ++\e[0m${spaces}"
            echo -e "\e[${color}m##    /   ++\e[0m${spaces}"
            echo -e "\e[${color}m##   /    ++\e[0m${spaces}"
            echo -e "\e[${color}m##        ++\e[0m${spaces}"
            echo -e "\e[${color}m## 7      ++\e[0m${spaces}"
            echo -e "\e[${color}m#######+++++\e[0m${spaces}"
            echo -e "\e[${color}m ########## \e[0m${spaces}"
            ;;
        8)
            echo -e "\e[${color}m ++++++++++ \e[0m${spaces}"
            echo -e "\e[${color}m######++++++\e[0m${spaces}"
            echo -e "\e[${color}m##      8 ++\e[0m${spaces}"
            echo -e "\e[${color}m##   __   ++\e[0m${spaces}"
            echo -e "\e[${color}m##  |__|  ++\e[0m${spaces}"
            echo -e "\e[${color}m##  |__|  ++\e[0m${spaces}"
            echo -e "\e[${color}m##        ++\e[0m${spaces}"
            echo -e "\e[${color}m## 8      ++\e[0m${spaces}"
            echo -e "\e[${color}m#######+++++\e[0m${spaces}"
            echo -e "\e[${color}m ########## \e[0m${spaces}"
            ;;
        9)
            echo -e "\e[${color}m ++++++++++ \e[0m${spaces}"
            echo -e "\e[${color}m######++++++\e[0m${spaces}"
            echo -e "\e[${color}m##      9 ++\e[0m${spaces}"
            echo -e "\e[${color}m##   __   ++\e[0m${spaces}"
            echo -e "\e[${color}m##  |__|  ++\e[0m${spaces}"
            echo -e "\e[${color}m##   __|  ++\e[0m${spaces}"
            echo -e "\e[${color}m##        ++\e[0m${spaces}"
            echo -e "\e[${color}m## 9      ++\e[0m${spaces}"
            echo -e "\e[${color}m#######+++++\e[0m${spaces}"
            echo -e "\e[${color}m ########## \e[0m${spaces}"
            ;;
        10)
            echo -e "\e[90m ++++++++++ \e[0m${spaces}"
            echo -e "\e[90m#####+++++++\e[0m${spaces}"
            echo -e "\e[90m##\e[31mXXXX\e[34mXXXX\e[90m++\e[0m${spaces}"
            echo -e "\e[90m##\e[31mXXXX\e[34mXXXX\e[90m++\e[0m${spaces}"
            echo -e "\e[90m##\e[31mXXXX\e[34mXXXX\e[90m++\e[0m${spaces}"
            echo -e "\e[90m##\e[33mXXXX\e[32mXXXX\e[90m++\e[0m${spaces}"
            echo -e "\e[90m##\e[33mXXXX\e[32mXXXX\e[90m++\e[0m${spaces}"
            echo -e "\e[90m##\e[33mXXXX\e[32mXXXX\e[90m++\e[0m${spaces}"
            echo -e "\e[90m#######+++++\e[0m${spaces}"
            echo -e "\e[90m ########## \e[0m${spaces}"
            ;;
        e)
            echo -e "\e[${color}m ++++++++++ \e[0m${spaces}"
            echo -e "\e[${color}m######++++++\e[0m${spaces}"
            echo -e "\e[${color}m##        ++\e[0m${spaces}"
            echo -e "\e[${color}m##        ++\e[0m${spaces}"
            echo -e "\e[${color}m##        ++\e[0m${spaces}"
            echo -e "\e[${color}m##        ++\e[0m${spaces}"
            echo -e "\e[${color}m##        ++\e[0m${spaces}"
            echo -e "\e[${color}m##        ++\e[0m${spaces}"
            echo -e "\e[${color}m#######+++++\e[0m${spaces}"
            echo -e "\e[${color}m ########## \e[0m${spaces}"
            ;;
        *)
            echo "Invalid card number."
            ;;
    esac
}

# Ensure we have enough arguments
if [ "$#" -lt 2 ] || [ $(( $# % 2 )) -ne 0 ]; then
    echo "Usage: $0 card_number1 card_color1 [card_number2 card_color2 ...]"
    exit 1
fi

# Create temporary files for card outputs
temp_files=()
trap 'rm -f "${temp_files[@]}"' EXIT

while [ "$#" -gt 0 ]; do
    card_num=$1
    card_color=$2
    temp_file=$(mktemp)
    draw_card "$card_num" "$card_color" > "$temp_file"
    temp_files+=("$temp_file")
    shift 2
done

# Use paste to display cards side by side
paste -d ' ' "${temp_files[@]}"
