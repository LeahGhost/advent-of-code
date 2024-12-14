const fs = require('fs');
const [w, h] = [101, 103];
const input = fs.readFileSync('input.txt', 'utf8').trim().split('\n');
const robots = input.map(line => {
  const [, px, py, vx, vy] = line.match(/p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)/).map(Number);
  return { px, py, vx, vy };
});

const getPosition = ({ px, py, vx, vy }, t) => ({
  x: (((px + vx * t) % w) + w) % w,
  y: (((py + vy * t) % h) + h) % h
});

const isChristmasTree = (positions) => {
  const grid = Array.from({ length: h }, () => Array(w).fill(false));
  positions.forEach(({ x, y }) => grid[y][x] = true);
  const tree = [
    [0, 0], [-1, 1], [0, 1], [1, 1], [-2, 2], [-1, 2], [0, 2], [1, 2], [2, 2]
  ];
  return grid.some((_, y) => grid[y].some((_, x) => 
    tree.every(([dx, dy]) => grid[(y + dy + h) % h][(x + dx + w) % w])
  ));
};

const findFewestSecondsForEasterEgg = () => {
  let t = 0;
  while (!isChristmasTree(robots.map(r => getPosition(r, t)))) t++;
  return t;
};

const calculatePart1 = (t) => {
  const positions = robots.map(r => getPosition(r, t));
  const [midX, midY] = [Math.floor(w / 2), Math.floor(h / 2)];
  const counts = positions.reduce((counts, { x, y }) => {
    if (x === midX || y === midY) return counts;
    counts[(x >= midX) + 2 * (y >= midY)]++;
    return counts;
  }, [0, 0, 0, 0]);
  return counts.reduce((a, c) => a * c, 1);
};

console.log('Part 1:', calculatePart1(100));
console.log('Part 2:', findFewestSecondsForEasterEgg());
