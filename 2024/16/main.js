const fs = require('fs');
const PriorityQueue = require('js-priority-queue');

function readFileLines() {
    return fs.readFileSync('input.txt', 'utf-8')
             .split('\n')
             .filter(str => str !== '')
             .map(line => line.split(''));
  }  

function getAlternateDirections(direction) {
  switch (direction) {
    case '^': return ['>', '<', 'v'];
    case '>': return ['^', '<', 'v'];
    case 'v': return ['^', '>', '<'];
    case '<': return ['^', '>', 'v'];
    default: return [];
  }
}

function calculateNextMove(grid, lab, x, y, direction) {
    const currentScore = lab[[x, y]]?.[direction] || 0;
    const possibleMoves = [
      { move: 'v', x: x + 1, y: y },
      { move: '^', x: x - 1, y: y },
      { move: '>', x: x, y: y + 1 },
      { move: '<', x: x, y: y - 1 }
    ].filter(({ x, y, move }) => grid[x]?.[y] !== '#' && grid[x][y] !== '#' && move);
  
    const pq = new PriorityQueue({ comparator: (a, b) => a.score - b.score });
  
    possibleMoves.forEach(({ move, x: newX, y: newY }) => {
      const newScore = move === direction ? currentScore + 1 : currentScore + 1001;
      pq.queue({ move, x: newX, y: newY, score: newScore });
    });
  
    while (pq.length) {
      const { move, x: newX, y: newY, score } = pq.dequeue();
      const existingScore = lab[[newX, newY]]?.[move];
      if (!existingScore || score < existingScore) {
        lab[[newX, newY]] = { ...lab[[newX, newY]], [move]: score };
        updateAlternateDirections(lab, newX, newY, score + 1000, move);
        calculateNextMove(grid, lab, newX, newY, move);
      }
    }
  }
  
function updateAlternateDirections(lab, x, y, score, direction) {
  getAlternateDirections(direction).forEach(otherMove => {
    lab[[x, y]] = lab[[x, y]] || {};
    const oldScore = lab[[x, y]][otherMove];
    if (oldScore !== undefined && score < oldScore) {
      lab[[x, y]][otherMove] = score;
    } else if (oldScore === undefined) {
      lab[[x, y]][otherMove] = score;
    }
  });
}

function initaliseLab(grid) {
  let startX, startY;
  grid.forEach((row, i) => {
    row.forEach((cell, j) => {
      if (cell === 'S') {
        startX = i;
        startY = j;
      }
    });
  });

  const lab = {};
  lab[[startX, startY]] = { '>': 0 };
  calculateNextMove(grid, lab, startX, startY, '>');
  return lab;
}

function getMinimalScoreToEnd(grid, lab) {
  let endX, endY;
  grid.forEach((row, i) => {
    row.forEach((cell, j) => {
      if (cell === 'E') {
        endX = i;
        endY = j;
      }
    });
  });

  const scores = lab[[endX, endY]];
  return scores ? Math.min(...Object.values(scores)) : -1;
}

function getPreviousCoordinates(x, y, direction) {
  switch (direction) {
    case '^': return [x + 1, y];
    case '>': return [x, y - 1];
    case '<': return [x, y + 1];
    case 'v': return [x - 1, y];
    default: return [-1, -1];
  }
}

function findDirectionToPreviousCell(x, y, nextX, nextY) {
  if (x === nextX + 1) return '^';
  if (x === nextX - 1) return 'v';
  if (y === nextY + 1) return '<';
  if (y === nextY - 1) return '>';
  return '';
}

function getInitialDirectionAndCoords(grid, lab) {
  let endX, endY;
  grid.forEach((row, i) => {
    row.forEach((cell, j) => {
      if (cell === 'E') {
        endX = i;
        endY = j;
      }
    });
  });

  const scores = lab[[endX, endY]];
  const minScore = Math.min(...Object.values(scores));
  const previousCoords = Object.entries(scores)
    .filter(([key, score]) => score === minScore)
    .map(([key]) => getPreviousCoordinates(endX, endY, key));

  return [endX, endY, previousCoords];
}

function tracePathToStart(grid, lab, x, y, prevX, prevY) {
    const direction = findDirectionToPreviousCell(x, y, prevX, prevY);
    const currentScore = lab[[x, y]]?.[direction] || 0;
    grid[x][y] = 'O';
  
    const directions = [
      { dx: 1, dy: 0, move: '^' },
      { dx: -1, dy: 0, move: 'v' },
      { dx: 0, dy: 1, move: '<' },
      { dx: 0, dy: -1, move: '>' },
    ];
  
    directions.forEach(({ dx, dy, move }) => {
      if (grid[x + dx]?.[y + dy] !== '#' && lab[[x + dx, y + dy]]?.[move] === currentScore - (move === direction ? 1 : 1001)) {
        grid[x + dx][y + dy] = 'O';
        tracePathToStart(grid, lab, x + dx, y + dy, x, y);
      }
    });
  }
  
function calculateChecksum(grid) {
  return grid.reduce((sum, row) => sum + row.filter(c => ['O', 'E', 'S'].includes(c)).length, 0);
}

function main() {
  const grid = readFileLines();
  const lab = initaliseLab(grid);
  console.log('Part 1:', getMinimalScoreToEnd(grid, lab));

  const [endX, endY, previousCoords] = getInitialDirectionAndCoords(grid, lab);
  previousCoords.forEach(([prevX, prevY]) => {
    tracePathToStart(grid, lab, prevX, prevY, endX, endY);
  });

  console.log('Part 2:', calculateChecksum(grid));
}

main();
