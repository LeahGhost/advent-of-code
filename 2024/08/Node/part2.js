const fs = require('fs');

const data = fs.readFileSync('../input.txt', 'utf-8').split('\n').filter(line => line.trim() !== '');
const rows = data.length;
const cols = data[0].length;

class Antenna {
    constructor(row, col) {
        this.row = row;
        this.col = col;
    }
}

const frequencies = {};

for (let r = 0; r < rows; r++) {
    for (let c = 0; c < cols; c++) {
        const char = data[r][c];
        if (char !== '.') {
            if (!frequencies[char]) {
                frequencies[char] = [];
            }
            frequencies[char].push(new Antenna(r, c));
        }
    }
}

const antinodes = new Set();
function toString(row, col) {
    return `${row},${col}`;
}

for (const antennas of Object.values(frequencies)) {
    if (antennas.length < 2) {
        continue;
    }
    for (let i = 0; i < antennas.length; i++) {
        const a1 = antennas[i];
        antinodes.add(toString(a1.row, a1.col));
        for (let j = i + 1; j < antennas.length; j++) {
            const a2 = antennas[j];
            const dr = a2.row - a1.row;
            const dc = a2.col - a1.col;
            for (let k = 1; ; k++) {
                const r = a1.row + k * dr;
                const c = a1.col + k * dc;
                if (r < 0 || r >= rows || c < 0 || c >= cols) {
                    break;
                }
                antinodes.add(toString(r, c));
            }
            for (let k = 1; ; k++) {
                const r = a1.row - k * dr;
                const c = a1.col - k * dc;
                if (r < 0 || r >= rows || c < 0 || c >= cols) {
                    break;
                }
                antinodes.add(toString(r, c));
            }
        }
    }
}

console.log(antinodes.size);
