const { readFileSync } = require("fs");

const data = readFileSync("./input.txt", "utf-8").trim();

function processData() {
  const groups = data.split("\n\n").map(g => g.split("\n").map(r => r.split("")));
  const patterns = groups.map(g => {
    const counts = Array(g[0].length).fill(0);
    g.forEach(row => row.forEach((ch, i) => { if (ch === "#") counts[i]++; }));
    return { type: g[0][0] === "#" ? "lock" : "key", counts };
  });
  const locks = patterns.filter(p => p.type === "lock").map(p => p.counts);
  const keys = patterns.filter(p => p.type === "key").map(p => p.counts);
  let matchCount = 0;
  locks.forEach(lock => {
    keys.forEach(key => {
      if (lock.every((v, i) => v + key[i] <= 7)) matchCount++;
    });
  });
  return matchCount;
}

console.log(processData());
