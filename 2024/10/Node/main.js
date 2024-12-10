const fs = require('fs');
const data = fs.readFileSync('../input.txt', 'utf8').trim().split('\n');
const grid = data.map(line => line.split('').map(Number));
const rows = grid.length, cols = grid[0].length, moves = [[-1, 0], [1, 0], [0, -1], [0, 1]];
const bfsScores = (r, c) => {
  const q = [[r, c]], visited = Array.from({ length: rows }, () => Array(cols).fill(false));
  visited[r][c] = true;
  let nines = 0;
  while (q.length) {
    const [x, y] = q.shift();
    const h = grid[x][y];
    for (const [dx, dy] of moves) {
      const nx = x + dx, ny = y + dy;
      if (nx >= 0 && nx < rows && ny >= 0 && ny < cols && !visited[nx][ny] && grid[nx][ny] === h + 1) {
        visited[nx][ny] = true;
        if (grid[nx][ny] === 9) nines++;
        q.push([nx, ny]);
      }
    }
  }
  return nines;
};
const dfsRatings = (r, c, path) => {
  const pathKey = path.map(([x, y]) => `${x},${y}`).join('|');
  if (visitedPaths.has(pathKey)) return 0;
  visitedPaths.add(pathKey);
  if (grid[r][c] === 9) return 1;
  let trails = 0;
  for (const [dx, dy] of moves) {
    const nx = r + dx, ny = c + dy;
    if (nx >= 0 && nx < rows && ny >= 0 && ny < cols && grid[nx][ny] === grid[r][c] + 1) {
      trails += dfsRatings(nx, ny, [...path, [nx, ny]]);
    }
  }
  return trails;
};
let totalScores = 0, totalRatings = 0;
visitedPaths = new Set();
for (let r = 0; r < rows; r++) {
  for (let c = 0; c < cols; c++) {
    if (grid[r][c] === 0) {
      totalScores += bfsScores(r, c);
      totalRatings += dfsRatings(r, c, [[r, c]]);
    }
  }
}
console.log("PART 1:",totalScores,"PART 2:", totalRatings);
