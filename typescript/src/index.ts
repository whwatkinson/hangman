// in browser
import {readFileSync} from 'fs';

// out of browser
// const {readFileSync, promises: fsPromises} = require('fs');


const FILE_PATH_WORDS: String = "words.txt"
// let LIVES: Number = 10


function syncReadFile(filename: string)  {
    const contents = readFileSync(filename, 'utf-8');
  
    const words = contents.split(/\r?\n/);

    var word = words[Math.floor(Math.random()*words.length)];

  
    console.log(word);
  
    return word;
  }

  syncReadFile("words.txt")
