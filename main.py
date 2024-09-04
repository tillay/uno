import os, random
cards = []
ecards = []
rcards = []

############# INSTANCE VARAIBLES
deal = 9    # num cards to start with
width = 10  # width per row, recommended 10 (number labels get wonky after 10)
#############

def printarray(start, end, array): #print cards in *array* from *start* to *end*
    print()
    selected = array[(2 * start) - 2:(2 * end)]
    values = ' '.join(map(str, selected))
    os.system(f'bash cards.sh {values}')
def ethink(pcard): #main computer logic function
    intersect = -1

    for i in range(int(len(rcards)/2)):
        if rcards[i*2] == pcard[0] or rcards[i*2 + 1] == pcard[1]:
            intersect = int(i) #find a card that lines up
    if intersect != -1:
        pcard[0] = rcards[intersect*2] #set middle card to computer's played card
        pcard[1] = rcards[intersect*2+1]

        rcards.pop(intersect*2) #remove real cards
        rcards.pop(intersect*2)
        ecards.pop(0) #remove fake cards
        ecards.pop(0)

    else: #computer draws a card
        ecards.extend(["e", 90])
        rcards.extend([random.randint(0, 9), random.randint(31, 34)])
        print("computer drew a card")
def rmcard(num): #remove a singular card
    cards.pop(2 * num)
    cards.pop(2 * num)
def randcard(): #draw a random card for user
    cards.append(random.randint(0, 10))
    cards.append(random.randint(31, 34))
def wildask():
    while(True):
        color = input("New color: ")
        if color in ["red", "r"]:
            out = ["e",31]
            break
        elif color in ["green", "g"]:
            out = ["e",32]
            break
        elif color in ["yellow", "y"]:
            out = ["e",33]
            break
        elif color in ["blue", "b"]:
            out = ["e",34]
            break
        if (1 <= int(color) <= len(cards) / 2):
                out = ["e",cards[2 * (int(color) - 1)+1]]
                break
        else:
            print("Invalid color!")
    return out
def printcards(array,wid): #print all cards in order with nice rows
    i = 1
    while i <= (len(array) // 2):         # go through every value in the card array, printing
        printarray(i, i + wid - 1, array) # and then incrementing by the width of each row
        i += wid
def game(): #single player main function
    os.system("clear")
    pcard = [random.randint(0, 9), random.randint(31, 34)] #set center card to random
    for i in range(deal): # give user cards
        randcard()
    while len(cards) > 0:
        printarray(0,1,pcard)
        for i in range(min(width, int(len(cards) / 2))): #display nice little numbers over the cards
            print(i + 1, end="              ")
        print()
        printcards(cards,width)
        selec = input("Enter your choice: ")
        try:
            selec = int(selec) #check if selected card is valid to play
            if 1 <= selec <= len(cards) / 2:
                if (cards[2 * (selec - 1)] == pcard[0] or cards[2 * (selec - 1) + 1] == pcard[1]):
                    pcard = [cards[2 * (selec - 1)], cards[2 * (selec - 1) + 1]]
                    rmcard(selec - 1)
                    if (pcard[0] == 10):
                        pcard=(wildask())
                    os.system("clear")
                elif cards[2 * (selec - 1)] == 10: #wildcard logic
                    pcard=(wildask())
                    os.system("clear")
                    rmcard(selec - 1)
                else:
                    os.system("clear")
                    print("Invalid input!")
            else:
                os.system("clear")
                print("Invalid input!")
        except ValueError: # dark magic logic from stack overflow
            if selec == "d" or selec == "": # code to draw a card
                randcard()
                os.system("clear")
            else:
                os.system("clear")
                print("Invalid input!")
    os.system("figlet YOU WIN!")
    os.system("neofetch")
    print()


def bot(): #basically singleplayer with some tweaks and computer logic added
    os.system("clear")
    for i in range(deal):
        ecards.extend(["e", 90])
        rcards.extend([random.randint(0, 9), random.randint(31, 34)])
    pcard = [random.randint(0, 9), random.randint(31, 34)]
    for i in range(deal):
        randcard()
    while len(cards) > 0 and len(ecards) > 0: # Main loop
        ethink(pcard)  # compooter thinky time
        printcards(ecards,width) #computer prints cards
        printarray(0,1,pcard)
        for i in range(min(width, int(len(cards) / 2))):
            print(i + 1, end="               ")
        print()
        printcards(cards,width)
        selec = input("Enter your choice: ")
        try:
            selec = int(selec)
            if 1 <= selec <= len(cards) / 2:
                if (cards[2 * (selec - 1)] == pcard[0] or cards[2 * (selec - 1) + 1] == pcard[1]):
                    pcard = [cards[2 * (selec - 1)], cards[2 * (selec - 1) + 1]]
                    rmcard(selec - 1)
                    os.system("clear")
                    if (pcard[0] == 10):
                        pcard=(wildask())
                elif cards[2 * (selec - 1)] == 10:
                    pcard=(wildask())
                    rmcard(selec - 1)
                else:
                    os.system("clear")
                    print("Invalid input!")
            else:
                os.system("clear")
                print("Invalid input!")
        except ValueError:
            if selec == "d" or selec == "":
                randcard()
                os.system("clear")
            else:
                os.system("clear")
                print("Invalid input!")
    if len(cards) == 0:
        os.system("figlet YOU WIN!")
        os.system("neofetch")
    else:
        os.system("figlet YOU LOSE")
        for i in range( 0):
            os.system("figlet BRUH")
    print()  
game()
