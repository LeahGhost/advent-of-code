const fs = require('fs');

const data = fs.readFileSync('../input.txt', 'utf8').trim().split('\n');

const rows = data.length;
const cols = data[0].length;

const signals = [];
for (let r = 0; r < rows; r++) {
  for (let c = 0; c < cols; c++) {
    const char = data[r][c];
    if (char !== '.') {
      signals.push({ r, c, type: char });
    }
  }
}

const recordAntinode = (r, c, locations) => {
  if (r >= 0 && r < rows && c >= 0 && c < cols) {
    locations.add(`${r},${c}`);
  }
};

const locations = new Set();
for (let x = 0; x < signals.length; x++) {
  for (let y = x + 1; y < signals.length; y++) {
    const sig1 = signals[x];
    const sig2 = signals[y];
    if (sig1.type !== sig2.type) continue;
    const diffR = sig2.r - sig1.r;
    const diffC = sig2.c - sig1.c;
    if (diffR % 2 === 0 && diffC % 2 === 0) {
      const midR = sig1.r + diffR / 2;
      const midC = sig1.c + diffC / 2;
      recordAntinode(midR, midC, locations);
    }
    recordAntinode(sig2.r + diffR, sig2.c + diffC, locations);
    recordAntinode(sig1.r - diffR, sig1.c - diffC, locations);
  }
}
console.log(locations.size);
