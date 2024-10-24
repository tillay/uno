# UNO GAME
This is a simple uno game for a linux terminal.

Made for my AP compsci class but i still maintain and improve it

The only requirements are a computer that can run python and bash, and a terminal that accepts ansii escape sequences.

Installation:

`git clone https://github.com/tillay/uno&&cd uno&&./install.sh`

To play in the future, make sure you are in the uno directory and run: `python3 main.py`

Modification:
to change number of starting cards and type of game (singleplayer or against computer) run the installer script again. (`./install.sh`)

to change the width of each column, modify the instance variable with `width = 10`. Not reccomended unless you dont prefer base 10

(numbers also get kinda offset too)

I didnt add AI at all, the computer version is programmed to only play valid cards with no strategy

im working on a 2 player version with 2 ppl on same ssh server but its a buncha work and confusing

Known issues: 
- installer does not properly change number of starting cards on macOS
- windows support is nonexistent (barring WSL)
- cards script can be optimized (loads of repeated code rn)
- user must be cd'ed in the directory with both the main.py and cards.sh, abselute filepaths don't work
