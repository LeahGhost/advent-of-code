const fs = require('fs');

const input = fs.readFileSync('../input.txt', 'utf-8').trim().split('\n\n');
const rules = input[0].split('\n').map((r) => r.split('|').map(Number));
const updates = input[1].split('\n').map((line) => line.split(',').map(Number));

const ruleMap = rules.reduce((map, [x, y]) => {
  (map[x] ??= new Set()).add(y);
  return map;
}, {});

let validMiddleSum = 0,
  fixedMiddleSum = 0;
const invalidUpdates = [];

for (const update of updates) {
  isValidUpdate(update, ruleMap)
    ? (validMiddleSum += findMiddle(update))
    : invalidUpdates.push(update);
}

for (const update of invalidUpdates) {
  fixedMiddleSum += findMiddle(fixUpdate(update, ruleMap));
}

console.log('Part 1:', validMiddleSum);
console.log('Part 2:', fixedMiddleSum);

function isValidUpdate(update, ruleMap) {
  const pos = Object.fromEntries(update.map((p, i) => [p, i]));
  for (const [x, targets] of Object.entries(ruleMap)) {
    for (const y of targets) {
      if (pos[x] !== undefined && pos[y] !== undefined && pos[x] >= pos[y]) {
        return false;
      }
    }
  }
  return true;
}

function fixUpdate(update, ruleMap) {
  const inDegree = {},
    graph = {},
    set = new Set(update);
  for (const p of update) (inDegree[p] = 0), (graph[p] = []);
  for (const [x, targets] of Object.entries(ruleMap)) {
    const xNum = parseInt(x, 10);
    if (set.has(xNum)) {
      for (const y of targets) {
        if (set.has(y)) {
          graph[xNum].push(y);
          inDegree[y] = (inDegree[y] || 0) + 1;
        }
      }
    }
  }
  const queue = update.filter((p) => inDegree[p] === 0),
    sorted = [];
  while (queue.length) {
    const curr = queue.shift();
    sorted.push(curr);
    for (const n of graph[curr]) if (--inDegree[n] === 0) queue.push(n);
  }
  return sorted;
}

function findMiddle(arr) {
  return arr[Math.floor(arr.length / 2)];
}
