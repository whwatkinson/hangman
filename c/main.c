#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

//from chat gpt

// Function to get a random word from a list
char* get_word() {
    char* word_list[] = {"apple", "banana", "orange", "grape", "melon"};
    int list_size = sizeof(word_list) / sizeof(word_list[0]);
    srand(time(NULL));
    int index = rand() % list_size;
    return word_list[index];
}

int main() {
    char* word = get_word(); // Get a random word from the list
    int word_length = strlen(word);
    char guessed_letters[26] = {0}; // Array to track guessed letters
    int num_guesses = 0;
    int max_guesses = 6; // Number of maximum incorrect guesses allowed

    printf("Welcome to Hangman!\n");
    printf("The word has %d letters. You have %d chances to guess.\n", word_length, max_guesses);

    while (num_guesses < max_guesses) {
        printf("\n");
        for (int i = 0; i < word_length; i++) {
            if (guessed_letters[word[i] - 'a']) {
                printf("%c ", word[i]); // Print guessed letters
            } else {
                printf("_ ");
            }
        }

        char guess;
        printf("\nEnter a letter: ");
        scanf(" %c", &guess);
        if (guessed_letters[guess - 'a']) {
            printf("You already guessed that letter!\n");
        } else {
            guessed_letters[guess - 'a'] = 1;
            if (strchr(word, guess) != NULL) {
                printf("Good guess!\n");
            } else {
                printf("Incorrect guess!\n");
                num_guesses++;
            }
        }

        int num_letters_guessed = 0;
        for (int i = 0; i < 26; i++) {
            if (guessed_letters[i]) {
                num_letters_guessed++;
            }
        }
        if (num_letters_guessed == word_length) {
            printf("\nCongratulations, you guessed the word %s!\n", word);
            return 0;
        }
    }

    printf("\nSorry, you ran out of chances. The word was %s.\n", word);
    return 0;
}
