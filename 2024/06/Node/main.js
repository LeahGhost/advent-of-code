const fs = require('fs');

function main() {
  const input = fs.readFileSync('../input.txt', 'utf8');
  const { part1, part2 } = execute(input);

  console.log(`Part 1: ${part1}`);
  console.log(`Part 2: ${part2}`);
}

function execute(input) {
  const grid = parseGrid(input);
  const part1 = trackGuard(grid).visitedCount;

  let part2 = 0;
  for (let y = 0; y < grid.length; y++) {
    for (let x = 0; x < grid[y].length; x++) {
      if (grid[y][x] === '.') {
        const updatedGrid = parseGrid(input);
        updatedGrid[y][x] = '#';

        if (trackGuard(updatedGrid).loopDetected) {
          part2++;
        }
      }
    }
  }

  return { part1, part2 };
}

function parseGrid(input) {
  return input.split('\n').filter(line => line).map(line => [...line]);
}

function trackGuard(grid) {
  const visited = new Set();
  const visitedWithDirection = new Set();

  let [cell, x, y] = findGuard(grid);

  while (x >= 0 && x < grid[0].length && y >= 0 && y < grid.length) {
    const positionKey = `${x},${y}`;
    const directionKey = `${x},${y},${cell}`;

    if (visitedWithDirection.has(directionKey)) {
      return { visitedCount: visited.size, loopDetected: true };
    }

    visited.add(positionKey);
    visitedWithDirection.add(directionKey);

    let turn = true;
    while (turn) {
      turn = false;

      switch (cell) {
        case '^':
          if (grid[y - 1] && grid[y - 1][x] === '#') {
            cell = '>';
            turn = true;
          }
          break;
        case '>':
          if (grid[y][x + 1] === '#') {
            cell = 'v';
            turn = true;
          }
          break;
        case 'v':
          if (grid[y + 1] && grid[y + 1][x] === '#') {
            cell = '<';
            turn = true;
          }
          break;
        case '<':
          if (grid[y][x - 1] === '#') {
            cell = '^';
            turn = true;
          }
          break;
        default:
          throw new Error("Invalid cell");
      }
    }

    switch (cell) {
      case '^': y--; break;
      case '>': x++; break;
      case 'v': y++; break;
      case '<': x--; break;
      default:
        throw new Error("Invalid cell");
    }
  }

  return { visitedCount: visited.size, loopDetected: false };
}

function findGuard(grid) {
  for (let y = 0; y < grid.length; y++) {
    for (let x = 0; x < grid[y].length; x++) {
      if (['^', 'v', '<', '>'].includes(grid[y][x])) {
        return [grid[y][x], x, y];
      }
    }
  }

  throw new Error("Guard not found");
}

main();
