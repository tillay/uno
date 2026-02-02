# UNO GAME
This is a simple uno game for a linux terminal.

I first made this for my AP compsci class, but I still maintain and improve it.

I recently rewrote it completely in go (instead of the original python) because im better at code now.

The only requirements is a terminal that accepts ansi escape sequences and a computer that can run go.

im working on a 2 player version with 2 ppl (networking is goofy to figure out tho)

### Installation:
make sure go is installed on your system (`sudo apt install golang-go` on debian and `sudo pacman -S go` on arch)
```bash
git clone https://github.com/tillay/uno
cd uno
go run uno.go
```

### Gameplay
You can play against either computer or yourself.

To play, type the index of the card you want to play and press enter.

To draw a card, press enter without typing anything first. 

If you play a wild card, type the first letter of the color you want (`r`,`g`,`y`,`b`) or the index of a card that is the color you want

The prompt being `>` means you are supposed to play a card (it expects a number) and if its `->` then that means you played a wildcard, and it expects a color letter or index.

### Modification:
to change number of starting cards, gamemode, and some other stuff, change the instance variables listed under settings

line width is how many cards it puts on each line (before carrying over to next one)

init cards is how many cards it initially serves to each player

against ai is whether its singleplayer or against the computer (ai)

enable hints controls whether to underline index labels for cards that can be played

debugging mode shows the computer's cards and doesn't clear the screen every turn
