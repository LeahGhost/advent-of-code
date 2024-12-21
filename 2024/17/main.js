const fs = require("fs");

const data = fs.readFileSync('./input.txt', 'utf-8').trim();

function parseData(data) {
  const [registersData, instructionsData] = data.split('\n\n');
  const [X, Y, Z] = registersData.match(/\d+/g).map(Number);
  const instructionsList = instructionsData.split(' ')[1].split(',').map(Number);
  return { X, Y, Z, instructionsList };
}

function runInstructions({ X, Y, Z, instructionsList }) {
  let Xval = X;
  let Yval = Y;
  let Zval = Z;
  const output = [];

  const getValue = (param) => {
    switch (param) {
      case 0: case 1: case 2: case 3: return param;
      case 4: return Xval;
      case 5: return Yval;
      case 6: return Zval;
    }
  };

  for (let i = 0; i < instructionsList.length; i += 2) {
    const command = instructionsList[i];
    const argument = instructionsList[i + 1];
    
    switch (command) {
      case 0:
        Xval = Math.trunc(Xval / Math.pow(2, getValue(argument)));
        break;
      case 1:
        Yval ^= argument;
        break;
      case 2:
        Yval = getValue(argument) % 8;
        break;
      case 3:
        if (Xval === 0) break;
        i = argument - 2;
        break;
      case 4:
        Yval ^= Zval;
        break;
      case 5:
        output.push(getValue(argument) % 8);
        break;
      case 6:
        Yval = Math.trunc(Xval / Math.pow(2, getValue(argument)));
        break;
      case 7:
        Zval = Math.trunc(Xval / Math.pow(2, getValue(argument)));
        break;
    }
  }

  return output;
}

function generateOutput({ X, Y, Z, instructionsList }) {
  let Xval = X;
  let Yval = Y;
  let Zval = Z;
  const output = [];

  const getValue = (param) => {
    switch (param) {
      case 0n: case 1n: case 2n: case 3n: return param;
      case 4n: return Xval;
      case 5n: return Yval;
      case 6n: return Zval;
    }
  };

  for (let i = 0; i < instructionsList.length; i += 2) {
    const command = instructionsList[i];
    const argument = instructionsList[i + 1];
    
    switch (command) {
      case 0n:
        Xval = Xval / BigInt(Math.pow(2, Number(getValue(argument))));
        break;
      case 1n:
        Yval ^= argument;
        break;
      case 2n:
        Yval = getValue(argument) % 8n;
        break;
      case 3n:
        if (Xval === 0n) break;
        i = Number(argument) - 2;
        break;
      case 4n:
        Yval ^= Zval;
        break;
      case 5n:
        output.push(getValue(argument) % 8n);
        break;
      case 6n:
        Yval = Xval / BigInt(Math.pow(2, Number(getValue(argument))));
        break;
      case 7n:
        Zval = Xval / BigInt(Math.pow(2, Number(getValue(argument))));
        break;
    }
  }

  return output;
}

function solvePart1({ X, Y, Z, instructionsList }) {
  const output = runInstructions({ X, Y, Z, instructionsList });
  return output.join(',');
}

function solvePart2({ Y, Z, instructionsList }) {
  instructionsList = instructionsList.map(BigInt);
  let possibleX = [0n];

  for (let i = instructionsList.length - 1; i >= 0; i--) {
    const newPossibleX = [];

    for (const X of possibleX) {
      for (let rem = 0n; rem < 8n; rem++) {
        const output = generateOutput({ X: X * 8n + rem, Y, Z, instructionsList });
        if (output[0] === instructionsList[i]) {
          newPossibleX.push(X * 8n + rem);
        }
      }
    }

    possibleX = newPossibleX;
  }

  return Math.min(...possibleX.map(Number));
}

console.log(solvePart1(parseData(data)));
console.log(solvePart2(parseData(data)));