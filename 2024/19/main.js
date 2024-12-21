const fs = require('fs');
const input = fs.readFileSync('input.txt', 'utf-8').trim().split('\n');
const [patternsLine, , ...designs] = input;
const patterns = patternsLine.split(', ').map((p) => p.trim());

const countWays = (design, memo = {}) => {
    if (design in memo) return memo[design];
    if (design === '') return 1;
    memo[design] = patterns.reduce((sum, p) => sum + (design.startsWith(p) ? countWays(design.slice(p.length), memo) : 0), 0);
    return memo[design];
};

let possibleDesigns = 0, totalArrangements = 0;

designs.forEach((design) => {
    const ways = countWays(design);
    if (ways > 0) possibleDesigns++, totalArrangements += ways;
});

console.log(`Part 1: ${possibleDesigns}`);
console.log(`Part 2: ${totalArrangements}`);
