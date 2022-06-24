import { readFileSync } from 'fs';


const FILE_PATH_WORDS: string = "words.txt"
let LIVES: Number = 10


function randomWordGenerator(filename: string): string {
  const contents = readFileSync(filename, 'utf-8');

  const words: string[] = contents.split(/\r?\n/);

  var word: string = words[Math.floor(Math.random() * words.length)];

  return word;
}


function hangman() {
  let word: string = randomWordGenerator(FILE_PATH_WORDS)
  let wordChars: string[] = word.split("")
  let lettersRemaining: string[] = [
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
  ]
  var lives: Number = LIVES
  var roundNumber: Number = 0
  var wordBlank = []

  // in lieu of a list comprehension 
  for (let i =0; i < word.length; i++){
    wordBlank.push('_')
  }

  var rightGuess: string[] = []
  var wrongGuess: string[] = []

  // Start and welcome
  console.log("******************")
  console.log("Welcome to Hangman")
  console.log("******************\n")
  




  console.log(word)
  console.log(wordChars)
  console.log(lettersRemaining)
  console.log(lives)
  console.log(roundNumber)
  console.log(wordBlank)
  console.log(rightGuess)
  console.log(wrongGuess)

}


if (require.main === module) {
  hangman();
}
