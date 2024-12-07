const fs = require('fs');

function evaluateExpression(numbers, target, index, currentResult, part) {
  if (index === numbers.length) {
    return currentResult === target;
  }

  const addResult = evaluateExpression(
    numbers,
    target,
    index + 1,
    currentResult + numbers[index],
    part,
  );
  const multiplyResult = evaluateExpression(
    numbers,
    target,
    index + 1,
    currentResult * numbers[index],
    part,
  );

  if (part === 2) {
    const concatResult = evaluateExpression(
      numbers,
      target,
      index + 1,
      parseInt(currentResult.toString() + numbers[index].toString()),
      part,
    );
    return addResult || multiplyResult || concatResult;
  }

  return addResult || multiplyResult;
}

function solveCalibration(inputFile, part) {
  const input = fs.readFileSync(inputFile, 'utf-8').trim().split('\n');
  let totalCalibrationResult = 0;

  for (const line of input) {
    const [testValue, numbersString] = line.split(': ');
    const target = parseInt(testValue, 10);
    const numbers = numbersString.split(' ').map(Number);

    if (evaluateExpression(numbers, target, 1, numbers[0], part)) {
      totalCalibrationResult += target;
    }
  }

  return totalCalibrationResult;
}

const inputFile = '../input.txt';

const part1Result = solveCalibration(inputFile, 1);
console.log('Part 1 - Total Calibration Result:', part1Result);

const part2Result = solveCalibration(inputFile, 2);
console.log('Part 2 - Total Calibration Result:', part2Result);
