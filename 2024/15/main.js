const fs = require('fs');

function loadInputData(fileName) {
  return fs.readFileSync(fileName, 'utf8').split('\n');
}

function expandWarehouseLayout(warehouse) {
  const modified = [];
  for (const row of warehouse) {
    const newRow = [];
    for (const cell of row) {
      switch (cell) {
        case '.':
          newRow.push('.', '.');
          break;
        case '#':
          newRow.push('#', '#');
          break;
        case '@':
          newRow.push('@', '.');
          break;
        default:
          newRow.push('[', ']');
          break;
      }
    }
    modified.push(newRow);
  }
  return modified;
}

function locateCharacter(target, warehouse) {
  for (let i = 0; i < warehouse.length; i++) {
    for (let j = 0; j < warehouse[i].length; j++) {
      if (warehouse[i][j] === target) {
        return [i, j];
      }
    }
  }
  throw new Error(`Character '${target}' not found`);
}

function translateDirectionToSteps(direction) {
  switch (direction) {
    case '^':
      return [-1, 0];
    case 'v':
      return [1, 0];
    case '<':
      return [0, -1];
    case '>':
      return [0, 1];
    default:
      throw new Error(`Invalid direction '${direction}'`);
  }
}

function part1(data) {
  const [warehouse, instructions] = parseInput(data);
  let [row, col] = locateCharacter('@', warehouse);

  for (const instruction of instructions) {
    const [dRow, dCol] = translateDirectionToSteps(instruction);
    let nRow = row + dRow,
      nCol = col + dCol;
    let steps = 0;

    while (true) {
      const newSquare = warehouse[nRow + steps * dRow][nCol + steps * dCol];
      if (newSquare === 'O') {
        steps++;
      } else if (newSquare === '.') {
        warehouse[row][col] = '.';
        warehouse[nRow][nCol] = '@';
        for (let i = 1; i <= steps; i++) {
          warehouse[nRow + i * dRow][nCol + i * dCol] = 'O';
        }
        [row, col] = [nRow, nCol];
        break;
      } else if (newSquare === '#') {
        break;
      }
    }
  }

  let total = 0;
  for (let i = 0; i < warehouse.length; i++) {
    for (let j = 0; j < warehouse[i].length; j++) {
      if (warehouse[i][j] === 'O') {
        total += i * 100 + j;
      }
    }
  }
  return total;
}

function findMovableBlocks(warehouse, instruction, row, col) {
  const [dRow, dCol] = translateDirectionToSteps(instruction);
  const newRow = row + dRow,
    newCol = col + dCol;
  let blocksToCheck = [];

  if (warehouse[newRow][newCol] === '#') return null;
  if (warehouse[newRow][newCol] === ']') {
    blocksToCheck.push([newRow, newCol - 1, newCol]);
  } else if (warehouse[newRow][newCol] === '[') {
    blocksToCheck.push([newRow, newCol, newCol + 1]);
  } else if (warehouse[newRow][newCol] === '.') {
    return [[row, col]];
  }

  let blocks = [[row, col]];
  while (blocksToCheck.length > 0) {
    const block = blocksToCheck.shift();
    if (instruction === 'v' || instruction === '^') {
      if (
        warehouse[block[0] + dRow][block[1]] === '#' ||
        warehouse[block[0] + dRow][block[2]] === '#'
      ) {
        return null;
      }
      if (warehouse[block[0] + dRow][block[1]] === ']') {
        blocksToCheck.push([block[0] + dRow, block[1] - 1, block[1]]);
      } else if (warehouse[block[0] + dRow][block[1]] === '[') {
        blocksToCheck.push([block[0] + dRow, block[1], block[2]]);
      }
      if (warehouse[block[0] + dRow][block[2]] === '[') {
        blocksToCheck.push([block[0] + dRow, block[2], block[2] + 1]);
      }
    } else if (instruction === '<') {
      if (warehouse[block[0]][block[1] - 1] === '#') {
        return null;
      } else if (warehouse[block[0]][block[1] - 1] === ']') {
        blocksToCheck.push([block[0], block[1] - 2, block[1] - 1]);
      }
    } else if (instruction === '>') {
      if (warehouse[block[0]][block[2] + 1] === '#') {
        return null;
      } else if (warehouse[block[0]][block[2] + 1] === '[') {
        blocksToCheck.push([block[0], block[2] + 1, block[2] + 2]);
      }
    }
    blocks.push(block);
  }
  return blocks;
}

function calculateGPSScore(warehouse) {
  let total = 0;
  for (let i = 0; i < warehouse.length; i++) {
    for (let j = 0; j < warehouse[i].length; j++) {
      if (warehouse[i][j] === '[') {
        total += i * 100 + j;
      }
    }
  }
  return total;
}

function moveBlocks(warehouse, blocks, instruction) {
  const [dRow, dCol] = translateDirectionToSteps(instruction);
  for (let n = blocks.length - 1; n >= 0; n--) {
    const block = blocks[n];
    if (block.length === 2) {
      warehouse[block[0] + dRow][block[1] + dCol] = '@';
      warehouse[block[0]][block[1]] = '.';
    } else {
      warehouse[block[0]][block[1]] = '.';
      warehouse[block[0]][block[2]] = '.';
      warehouse[block[0] + dRow][block[1] + dCol] = '[';
      warehouse[block[0] + dRow][block[2] + dCol] = ']';
    }
  }
}

function part2(data) {
  let [warehouse, instructions] = parseInput(data);
  warehouse = expandWarehouseLayout(warehouse);
  let [rx, ry] = locateCharacter('@', warehouse);

  for (const instruction of instructions) {
    const blocks = findMovableBlocks(warehouse, instruction, rx, ry);

    if (blocks === null) continue;

    moveBlocks(warehouse, blocks, instruction);

    const [dx, dy] = translateDirectionToSteps(instruction);
    rx += dx;
    ry += dy;
  }

  return calculateGPSScore(warehouse);
}

function parseInput(data) {
  let warehouse = [];
  let j = 0;

  for (let i = 0; i < data.length; i++) {
    if (data[i] === '') {
      j = i;
      break;
    }
    warehouse.push(data[i].split(''));
  }

  const instructions = data.slice(j + 1).join('');
  return [warehouse, instructions];
}

const fileName = 'input.txt';
const data = loadInputData(fileName);
console.log('Part 1:', part1(data));
console.log('Part 2:', part2(data));
