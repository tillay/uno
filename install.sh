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

# Prompt user for input
read -p "Amount of cards to start with: " deal

# Replace the value in main.py
sed -i '' -e "s/deal = 6/deal = $deal/g" main.py

chmod u+x 0 1 2 3 4 5 6 7 8 9 10 e cards.sh
