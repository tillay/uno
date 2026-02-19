# UNO GAME
This is a simple uno game for a linux terminal.

I first made this for my AP compsci class, but I still maintain and improve it.

I recently rewrote it completely in go (instead of the original python) because im better at code now.

The only requirements are a terminal that accepts ansi escape sequences and a computer that can run go.

If you want to play online, an internet connection would be nice as well.

### Installation:
make sure go is installed on your system (`sudo apt install golang-go` on debian and `sudo pacman -S go` on arch)
```bash
git clone https://github.com/tillay/uno
cd uno
go build
go run uno
```
> [!IMPORTANT]
> do not run just `go run uno.go` this will *not work*. Make sure to run `go build` first!

### Gameplay
You can play against the computer, yourself, or against a friend.

To play a card, type the index (starting from 1 - there are labels) of the card you want to play, and press enter to play it.

To draw a card, press enter without typing anything first.

If you play a wild card, type the first letter of the color you want (`r`,`g`,`y`,`b`) or the index of a card that is the color you want to switch gameplay to.

A played wildcard will put a blank color card on top of the stack

The prompt being `>` means you are supposed to play a card (it expects a number) and if its `->` then that means you played a wildcard, and it expects a color letter or index.

### Modification:
There are a lot of flags that one can use to change operation. Run `./uno -help` for more infos.

### Docker Server Instructions

Note: due to some changes this is currently not dockerized

```bash
git clone https://github.com/tillay/uno
cd uno
cd server
docker compose up --build -d
```
> [!IMPORTANT]
> make sure to uncomment or change the websocket url in `uno.go` on your client device so you can use your websocket instead of default.