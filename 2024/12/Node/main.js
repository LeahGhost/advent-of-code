const fs = require('fs');

const parseInput = (input) => {
  return input.split('\n').map(line => [...line]);
};

const isValid = (grid, p) => p.y >= 0 && p.y < grid.length && p.x >= 0 && p.x < grid[0].length;

const explore = (grid, p, visited) => {
  let area = 0, perimeter = 0;
  const directions = [{x: 0, y: 1}, {x: 1, y: 0}, {x: 0, y: -1}, {x: -1, y: 0}];

  for (let d of directions) {
    const np = {x: p.x + d.x, y: p.y + d.y}; 
    if (!isValid(grid, np) || grid[np.y][np.x] !== grid[p.y][p.x]) {
      perimeter++;
    } else if (!visited.has(`${np.x},${np.y}`)) {
      visited.add(`${np.x},${np.y}`);
      const [na, npPerimeter] = explore(grid, np, visited); 
      area += na;
      perimeter += npPerimeter;
    }
  }

  return [area + 1, perimeter];
};

const calculateResult = (input) => {
  const grid = parseInput(input);
  const visited = new Set();
  let result = 0;

  for (let y = 0; y < grid.length; y++) {
    for (let x = 0; x < grid[y].length; x++) {
      const p = {x, y};
      if (!visited.has(`${x},${y}`)) {
        visited.add(`${x},${y}`);
        const [area, perimeter] = explore(grid, p, visited);
        result += area * perimeter;
      }
    }
  }
  return result;
};

fs.readFile('../input.txt', 'utf8', (err, data) => {
  if (err) throw err;
  const result = calculateResult(data);
  console.log(`Result: ${result}`);
});
