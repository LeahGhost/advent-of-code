const fs = require('fs');

const data = fs.readFileSync('../input.txt', 'utf8').trim().split('\n');
const grid = data.map(line => line.split('').map(Number));
const rows = grid.length, cols = grid[0].length;
const moves = [[-1, 0], [1, 0], [0, -1], [0, 1]];

const explore = (r, c) => {
  const q = [[r, c]], seen = Array.from({ length: rows }, () => Array(cols).fill(false));
  seen[r][c] = true;
  let count = 0;
  while (q.length) {
    const [x, y] = q.shift(), h = grid[x][y];
    for (const [dx, dy] of moves) {
      const nx = x + dx, ny = y + dy;
      if (nx >= 0 && nx < rows && ny >= 0 && ny < cols && !seen[nx][ny] && grid[nx][ny] === h + 1) {
        seen[nx][ny] = true;
        if (grid[nx][ny] === 9) count++;
        q.push([nx, ny]);
      }
    }
  }
  return count;
};

const total = grid.reduce((sum, row, r) => 
  sum + row.reduce((sub, val, c) => sub + (val === 0 ? explore(r, c) : 0), 0), 0);

console.log(total);
