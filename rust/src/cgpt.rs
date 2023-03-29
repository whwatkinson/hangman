use std::io;
use rand::Rng;

// Function to get a random word from a list
fn get_word() -> &'static str {
    let word_list = ["apple", "banana", "orange", "grape", "melon"];
    let mut rng = rand::thread_rng();
    let index = rng.gen_range(0..word_list.len());
    return word_list[index];
}

fn main() {
    let word = get_word(); // Get a random word from the list
    let word_length = word.len();
    let mut guessed_letters = vec![false; 26]; // Vector to track guessed letters
    let mut num_guesses = 0;
    let max_guesses = 6; // Number of maximum incorrect guesses allowed

    println!("Welcome to Hangman!");
    println!("The word has {} letters. You have {} chances to guess.", word_length, max_guesses);

    while num_guesses < max_guesses {
        println!("");
        for (i, c) in word.chars().enumerate() {
            if guessed_letters[c as usize - 'a' as usize] {
                print!("{} ", c); // Print guessed letters
            } else {
                print!("_ ");
            }
        }

        let mut guess = String::new();
        println!("\nEnter a letter: ");
        io::stdin().read_line(&mut guess).expect("Failed to read line");
        let guess = guess.trim().chars().next().unwrap();

        if guessed_letters[guess as usize - 'a' as usize] {
            println!("You already guessed that letter!");
        } else {
            guessed_letters[guess as usize - 'a' as usize] = true;
            if word.contains(guess) {
                println!("Good guess!");
            } else {
                println!("Incorrect guess!");
                num_guesses += 1;
            }
        }

        let num_letters_guessed = guessed_letters.iter().filter(|&&x| x).count();
        if num_letters_guessed == word_length {
            println!("\nCongratulations, you guessed the word {}!", word);
            return;
        }
    }

    println!("\nSorry, you ran out of chances. The word was {}.", word);
}
