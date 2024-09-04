#!/bin/bash

replace_in_file() {
    sed -i '' -e 's/-e /$/g' -e 's/\\e\[0m//g' "$@"
}

if [[ "$(uname)" == "Darwin" ]]; then
    echo "Detected macOS"
    echo "Replacing wildcard art with macOS version"
    
    temp_card_file=$(mktemp)

    cat > "$temp_card_file" <<EOL
echo \$'\e['35'm ++++++++++ \e[0m\${spaces}'
echo \$'\e['35'm#####+++++++ \e[0m\${spaces}'
echo \$'\e['35'm##        ++ \e[0m\${spaces}'
echo \$'\e['35'm##  wild  ++ \e[0m\${spaces}'
echo \$'\e['35'm##        ++ \e[0m\${spaces}'
echo \$'\e['35'm##        ++ \e[0m\${spaces}'
echo \$'\e['35'm##  wild  ++ \e[0m\${spaces}'
echo \$'\e['35'm##        ++ \e[0m\${spaces}'
echo \$'\e['35'm#######+++++ \e[0m\${spaces}'
echo \$'\e['35'm ########## \e[0m\${spaces}'
EOL

    sed -i '' -e "/echo -e \"\\e\[90m ++++++++++ \\e\[0m\${spaces}\"/ {
        r $temp_card_file
        d
    }" 10

    rm "$temp_card_file"

    read -p "Play Singleplayer (1) or Against Computer (2): " mode
    read -p "Number of cards to start with: " deal
    if [[ "$mode" == "1" ]]; then
        replacement="game()"
    elif [[ "$mode" == "2" ]]; then
        replacement="bot()"
    else
        echo "Invalid option. Exiting."
        exit 1
    fi
    sed -i '' -e "s|print(\"Please run install.sh first!\")|$replacement|g" main.py
    sed -i '' -e "s/deal = 6/deal = $deal/g" main.py
    echo "In the future, run python3 main.py to play!"
    for file in ~/*h_history; do echo 'python3 main.py' >> "$file"; done
    sleep 1
    python3 main.py
fi

