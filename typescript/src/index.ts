import { readFileSync } from 'fs';


const FILE_PATH_WORDS: string = "words.txt"
var LIVES: number = 10


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
  var lives: number = LIVES
  var roundNumber: number = 1
  var wordBlank = []

  // in lieu of a list comprehension 
  for (let i = 0; i < word.length; i++) {
    wordBlank.push('_')
  }

  var rightGuess: string[] = []
  var wrongGuess: string[] = []

  // Start and welcome
  console.log("******************")
  console.log("Welcome to Hangman")
  console.log("******************\n")


  while (true) {

    // Win check
    if (!wordBlank.includes("_")){
      console.log("Congratulations, you won!")
      console.log("The word was %s!", word)
      break
    }

    // Game over check
    if (lives === 0){
      // todo board display
      console.log("The word I was thinking of was: %s\n", word)
      console.log("*********")
      console.log("GAME OVER")
      console.log("*********\n")
      break
    }

    // Round setup
    console.log("*******")
    console.log("Round %s", roundNumber)
    console.log("*******\n")
    roundNumber += 1


    console.log("You have %i live(s)\n", lives)
    // console.log("You have %i live{'s' if lives > 1 else ''} left\n", lives)
    console.log("word:        %s\n", wordBlank.join(" "))
    console.log("worng:       %s\n", wrongGuess.sort().join(' '))
    console.log("letters:     %s\n", lettersRemaining.sort().join(' '))





    console.log(word)
    console.log(wordChars)
    console.log(lettersRemaining)
    console.log(lives)
    console.log(roundNumber)
    console.log(wordBlank)
    console.log(rightGuess)
    console.log(wrongGuess)

    lives -= 1
  }
}


if (require.main === module) {
  hangman();
}
