const fs = require('fs');

// Calculate the total distance between two lists of numbers
function calculateTotalDistance(leftList, rightList) {
  leftList.sort((a, b) => a - b);
  rightList.sort((a, b) => a - b);

  let totalDistance = 0;
  for (let i = 0; i < leftList.length; i++) {
    totalDistance += Math.abs(leftList[i] - rightList[i]);
  }
  return totalDistance;
}

// Read the file and parse it
function readFile(filePath) {
  const data = fs.readFileSync(filePath, 'utf-8');
  const lines = data.split('\n');

  const leftList = [];
  const rightList = [];

  lines.forEach(line => {
    if (line.trim()) {
      const [left, right] = line.split(' ').map(Number);
      leftList.push(left);
      rightList.push(right);
    }
  });

  return { leftList, rightList };
}

function main() {
  const inputFilePath = '../input.txt'

  // Read the input from the file
  const { leftList, rightList } = readFile(inputFilePath);

  // Calculate the total distance
  const result = calculateTotalDistance(leftList, rightList);
  console.log(`Total Distance: ${result}`);
}

main();
