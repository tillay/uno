# UNO GAME
This is a simple uno game for a linux terminal.

The only requirments are a computer that can run python and bash, and a terminal that accepts ansii escape sequences.

Installation:

git clone https://github.com/tillay/uno

cd into it and install:

cd uno

chmod +x install.sh

Then:
./install.sh

to play:
python3 main.py

To play in the future, make sure you are in the correct directory and type the same command.

Modification:
there are two varaibles that can be changed in the main.py:
deal, the number of cards dealt at the start of the game
and width, the width of each column.
To switch from singleplayer to against the computer, change the last line from bot() to game() and vise versa.
