const fs = require('fs');

fs.readFile('../input.txt', 'utf8', (err, data) => {
  if (err) {
    console.error('Error opening file:', err);
    return;
  }

  const pattern = /mul\((\d{1,3}),(\d{1,3})\)/g;
  const matches = [...data.matchAll(pattern)];

  let total = 0;

  matches.forEach((match) => {
    const x = parseInt(match[1], 10);
    const y = parseInt(match[2], 10);
    total += x * y;
  });

  console.log('Sum of all multiplications:', total);
});
