const { readFileSync } = require("fs");
const inputData = readFileSync("./input.txt", "utf-8").trim().split("\n").map(BigInt);

const transformValue = (v) => ((v = (v ^ (v * 64n)) % 16777216n), (v = (v ^ (v / 32n)) % 16777216n), (v ^ (v * 2048n)) % 16777216n);

const calculatePart1 = (nums) => nums.reduce((sum, num) => {
  for (let i = 0; i < 2000; i++) num = transformValue(num);
  return sum + num;
}, 0n);

const calculatePart2 = (numbers) => {
  const frequencyMap = new Map();
  numbers.forEach((num) => {
    const sequence = [null];
    const seenKeys = new Set();
    let remainder = Number(num) % 10;
    for (let i = 0; i < 3; i++) {
      num = transformValue(num);
      const digit = Number(num) % 10;
      sequence.push(digit - remainder);
      remainder = digit;
    }
    for (let i = 3; i < 2000; i++) {
      num = transformValue(num);
      const digit = Number(num) % 10;
      sequence.shift();
      sequence.push(digit - remainder);
      const key = sequence.join();
      if (!seenKeys.has(key)) {
        seenKeys.add(key);
        frequencyMap.set(key, (frequencyMap.get(key) || 0) + digit);
      }
      remainder = digit;
    }
  });
  return Math.max(...frequencyMap.values());
};

console.log(calculatePart1(inputData));
console.log(calculatePart2(inputData));
