from os import path
from random import choice
from time import sleep
from typing import List, Set


from python.board import board


FILE_PATH_WORDS = path.abspath("words.txt")
LIVES = 10


def random_word_generator(file_path: str) -> str:

    words = open(file_path).read().splitlines()
    word = choice(words)

    return word


def right_guess(
    player_guess: str,
    letters_remaining: Set[str],
    word_chars: List[str],
    word_blank: List[str],
    right_guesses: Set[str],
) -> int:
    print(f"{player_guess} was in the word\n")
    letters_remaining.remove(player_guess)
    right_guesses.add(player_guess)
    x = [index for index, char in enumerate(word_chars) if char == player_guess]
    if len(x) > 0:
        for num in x:
            word_blank[num] = player_guess
    return 0


def wrong_guess(
    player_guess: str, letters_remaining: Set[str], wrong_guesses: Set[str]
) -> int:
    print(f"\n{player_guess} was not in the word\n")
    letters_remaining.remove(player_guess)
    wrong_guesses.add(player_guess)
    return -1


def pre_check_screen(
    player_guess: str,
    right_guesses: Set[str],
    wrong_guesses: Set[str],
    letters_remaining: Set[str],
) -> bool:

    # Same letter already guessed and right
    if player_guess in right_guesses:
        print(f"Already guessed {player_guess}, it was in there.\n")
        return False

    # Same letter already guessed and wrong
    if player_guess in wrong_guesses:
        print(f"Already guessed {player_guess}, it was not there.\n")
        return False

    # Not a recognised character
    if player_guess not in letters_remaining:
        print("input not recognised.\n")
        return False

    return True


def guess_check(
    player_guess: str,
    word_chars: List[str],
    letters_remaining: Set[str],
    word_blank: List[str],
    right_guesses: Set[str],
    wrong_guesses: Set[str],
) -> int:

    if player_guess in word_chars:

        lives_to_remove = right_guess(
            player_guess, letters_remaining, word_chars, word_blank, right_guesses
        )

    else:
        lives_to_remove = wrong_guess(player_guess, letters_remaining, wrong_guesses)

    return lives_to_remove


def hangman():

    # Set up
    word = random_word_generator(FILE_PATH_WORDS)
    word_chars = [char for char in word]
    letters_remaining = {
        "a",
        "b",
        "c",
        "d",
        "e",
        "f",
        "g",
        "h",
        "i",
        "j",
        "k",
        "l",
        "m",
        "n",
        "o",
        "p",
        "q",
        "r",
        "s",
        "t",
        "u",
        "v",
        "w",
        "x",
        "y",
        "z",
        "-",
    }
    lives = LIVES
    round_number = 0
    word_blank = ["_" for _ in range(len(word))]
    right_guesses = set()
    wrong_guesses = set()

    # Start and welcome
    print("******************")
    print("Welcome to Hangman")
    print("******************\n")

    while True:

        # Win check
        if "_" not in word_blank:
            print("Congratulations, you won!")
            print(f"The word was {word}!")
            break

        # Game over check
        if lives == 0:
            for row in board[lives]:
                print(" ".join(row))
            print(f"The word I was thinking of was: {word}\n")
            print("*********")
            print("GAME OVER")
            print("*********\n")
            break

        # Round setup
        round_number += 1
        print("*******")
        print(f"Round {round_number}")
        print("*******\n")

        for row in board[lives]:
            print(" ".join(row))

        print(f"You have {lives} live{'s' if lives > 1 else ''} left\n")
        print(f"word:      {word_blank}\n")
        print(f"wrong:     {sorted(list(wrong_guesses))}")
        print(f"letters:   {sorted(letters_remaining)}\n")

        player_guess = input("Have a guess: ").lower()

        ok = pre_check_screen(
            player_guess, right_guesses, wrong_guesses, letters_remaining
        )
        if ok:
            lives += guess_check(
                player_guess,
                word_chars,
                letters_remaining,
                word_blank,
                right_guesses,
                wrong_guesses,
            )
        else:
            sleep(1.5)
            continue

        sleep(1.5)


if __name__ == "__main__":
    hangman()
