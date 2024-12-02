const fs = require('fs');

function main() {
  fs.readFile('../input.txt', 'utf8', (err, data) => {
    if (err) {
      console.error('Error opening file:', err);
      return;
    }

    // Trim whitespace and filter out empty lines
    const lines = data
      .trim()
      .split('\n')
      .filter((line) => line.trim() !== '');

    let safeCount = 0;

    lines.forEach((line) => {
      const levels = parseLine(line);
      if (isSafe(levels)) {
        safeCount++;
      }
    });

    console.log('Number of safe reports:', safeCount);
  });
}

// Converts a space-separated line of numbers into an array of integers
function parseLine(line) {
  return line.split(' ').map(Number);
}

// Checks if a report satisfies the safety rules
function isSafe(levels) {
  if (levels.length < 2) {
    return true;
  }

  let increasing = true;
  let decreasing = true;

  for (let i = 1; i < levels.length; i++) {
    const diff = levels[i] - levels[i - 1];
    if (diff < -3 || diff > 3) {
      return false;
    }
    if (diff > 0) {
      decreasing = false;
    } else if (diff < 0) {
      increasing = false;
    } else {
      return false;
    }
  }
  return increasing || decreasing;
}

main();
