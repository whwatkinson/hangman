use std::fs;


const FILE_PATH_WORDS: &str = "../words.txt";
const LIVES: i32 = 10;



fn random_word_generator(file_path: &str) -> &str {
    let data = fs::read_to_string(file_path).expect("Unable to read file");
    println!("{}", data);
    return "asd"
}

fn right_guess(
    player_guess: &char,
    letters_remaining: &[char],
    word_chars: &[char],
    word_blank: &[char],
    right_guesses: &[char]
) -> i32 {
    return 0
}

fn wrong_guess(
    player_guess: &char,
    letters_remaining: &[char],
    wrong_guesses: &[char]
) -> i32 {
    return -1
}

fn pre_check_screen(
    player_guess: &char,
    right_guesses: &[char],
    wrong_guesses: &[char]

) -> bool {
    return true
}

fn guess_check(

    player_guess: &char,
    word_chars: &[char],
    letters_remaining: &[char],
    word_blank: &[char],
    right_guesses: &[char],
    wrong_guesses: &[char]
) -> i32 {
    return 0
}

fn main() {
    random_word_generator(FILE_PATH_WORDS);
}