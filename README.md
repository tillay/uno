# UNO GAME
This is a simple uno game for a linux terminal.

Made for my AP compsci class but i still maintain and improve it

I recently rewrote it completley in go (instead of the original python) because im better at code now

The only requirements is a terminal that accepts ansii escape sequences and a computer that can run go.

Installation:

```bash
git clone https://github.com/tillay/uno
cd uno
go run uno.go
```

Modification:
to change number of starting cards, type of game (singleplayer or against computer), and some other stuff, change the instance variables listed under settings

i recommend keeping width at 10, its a lot easier if youre best with base10

im working on a 2 player version with 2 ppl (networking is goofy to figure out tho)
