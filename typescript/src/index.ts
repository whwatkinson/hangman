import {readFileSync} from 'fs';


const FILE_PATH_WORDS: String = "words.txt"
let LIVES: Number = 10


function randomWordGenerator(filename: string)  {
    const contents = readFileSync(filename, 'utf-8');
  
    const words = contents.split(/\r?\n/);

    var word = words[Math.floor(Math.random()*words.length)];
  
    console.log(word);
  
    return word;
  }


function hangman(){
  let word: string = randomWordGenerator("words.txt")
}


if (require.main === module) {
  hangman();
}
