const fs = require('fs');

fs.readFile('../input.txt', 'utf8', (err, data) => {
  if (err) {
    console.log('Error opening file:', err);
    return;
  }

  const lines = data.trim().split('\n');
  let safeCount = 0;

  for (let line of lines) {
    const levels = parseLine(line);

    if (isSafeWithDampener(levels)) {
      safeCount++;
    }
  }

  console.log('Number of safe reports:', safeCount);
});

function parseLine(line) {
  return line.split(' ').map(Number);
}

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

function isSafeWithDampener(levels) {
  if (isSafe(levels)) {
    return true;
  }

  for (let i = 0; i < levels.length; i++) {
    const modified = [...levels.slice(0, i), ...levels.slice(i + 1)];

    if (isSafe(modified)) {
      return true;
    }
  }

  return false;
}
