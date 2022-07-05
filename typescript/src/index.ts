import { readFileSync } from 'fs';


const FILE_PATH_WORDS: string = "../../words.txt"
var LIVES: number = 10



function playerGuesser(): string {

  var letters: string = "abcdefghijklmnopqrstuvwxyz-"

  var guess: string = letters[Math.floor(Math.random() * letters.length)];
  
  return guess
}


function randomWordGenerator(filename: string): string {
  const contents = readFileSync(filename, 'utf-8');

  const words: string[] = contents.split(/\r?\n/);

  var word: string = words[Math.floor(Math.random() * words.length)];

  return word;
}


function rightGuess(playerGuess: string, lettersRemaining: string[], wordChars: string[], wordBlank: string[], rightGuesses: string[]): number {

  console.log("%s was in the word\n", playerGuess)
  var indexGuess = lettersRemaining.indexOf(playerGuess)
  lettersRemaining[indexGuess] = "_"

  rightGuesses.push(playerGuess)



  for (var i = 0; i < wordChars.length; i++) {

    if (playerGuess === wordChars[i]) {
      wordBlank[i] = playerGuess
    }
  }

  return 0
}

function wrongGuess(playerGuess: string, lettersRemaining: string[], wrongGuess: string[]): number {

  console.log("%s, was not in the word\n", playerGuess)
  var indexGuess = lettersRemaining.indexOf(playerGuess)
  lettersRemaining[indexGuess] = "_"
  wrongGuess.push(playerGuess)


  return -1
}


function preCheckGuess(
  playerGuess: string,
  rightGuesses: string[],
  wrongGuesses: string[],
  lettersRemaining: string[]
): boolean {

  // Same letter already guessed and right
  if (rightGuesses.includes(playerGuess)) {
    console.log("Already guessed %s, it was in there.\n", playerGuess)
    return false
  }

  // Same letter already guessed and wrong
  if (wrongGuesses.includes(playerGuess)) {
    console.log("Already guessed %s, it was not there.\n", playerGuess)
    return false
  }

  // Not a recognised character
  if (!lettersRemaining.includes(playerGuess)) {
    console.log("input not recognised.\n")
    return false
  }

  return true
}


function guessCheck(playerGuess: string, wordChars: string[], lettersRemaining: string[], wordBlank: string[] ,rightGuesses: string[], wrongGuesses: string[]): number {
  if (wordChars.includes(playerGuess)) {
    var livesToRemove = rightGuess(playerGuess, lettersRemaining, wordChars, wordBlank, rightGuesses)
  }
  else {
    var livesToRemove = wrongGuess(playerGuess, lettersRemaining, wrongGuesses)
  }

  return livesToRemove
}


function hangman() {
  var word: string = randomWordGenerator(FILE_PATH_WORDS)
  var wordChars: string[] = word.split("")

  var lettersRemaining = "abcdefghijklmnopqrstuvwxyz".split("")
  var lives: number = LIVES
  var roundNumber: number = 1
  var wordBlank = []

  // in lieu of a list comprehension 
  for (let i = 0; i < word.length; i++) {
    wordBlank.push('_')
  }

  var rightGuesses: string[] = []
  var wrongGuesses: string[] = []

  // Start and welcome
  console.log("***************************************")
  console.log("Welcome to Hangman Written in TypeScript")
  console.log("****************************************\n")


  while (true) {

    // Win check
    if (!wordBlank.includes("_")) {
      console.log("\n\n\n\n")
      console.log("*************************")
      console.log("Congratulations, you won!")
      console.log("*************************\n")
      console.log("The word was %s!", word)
      break
    }

    // Game over check
    if (lives === 0) {
      // todo board display
      console.log("\n\n\n\n")
      console.log("*********")
      console.log("GAME OVER")
      console.log("*********\n")
      console.log("The word I was thinking of was: %s\n", word)
      break
    }

    // Round setup
    console.log("********")
    console.log("Round %s", roundNumber)
    console.log("********\n")
    roundNumber += 1

    // Dashboard
    console.log("You have %i live(s)\n", lives)
    console.log("word:        %s\n", wordBlank.join(" "))
    console.log("worng:       %s\n", wrongGuesses.sort().join(' '))
    console.log("letters:     %s\n", lettersRemaining.sort().join(' '))


    // var playerGuess: string = prompt("What is your name?")
    var playerGuess: string = playerGuesser()

    var okToPlay: boolean = preCheckGuess(playerGuess, rightGuesses, wrongGuesses, lettersRemaining)

    if (okToPlay) {
      
      lives += guessCheck(
        playerGuess, wordChars, lettersRemaining, wordBlank, rightGuesses, wrongGuesses
      )

      }

  }
}


if (require.main === module) {
  hangman();
}
