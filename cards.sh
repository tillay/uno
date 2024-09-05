#!/bin/bash

draw_card() {
    local card_num=$1
    local color=$2
    local spaces="  "

    case "$card_num" in
        0)
            echo -e "\033[0;${color}m ++++++++++ \033[0m${spaces}"
            echo -e "\033[0;${color}m#####+++++++\033[0m${spaces}"
            echo -e "\033[0;${color}m##      0 ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##   __   ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##  |  |  ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##  |__|  ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##        ++\033[0m${spaces}"
            echo -e "\033[0;${color}m## 0      ++\033[0m${spaces}"
            echo -e "\033[0;${color}m#######+++++\033[0m${spaces}"
            echo -e "\033[0;${color}m ########## \033[0m${spaces}"
            ;;
        1)
            echo -e "\033[0;${color}m ++++++++++ \033[0m${spaces}"
            echo -e "\033[0;${color}m#####+++++++\033[0m${spaces}"
            echo -e "\033[0;${color}m##      1 ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##   /|   ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##    |   ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##   _|_  ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##        ++\033[0m${spaces}"
            echo -e "\033[0;${color}m## 1      ++\033[0m${spaces}"
            echo -e "\033[0;${color}m#######+++++\033[0m${spaces}"
            echo -e "\033[0;${color}m ########## \033[0m${spaces}"
            ;;
        2)
            echo -e "\033[0;${color}m ++++++++++ \033[0m${spaces}"
            echo -e "\033[0;${color}m######++++++\033[0m${spaces}"
            echo -e "\033[0;${color}m##      2 ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##   __   ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##   __|  ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##  |__   ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##        ++\033[0m${spaces}"
            echo -e "\033[0;${color}m## 2      ++\033[0m${spaces}"
            echo -e "\033[0;${color}m#######+++++\033[0m${spaces}"
            echo -e "\033[0;${color}m ########## \033[0m${spaces}"
            ;;
        3)
            echo -e "\033[0;${color}m ++++++++++ \033[0m${spaces}"
            echo -e "\033[0;${color}m######++++++\033[0m${spaces}"
            echo -e "\033[0;${color}m##      3 ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##   __   ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##   __|  ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##   __|  ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##        ++\033[0m${spaces}"
            echo -e "\033[0;${color}m## 3      ++\033[0m${spaces}"
            echo -e "\033[0;${color}m#######+++++\033[0m${spaces}"
            echo -e "\033[0;${color}m ########## \033[0m${spaces}"
            ;;
        4)
            echo -e "\033[0;${color}m ++++++++++ \033[0m${spaces}"
            echo -e "\033[0;${color}m######++++++\033[0m${spaces}"
            echo -e "\033[0;${color}m##      4 ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##        ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##   |_|  ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##     |  ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##        ++\033[0m${spaces}"
            echo -e "\033[0;${color}m## 4      ++\033[0m${spaces}"
            echo -e "\033[0;${color}m#######+++++\033[0m${spaces}"
            echo -e "\033[0;${color}m ########## \033[0m${spaces}"
            ;;
        5)
            echo -e "\033[0;${color}m ++++++++++ \033[0m${spaces}"
            echo -e "\033[0;${color}m######++++++\033[0m${spaces}"
            echo -e "\033[0;${color}m##      5 ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##   __   ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##  |__   ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##   __|  ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##        ++\033[0m${spaces}"
            echo -e "\033[0;${color}m## 5      ++\033[0m${spaces}"
            echo -e "\033[0;${color}m#######+++++\033[0m${spaces}"
            echo -e "\033[0;${color}m ########## \033[0m${spaces}"
            ;;
        6)
            echo -e "\033[0;${color}m ++++++++++ \033[0m${spaces}"
            echo -e "\033[0;${color}m######++++++\033[0m${spaces}"
            echo -e "\033[0;${color}m##      6 ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##   __   ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##  |__   ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##  |__|  ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##        ++\033[0m${spaces}"
            echo -e "\033[0;${color}m## 6      ++\033[0m${spaces}"
            echo -e "\033[0;${color}m#######+++++\033[0m${spaces}"
            echo -e "\033[0;${color}m ########## \033[0m${spaces}"
            ;;
        7)
            echo -e "\033[0;${color}m ++++++++++ \033[0m${spaces}"
            echo -e "\033[0;${color}m######++++++\033[0m${spaces}"
            echo -e "\033[0;${color}m##      7 ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##   __   ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##    /   ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##   /    ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##        ++\033[0m${spaces}"
            echo -e "\033[0;${color}m## 7      ++\033[0m${spaces}"
            echo -e "\033[0;${color}m#######+++++\033[0m${spaces}"
            echo -e "\033[0;${color}m ########## \033[0m${spaces}"
            ;;
        8)
            echo -e "\033[0;${color}m ++++++++++ \033[0m${spaces}"
            echo -e "\033[0;${color}m######++++++\033[0m${spaces}"
            echo -e "\033[0;${color}m##      8 ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##   __   ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##  |__|  ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##  |__|  ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##        ++\033[0m${spaces}"
            echo -e "\033[0;${color}m## 8      ++\033[0m${spaces}"
            echo -e "\033[0;${color}m#######+++++\033[0m${spaces}"
            echo -e "\033[0;${color}m ########## \033[0m${spaces}"
            ;;
        9)
            echo -e "\033[0;${color}m ++++++++++ \033[0m${spaces}"
            echo -e "\033[0;${color}m######++++++\033[0m${spaces}"
            echo -e "\033[0;${color}m##      9 ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##   __   ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##  |__|  ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##   __|  ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##        ++\033[0m${spaces}"
            echo -e "\033[0;${color}m## 9      ++\033[0m${spaces}"
            echo -e "\033[0;${color}m#######+++++\033[0m${spaces}"
            echo -e "\033[0;${color}m ########## \033[0m${spaces}"
            ;;
        10)
	    echo -e "\033[0;90m ++++++++++ \033[0;0m${spaces}"
	    echo -e "\033[0;90m#####+++++++\033[0;0m${spaces}"
	    echo -e "\033[0;90m##\033[0;31mXXXX\033[0;34mXXXX\033[0;90m++\033[0;0m${spaces}"
	    echo -e "\033[0;90m##\033[0;31mXXXX\033[0;34mXXXX\033[0;90m++\033[0;0m${spaces}"
	    echo -e "\033[0;90m##\033[0;31mXXXX\033[0;34mXXXX\033[0;90m++\033[0;0m${spaces}"
	    echo -e "\033[0;90m##\033[0;33mXXXX\033[0;32mXXXX\033[0;90m++\033[0;0m${spaces}"
	    echo -e "\033[0;90m##\033[0;33mXXXX\033[0;32mXXXX\033[0;90m++\033[0;0m${spaces}"
	    echo -e "\033[0;90m##\033[0;33mXXXX\033[0;32mXXXX\033[0;90m++\033[0;0m${spaces}"
	    echo -e "\033[0;90m#######+++++\033[0;0m${spaces}"
	    echo -e "\033[0;90m ########## \033[0;0m${spaces}"
            ;;
        e)
            echo -e "\033[0;${color}m ++++++++++ \033[0m${spaces}"
            echo -e "\033[0;${color}m######++++++\033[0m${spaces}"
            echo -e "\033[0;${color}m##        ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##        ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##        ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##        ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##        ++\033[0m${spaces}"
            echo -e "\033[0;${color}m##        ++\033[0m${spaces}"
            echo -e "\033[0;${color}m#######+++++\033[0m${spaces}"
            echo -e "\033[0;${color}m ########## \033[0m${spaces}"
            ;;
        *)
            echo "Invalid card number."
            ;;
    esac
}

trap 'rm -f "${temp_files[@]}"' EXIT

while [ "$#" -gt 0 ]; do
    card_num=$1
    card_color=$2
    temp_file=$(mktemp)
    draw_card "$card_num" "$card_color" > "$temp_file"
    temp_files+=("$temp_file")
    shift 2
done

paste -d ' ' "${temp_files[@]}"
