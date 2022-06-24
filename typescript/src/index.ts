// in browser
// import {readFileSync} from 'fs';

// out of browser
const {readFileSync, promises: fsPromises} = require('fs');


const FILE_PATH_WORDS: String = "words.txt"
// let LIVES: Number = 10


function syncReadFile(filename: string)  {
    const contents = readFileSync(filename, 'utf-8');
  
    const arr = contents.split(/\r?\n/);
  
    console.log(arr);
  
    return arr;
  }

  syncReadFile("words.txt")