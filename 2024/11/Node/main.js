const fs = require('fs');

const stones = fs.readFileSync('../input.txt', 'utf-8').trim().split(' ').map(Number);

const processStonesOptimized = (stones, blinks) => {
    let counts = new Map();

    for (const stone of stones) {
        counts.set(stone, (counts.get(stone) || 0) + 1);
    }

    for (let i = 0; i < blinks; i++) {
        let newCounts = new Map();

        for (const [stone, count] of counts.entries()) {
            if (stone === 0) {
                newCounts.set(1, (newCounts.get(1) || 0) + count);
            } else if (stone.toString().length % 2 === 0) {
                const mid = stone.toString().length / 2;
                const left = parseInt(stone.toString().slice(0, mid), 10);
                const right = parseInt(stone.toString().slice(mid), 10);
                newCounts.set(left, (newCounts.get(left) || 0) + count);
                newCounts.set(right, (newCounts.get(right) || 0) + count);
            } else {
                const newStone = stone * 2024;
                newCounts.set(newStone, (newCounts.get(newStone) || 0) + count);
            }
        }
        counts = newCounts; 
    }
    return Array.from(counts.values()).reduce((sum, count) => sum + count, 0);
};

const part1 = processStonesOptimized([...stones], 25); 
const part2 = processStonesOptimized([...stones], 75); 

console.log(`Part 1: Number of stones after 25 blinks: ${part1}`);
console.log(`Part 2: Number of stones after 75 blinks: ${part2}`);
