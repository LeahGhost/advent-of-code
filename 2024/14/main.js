const fs = require('fs');

const [w, h, s] = [101, 103, 100];
const input = fs.readFileSync('input.txt', 'utf8').trim().split('\n');

const robots = input.map((line) => {
  const match = line.match(/p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)/);
  const [, px, py, vx, vy] = match.map(Number);
  return { px, py, vx, vy };
});

function getPositionAfterTime({ px, py, vx, vy }, t) {
  const wrappedX = (((px + vx * t) % w) + w) % w;
  const wrappedY = (((py + vy * t) % h) + h) % h;
  return { x: wrappedX, y: wrappedY };
}

const finalPositions = robots.map((robot) => getPositionAfterTime(robot, s));

const [midX, midY] = [Math.floor(w / 2), Math.floor(h / 2)];

const quadrantCounts = [0, 0, 0, 0];

finalPositions.forEach(({ x, y }) => {
  if (x === midX || y === midY) {
    return;
  }
  quadrantCounts[(x >= midX) + 2 * (y >= midY)]++;
});

console.log(
  'Part 1:',
  quadrantCounts.reduce((p, c) => p * c, 1),
);
