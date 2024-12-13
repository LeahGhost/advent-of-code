const fs = require('fs');
const path = require('path');

const input = fs
  .readFileSync(path.join(__dirname, '../input.txt'), 'utf-8')
  .split('\n\n')
  .map((line) => line.split('\n'));

function calc(A, b) {
  const det = A[0][0] * A[1][1] - A[0][1] * A[1][0];
  if (det === 0) return [0, 0];
  return [
    (b[0] * A[1][1] - b[1] * A[0][1]) / det,
    (A[0][0] * b[1] - A[1][0] * b[0]) / det,
  ];
}

class Machine {
  constructor(lines) {
    this.a = (lines[0].match(/\d+/g) || []).map(Number);
    this.b = (lines[1].match(/\d+/g) || []).map(Number);
    this.p = (lines[2].match(/\d+/g) || []).map(Number);
    this.p2 = this.p.map((x) => x + 1e13);
  }

  cost(part2 = false) {
    const [x, y] = calc(
      [
        [this.a[0], this.b[0]],
        [this.a[1], this.b[1]],
      ],
      part2 ? this.p2 : this.p,
    );
    return Number.isInteger(x) && Number.isInteger(y) ? x * 3 + y : 0;
  }
}

const machines = input.map((lines) => new Machine(lines));
console.log(machines.reduce((sum, m) => sum + m.cost(), 0));
console.log(machines.reduce((sum, m) => sum + m.cost(true), 0));
