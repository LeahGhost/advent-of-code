const fs = require('fs');

fs.readFile('../input.txt', 'utf8', (err, data) => {
  if (err) {
    console.error('Error opening file:', err);
    return;
  }

  const pattern = /mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)/g;
  const matches = [...data.matchAll(pattern)];

  let isEnabled = true;
  let total = 0;

  matches.forEach((match) => {
    if (match[0] === 'do()') {
      isEnabled = true;
    } else if (match[0] === "don't()") {
      isEnabled = false;
    } else if (match[1] && match[2]) {
      if (isEnabled) {
        const x = parseInt(match[1], 10);
        const y = parseInt(match[2], 10);
        total += x * y;
      }
    }
  });

  console.log('Sum of enabled multiplications:', total);
});
