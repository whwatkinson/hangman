from random import choice, randint
from os import path
import random

FILE_PATH_WORDS = path.abspath('words.txt')


def random_word_generator(file_path: str) -> str:

    words = open(file_path).read().splitlines()
    word = choice(words)

    return word


def hangman():

    # Set up
    word = random_word_generator(FILE_PATH_WORDS)
    word2 = []
    letters = {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '-'}
    lives = 10
    round_number = 0
    word_blank = ['_' for _ in range(len(word))]
    right_guess = set()
    wrong_guess = set()

    for char in word:
        word2.append(char)

    # Start and welcome
    print('\n******************')
    print('Welcome to Hangman')
    print('******************')

    playing = True

    while playing:
        print('\nI have thought of a word, can you guess it?\n')
        print('You have 10 lives good luck!\n')

        # Player guesses
        while True:

            # Win check
            if '_' not in word_blank:
                print('Congratulations, you won!')
                playing = False
                break

            # Game over check
            elif lives == 0:
                print('The word I was thinking of was: \n' + ''.join(word2))
                print('\n********')
                print('GAME OVER')
                print('********\n')
                playing = False
                break

            # Round setup
            round_number += 1
            print('\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n*******')
            print('Round ' + str(round))
            print('*******\n')
            print(word_blank)
            print(letters)
            print('You have ' + str(lives) + ' live(s) left\n')
            guess = input('\nHave a guess: \n').lower()

            # Same letter already guessed and right
            if guess in right_guess:
                print('Already guessed, it was in there.\n')
                continue

            # Same letter already guessed and wrong
            elif guess in wrong_guess:
                print('Already guessed, it wasnt there.\n')
                continue

            # Not a recognised character
            if guess not in letters:
                print('Computer says noo.')
                continue

            # Right guess
            elif guess in word2:
                letters.remove(guess)
                right_guess.add(guess)
                print('Well done!\n')

                # Adding in the correct letter to blank guess
                x = [pos for pos, char in enumerate(word2) if char == guess]
                if len(x) > 0:
                    for num in x:
                        word_blank[num] = guess

                # Display board and remaining letters
                print(str(word_blank))
                print(str(letters) + '\n')


            # Wrong Guess
            elif guess not in word2:
                lives -= 1
                letters.remove(guess)
                wrong_guess.add(guess)
                print('\n' + guess + ' was not in the word\n')
                print(str(word_blank) + '\n')
                print(str(letters) + '\n')

    print('\n*******')
    print('Goodbye')
    print('*******')

hangman()
