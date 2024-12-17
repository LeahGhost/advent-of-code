const fs = require("fs");

fs.readFile( "./input.txt", "utf8", (err, content) => {
  if (err) {
    console.error('Error reading file:', err);
    return;
  }

  const config = parseData(content);
  console.log(runProgram(config));
});

const parseData = (text) => {
  const result = {};
  text.split("\n").forEach(line => {
    if (line.includes("Register")) line = line.replace("Register ", "");
    const [key, value] = line.split(": ");
    result[key] = key === "Program" ? value.split(",").map(Number) : parseInt(value);
  });
  return result;
};

const runProgram = (data) => {
  const output = [];
  let { A, B, C, Program } = data;
  let pointer = 0;

  const getValue = operand => {
    if (operand <= 3) return operand;
    if (operand === 4) return A;
    if (operand === 5) return B;
    return C;
  };

  const executeInstruction = (opcode, operand) => {
    switch (opcode) {
      case 0: A = Math.trunc(A / (2 ** getValue(operand))); break;
      case 1: B ^= operand; break;
      case 2: B = getValue(operand) % 8; break;
      case 3: if (A !== 0) pointer = operand - 2; break;
      case 4: B ^= C; break;
      case 5: output.push(getValue(operand) % 8); break;
      case 6: B = Math.trunc(A / (2 ** getValue(operand))); break;
      case 7: C = Math.trunc(A / (2 ** getValue(operand))); break;
    }
    pointer += 2;
  };

  while (pointer < Program.length) {
    const opcode = Program[pointer];
    const operand = Program[pointer + 1];
    executeInstruction(opcode, operand);
  }

  return output.join(",");
};
