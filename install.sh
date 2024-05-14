#!/bin/bash

replace_in_files() {
    sed -i '' -e 's/-e /$/g' -e 's/\\e\[0m//g' "$@"
    echo "Modified $@ for macOS"
}

if [[ "$(uname)" == "Darwin" ]]; then
    echo "Detected macOS"
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
    brew install figlet
    brew install neofetch
    replace_in_files 0 1 2 3 4 5 6 7 8 9 10 e
        cat > 10 <<EOL
echo \$'\e['35'm ++++++++++'
echo \$'\e['35'm#####+++++++'
echo \$'\e['35'm##        ++'
echo \$'\e['35'm##  wild  ++'
echo \$'\e['35'm##        ++'
echo \$'\e['35'm##        ++'
echo \$'\e['35'm##  wild  ++'
echo \$'\e['35'm##        ++'
echo \$'\e['35'm#######+++++'
echo \$'\e['35'm ##########'
EOL
elif [ -x "$(command -v apt)" ]; then
    echo "Detected Debian-based distribution"
    sudo apt install figlet
    sudo apt install neofetch
elif [ -x "$(command -v dnf)" ]; then
    echo "Detected Fedora-based distribution"
    sudo dnf install figlet
    sudo dnf install neofetch
elif [ -x "$(command -v pacman)" ]; then
    echo "Detected Arch-based distribution"
    sudo pacman -S figlet
    sudo pacman -S neofetch
else
    echo "Unsupported distribution or package manager not found."
fi
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
chmod u+x 0 1 2 3 4 5 6 7 8 9 10 e cards.sh
