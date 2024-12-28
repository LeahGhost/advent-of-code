const fs = require('fs');

const input = fs.readFileSync('input.txt', 'utf-8').trim().split('\n');

function parseInput(input) {
  const wireValues = {};
  const gates = [];

  for (const line of input) {
    if (line.includes('->')) {
      const [expression, output] = line.split(' -> ');
      gates.push({ expression, output });
    } else {
      const [wire, value] = line.split(': ');
      wireValues[wire] = parseInt(value, 10);
    }
  }

  return { wireValues, gates };
}

function evaluateGate(expression, wireValues) {
  const [input1, operator, input2] = expression.split(' ');
  const value1 = isNaN(input1) ? wireValues[input1] : parseInt(input1, 10);
  const value2 = isNaN(input2) ? wireValues[input2] : parseInt(input2, 10);

  switch (operator) {
    case 'AND': return value1 & value2;
    case 'OR': return value1 | value2;
    case 'XOR': return value1 ^ value2;
    default: throw new Error(`Unknown operator: ${operator}`);
  }
}

function simulateSystem(wireValues, gates) {
  const pendingGates = [...gates];

  while (pendingGates.length > 0) {
    for (let i = 0; i < pendingGates.length; i++) {
      const { expression, output } = pendingGates[i];
      const [input1, , input2] = expression.split(' ');

      if (
        (!isNaN(input1) || wireValues[input1] !== undefined) &&
        (!isNaN(input2) || wireValues[input2] !== undefined)
      ) {
        wireValues[output] = evaluateGate(expression, wireValues);
        pendingGates.splice(i, 1); 
        i--;
      }
    }
  }

  return wireValues;
}

function part1(input) {
  const { wireValues, gates } = parseInput(input);
  const finalValues = simulateSystem(wireValues, gates);
  const zWires = Object.keys(finalValues)
    .filter(wire => wire.startsWith('z'))
    .sort();

  const binaryValue = zWires.map(wire => finalValues[wire]).reverse().join('');
  return parseInt(binaryValue, 2); 
}

console.log('Part 1:', part1(input));
