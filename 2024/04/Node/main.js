const fs = require('fs');

const directions = [
  [0, 1],
  [1, 1],
  [1, 0],
  [1, -1],
  [0, -1],
  [-1, -1],
  [-1, 0],
  [-1, 1],
];

function loadFile() {
  const filePath = __dirname + '/../input.txt';
  const data = fs.readFileSync(filePath, 'utf8');
  return data.split('\n').filter((line) => line.length > 0);
}

function isValid(row, col, numRows, numCols) {
  return row >= 0 && row < numRows && col >= 0 && col < numCols;
}

function part1(grid, numRows, numCols) {
  let count = 0;
  for (let row = 0; row < numRows; row++) {
    for (let col = 0; col < numCols; col++) {
      for (const direction of directions) {
        if (
          checkDirection(
            grid,
            row,
            col,
            direction[0],
            direction[1],
            numRows,
            numCols,
            'XMAS',
          )
        ) {
          count++;
        }
      }
    }
  }
  return count;
}

function checkDirection(
  grid,
  startX,
  startY,
  deltaX,
  deltaY,
  numRows,
  numCols,
  word,
) {
  for (let i = 0; i < word.length; i++) {
    const currRow = startX + i * deltaX;
    const currCol = startY + i * deltaY;
    if (
      !isValid(currRow, currCol, numRows, numCols) ||
      grid[currRow][currCol] !== word[i]
    ) {
      return false;
    }
  }
  return true;
}

function part2(grid, numRows, numCols) {
  let count = 0;

  function checkDirection(startX, startY, deltaX, deltaY) {
    startX = startX - deltaX;
    startY = startY - deltaY;
    let word = '';
    for (let i = 0; i < 3; i++) {
      const currRow = startX + i * deltaX;
      const currCol = startY + i * deltaY;
      if (!isValid(currRow, currCol, numRows, numCols)) {
        return false;
      }
      word += grid[currRow][currCol];
    }
    return word === 'MAS' || word === 'SAM';
  }

  function checkXPattern(row, col) {
    if (grid[row][col] !== 'A') {
      return false;
    }
    return (
      checkDirection(row, col, -1, 1) && // Top-left to bottom-right
      checkDirection(row, col, 1, -1) && // Bottom-right to top-left
      checkDirection(row, col, -1, -1) && // Top-right to bottom-left
      checkDirection(row, col, 1, 1)
    ); // Bottom-left to top-right
  }

  for (let row = 0; row < numRows; row++) {
    for (let col = 0; col < numCols; col++) {
      if (checkXPattern(row, col)) {
        count++;
      }
    }
  }
  return count;
}

function main() {
  const grid = loadFile();
  const numRows = grid.length;
  const numCols = grid[0].length;
  console.log(part1(grid, numRows, numCols));
  console.log(part2(grid, numRows, numCols));
}

main();
