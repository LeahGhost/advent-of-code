const fs = require('fs');
const path = require('path');
const PriorityQueue = require('js-priority-queue');

function readFileLines(filename) {
    const filepath = path.join(__dirname, filename);
    return fs.readFileSync(filepath, 'utf-8').trim().split('\n');
}

function buildGraph(filename) {
    const lines = readFileLines(filename);
    const rows = lines.length, cols = lines[0].length;
    const deltas = { up: [-1, 0], right: [0, 1], down: [1, 0], left: [0, -1] };
    let graph = {}, start = null, end = null;

    for (let i = 0; i < rows; i++) {
        for (let j = 0; j < cols; j++) {
            const char = lines[i][j];
            const node = `${i},${j}`;
            if (char === 'S') start = node;
            if (char === 'E') end = node;
            if (char === '#') continue;

            graph[node] = [];
            for (const [dy, dx] of Object.values(deltas)) {
                const ni = i + dy, nj = j + dx;
                if (ni >= 0 && ni < rows && nj >= 0 && nj < cols && lines[ni][nj] !== '#') {
                    graph[node].push(`${ni},${nj}`);
                }
            }
        }
    }
    return { graph, start, end };
}

function direction(from, to) {
    const [x1, y1] = from.split(',').map(Number);
    const [x2, y2] = to.split(',').map(Number);
    if (x2 > x1) return 'down';
    if (x1 > x2) return 'up';
    if (y2 > y1) return 'right';
    return 'left';
}

function calcWeight(currDir, nextDir) {
    const cost = {
        up: { left: 1001, right: 1001, down: 2001, up: 1 },
        down: { left: 1001, right: 1001, down: 1, up: 2001 },
        left: { left: 1, right: 2001, down: 1001, up: 1001 },
        right: { left: 2001, right: 1, down: 1001, up: 1001 }
    };
    return cost[currDir][nextDir];
}

function shortestPath(graph, start, end, initDir = 'right') {
    const pq = new PriorityQueue({ comparator: (a, b) => a[0] - b[0] });
    const dist = new Map(), prev = new Map(), visited = new Set();
    const directions = ['up', 'down', 'left', 'right'];

    directions.forEach(dir => dist.set(`${start},${dir}`, Infinity));
    dist.set(`${start},${initDir}`, 0);
    pq.queue([0, start, initDir]);

    while (pq.length) {
        const [cost, node, currDir] = pq.dequeue();
        if (visited.has(`${node},${currDir}`)) continue;
        visited.add(`${node},${currDir}`);

        for (const neighbor of graph[node]) {
            const nextDir = direction(node, neighbor);
            const weight = calcWeight(currDir, nextDir);
            const newCost = cost + weight;
            const key = `${neighbor},${nextDir}`;

            if (newCost < (dist.get(key) || Infinity)) {
                dist.set(key, newCost);
                prev.set(key, `${node},${currDir}`);
                pq.queue([newCost, neighbor, nextDir]);
            }
        }
    }

    let [minCost, bestDir] = directions.map(dir => [dist.get(`${end},${dir}`) || Infinity, dir]).reduce((a, b) => a[0] < b[0] ? a : b);
    if (minCost === Infinity) return { path: [], cost: Infinity };

    const path = [];
    let state = `${end},${bestDir}`;
    while (state) {
        const [node] = state.split(',');
        path.push(node);
        state = prev.get(state);
    }

    return { path: path.reverse(), cost: minCost };
}

function part1(file = 'input.txt') {
    const start = Date.now();
    const { graph, start: startNode, end } = buildGraph(file);
    const { path, cost } = shortestPath(graph, startNode, end);
    console.log(`Part 1: ${cost}`);
    return cost;
}

if (require.main === module) {
    part1();
}
