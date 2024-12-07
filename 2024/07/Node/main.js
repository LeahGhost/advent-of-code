const fs = require("fs");

function evaluateExpression(numbers, target, index = 1, currentResult = numbers[0]) {
  if (index === numbers.length) {
    return currentResult === target;
  }
  const addResult = evaluateExpression(numbers, target, index + 1, currentResult + numbers[index]);
  const multiplyResult = evaluateExpression(numbers, target, index + 1, currentResult * numbers[index]);
  return addResult || multiplyResult;
}

function solveCalibration(inputFile) {
  const input = fs.readFileSync(inputFile, "utf-8").trim().split("\n");
  let totalCalibrationResult = 0;

  for (const line of input) {
    const [testValue, numbersString] = line.split(": ");
    const target = parseInt(testValue, 10);
    const numbers = numbersString.split(" ").map(Number);

    if (evaluateExpression(numbers, target)) {
      totalCalibrationResult += target;
    }
  }

  return totalCalibrationResult;
}

const inputFile = "../input.txt";
const result = solveCalibration(inputFile);
console.log("Total Calibration Result:", result);
