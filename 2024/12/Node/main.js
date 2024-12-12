const fs = require('fs');

function readGrid(input) {
  const grid = input.split('\n').map((line) => line.split(''));
  return grid;
}

function isInside(grid, pt) {
  return (
    pt.y >= 0 && pt.y < grid.length && pt.x >= 0 && pt.x < grid[pt.y].length
  );
}

const dirs = [
  { x: 0, y: 1 },
  { x: 1, y: 0 },
  { x: 0, y: -1 },
  { x: -1, y: 0 },
];

function dfs(grid, pt, visited) {
  let area = 0;
  let perim = 0;

  dirs.forEach((d) => {
    const adj = { x: pt.x + d.x, y: pt.y + d.y };
    if (!isInside(grid, adj) || grid[adj.y][adj.x] !== grid[pt.y][pt.x]) {
      perim++;
    } else if (!visited.has(`${adj.x},${adj.y}`)) {
      visited.add(`${adj.x},${adj.y}`);
      const [a, p] = dfs(grid, adj, visited);
      area += a;
      perim += p;
    }
  });

  return [area + 1, perim];
}

function regionArea(grid, pt, visited, edges) {
  let area = 0;

  dirs.forEach((d) => {
    const adj = { x: pt.x + d.x, y: pt.y + d.y };
    if (!isInside(grid, adj) || grid[adj.y][adj.x] !== grid[pt.y][pt.x]) {
      const key = `${d.x},${d.y},${pt.x},${pt.y}`;
      edges.add(key);
    } else if (!visited.has(`${adj.x},${adj.y}`)) {
      visited.add(`${adj.x},${adj.y}`);
      area += regionArea(grid, adj, visited, edges);
    }
  });

  return area + 1;
}

function countEdges(edges) {
  let count = 0;
  dirs.forEach((d) => {
    const lines = {};

    edges.forEach((edge) => {
      const [dirX, dirY, ptX, ptY] = edge.split(',').map(Number);
      if (dirX === d.x && dirY === d.y) {
        const axis = d.x === 0 ? ptY : ptX;
        if (!lines[axis]) lines[axis] = [];
        lines[axis].push(d.x === 0 ? ptX : ptY);
      }
    });

    Object.values(lines).forEach((positions) => {
      positions.sort((a, b) => a - b);
      count++;
      for (let i = 1; i < positions.length; i++) {
        if (positions[i] - positions[i - 1] > 1) {
          count++;
        }
      }
    });
  });
  return count;
}

function part1(grid) {
  const visited = new Set();
  let result = 0;

  for (let y = 0; y < grid.length; y++) {
    for (let x = 0; x < grid[y].length; x++) {
      const pt = { x, y };
      if (!visited.has(`${x},${y}`)) {
        visited.add(`${x},${y}`);
        const [area, perim] = dfs(grid, pt, visited);
        result += area * perim;
      }
    }
  }

  return result;
}

function part2(grid) {
  const visited = new Set();
  let result = 0;

  for (let y = 0; y < grid.length; y++) {
    for (let x = 0; x < grid[y].length; x++) {
      const pt = { x, y };
      if (!visited.has(`${x},${y}`)) {
        visited.add(`${x},${y}`);
        const edges = new Set();
        const area = regionArea(grid, pt, visited, edges);
        result += area * countEdges(edges);
      }
    }
  }

  return result;
}

function main() {
  fs.readFile('../input.txt', 'utf8', (err, data) => {
    if (err) {
      console.error(err);
      return;
    }

    const grid = readGrid(data);

    console.log(`Part 1: ${part1(grid)}`);
    console.log(`Part 2: ${part2(grid)}`);
  });
}

main();
