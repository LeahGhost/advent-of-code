const fs = require('fs');

const directions = [
  [-1, 0], 
  [0, 1], 
  [1, 0], 
  [0, -1]
];

function parseInput(filename) {
  const grid = fs.readFileSync(filename, 'utf8').split('\n');
  let guardPos = null;
  let guardDir = 0;

  for (let row = 0; row < grid.length; row++) {
    for (let col = 0; col < grid[row].length; col++) {
      const cell = grid[row][col];
      if (cell === '^') { guardPos = [row, col]; guardDir = 0; }
      if (cell === '>') { guardPos = [row, col]; guardDir = 1; }
      if (cell === 'v') { guardPos = [row, col]; guardDir = 2; }
      if (cell === '<') { guardPos = [row, col]; guardDir = 3; }
    }
  }

  return { grid, guardPos, guardDir };
}

function simulateGuardMovement(grid, guardPos, guardDir) {
  const visited = new Set();
  let visitedPositions = 0;

  const visitPosition = (pos) => {
    const posStr = `${pos[0]},${pos[1]}`;
    if (!visited.has(posStr)) {
      visited.add(posStr);
      visitedPositions++;
    }
  };

  visitPosition(guardPos);

  while (true) {
    const [x, y] = guardPos;
    const [dx, dy] = directions[guardDir];
    const newPos = [x + dx, y + dy];

    if (newPos[0] < 0 || newPos[0] >= grid.length || newPos[1] < 0 || newPos[1] >= grid[newPos[0]].length) break;

    if (grid[newPos[0]][newPos[1]] === '#') {
      guardDir = (guardDir + 1) % 4;
    } else {
      guardPos = newPos;
      visitPosition(guardPos);
    }
  }

  return visitedPositions;
}

function main() {
  const { grid, guardPos, guardDir } = parseInput('../input.txt');
  const distinctPositions = simulateGuardMovement(grid, guardPos, guardDir);
  console.log(`Distinct positions visited: ${distinctPositions}`);
}

main();
