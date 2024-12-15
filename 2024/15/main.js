const fs = require('fs');

class GameData {
  constructor() {
    this.grid = [];
    this.moves = [];
    this.botX = 0;
    this.botY = 0;
  }
}

function loadGame(file) {
  try {
    const [gridData, movesData] = fs.readFileSync(file, 'utf-8').split('\n\n');
    const game = new GameData();

    game.grid = gridData.split('\n').map((line, y) =>
      line.split('').map((cell, x) => {
        if (cell === '@') [game.botX, game.botY] = [x, y];
        return cell;
      })
    );
    game.moves = movesData.trim().split('');

    return game;
  } catch (err) {
    console.error('Error reading file:', err);
    return null;
  }
}

function isMovable(obj) {
  return ['O', '[', ']'].includes(obj);
}

function shiftBox(game, x, y, dir) {
  const offsets = { left: [-1, 0], right: [1, 0], up: [0, -1], down: [0, 1] };
  const [dx, dy] = offsets[dir];
  const [newX, newY] = [x + dx, y + dy];
  const obj = game.grid[y][x];

  if (!isMovable(obj) || (isMovable(game.grid[newY][newX]) && !shiftBox(game, newX, newY, dir))) return false;

  if (obj === 'O' && game.grid[newY][newX] === '.') {
    [game.grid[y][x], game.grid[newY][newX]] = ['.', 'O'];
    return true;
  }

  if (obj === ']' && dir === 'left' && game.grid[newY][newX - 1] === '.') {
    [game.grid[y][x], game.grid[newY][newX - 1]] = ['.', '['];
    game.grid[newY][newX] = ']';
    return true;
  }

  if (obj === '[' && dir === 'right' && game.grid[newY][newX + 1] === '.') {
    [game.grid[y][x], game.grid[newY][newX + 1]] = ['.', ']'];
    game.grid[newY][newX] = '[';
    return true;
  }

  return false;
}

function calculateScore(game) {
  return game.grid.flat().reduce((score, cell, idx) => {
    if (cell === 'O') {
      const x = Math.floor(idx / game.grid[0].length);
      const y = idx % game.grid[0].length;
      score += 100 * x + y;
    }
    return score;
  }, 0);
}

function moveBot(game) {
  const directions = { '<': 'left', '>': 'right', '^': 'up', 'v': 'down' };
  const offsets = { left: [-1, 0], right: [1, 0], up: [0, -1], down: [0, 1] };

  for (const move of game.moves) {
    const dir = directions[move];
    const [dx, dy] = offsets[dir];
    const [newX, newY] = [game.botX + dx, game.botY + dy];

    shiftBox(game, newX, newY, dir);

    if (game.grid[newY]?.[newX] === '.') {
      [game.grid[game.botY][game.botX], game.grid[newY][newX]] = ['.', '@'];
      [game.botX, game.botY] = [newX, newY];
    }
  }

  return calculateScore(game);
}

function main() {
  const game = loadGame('input.txt');
  if (game) console.log(`Part One: ${moveBot(game)}`);
}

main();
