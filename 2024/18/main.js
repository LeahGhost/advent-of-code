const fs = require('fs');

const input = fs.readFileSync(__dirname + '/input.txt', 'utf8').trim().split('\n');
const corruptedPositions = input.map(line => line.split(',').map(Number));
const GRID_SIZE = 71, MAX_CORRUPTED_POSITIONS = 1024;
const directions = [[0, 1], [1, 0], [0, -1], [-1, 0]];

const bfs = (grid, start, end) => {
    const queue = [start], visited = new Set([`${start[0]},${start[1]}`]);
    for (let steps = 0; queue.length; steps++) {
        const nextQueue = [];
        for (const [x, y] of queue) {
            if (x === end[0] && y === end[1]) return steps;
            for (const [dx, dy] of directions) {
                const [nextX, nextY] = [x + dx, y + dy];
                if (nextX >= 0 && nextX < GRID_SIZE && nextY >= 0 && nextY < GRID_SIZE && grid[nextX][nextY] === '.' && !visited.has(`${nextX},${nextY}`)) {
                    visited.add(`${nextX},${nextY}`);
                    nextQueue.push([nextX, nextY]);
                }
            }
        }
        queue.splice(0, queue.length, ...nextQueue);
    }
    return -1;
};

const main = () => {
    const grid = Array.from({ length: GRID_SIZE }, () => Array(GRID_SIZE).fill('.'));
    corruptedPositions.slice(0, MAX_CORRUPTED_POSITIONS).forEach(([x, y]) => grid[x][y] = '#');
    
    console.log(`Minimum steps to reach the exit: ${bfs(grid, [0, 0], [70, 70])}`);

    for (const [x, y] of corruptedPositions) {
        grid[x][y] = '#';
        if (bfs(grid, [0, 0], [70, 70]) === -1) {
            console.log(`First byte that blocks the path: (${x}, ${y})`);
            break;
        }
    }
};

main();
