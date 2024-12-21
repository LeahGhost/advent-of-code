const { readFileSync } = require("fs");

const input = readFileSync("./input.txt", "utf-8").trim();
const keypadLayout = JSON.parse(readFileSync("./keypadLayout.json", "utf-8"));
const keyboardLayout = JSON.parse(readFileSync("./keyboardLayout.json", "utf-8"));

const getAllPaths = (grid) => {
  const findPaths = (startX, startY, endX, endY, path = "") => {
    if (startX === endX && startY === endY) return [path];
    if (startX < 0 || startX >= grid[0].length || startY < 0 || startY >= grid.length || grid[startY][startX] === null) return [];
    const paths = [];
    if (startX > endX) paths.push(...findPaths(startX - 1, startY, endX, endY, path + "<"));
    if (startX < endX) paths.push(...findPaths(startX + 1, startY, endX, endY, path + ">"));
    if (startY > endY) paths.push(...findPaths(startX, startY - 1, endX, endY, path + "^"));
    if (startY < endY) paths.push(...findPaths(startX, startY + 1, endX, endY, path + "v"));
    const minLength = Math.min(...paths.map(p => p.length));
    return paths.filter(p => p.length === minLength);
  };

  return grid.reduce((acc, row, startY) => ({
    ...acc,
    ...row.reduce((acc, start, startX) => {
      if (start === null) return acc;
      return {
        ...acc,
        [start]: grid.reduce((acc, row2, endY) => ({
          ...acc,
          ...row2.reduce((acc, end, endX) => {
            if (end === null) return acc;
            if (startX === endX && startY === endY) return { ...acc, [end]: [""] };
            return {
              ...acc,
              [end]: findPaths(startX, startY, endX, endY),
            };
          }, {}),
        }), {}),
      };
    }, {}),
  }), {});
};

const memoisation = new Map();
const computeShortestPath = (path, stepsLeft) => {
    if (stepsLeft === 0) return path.length;
    const subPaths = path.split(/A+/).map(s => s + "A").slice(0, -1);
    const Acount = (path.match(/A/g) || []).length - subPaths.length;
    
    return subPaths.reduce((totalLength, subPath) => {
      if (memoisation.has(stepsLeft + subPath)) return totalLength + memoisation.get(stepsLeft + subPath);
      
      const possiblePaths = subPath.split('').reduce((paths, char, i, arr) => {
        return paths.flatMap(p => keyboardPaths[arr[i - 1] || "A"][char].map(next => p + next + "A"));
      }, [""]);
      
      const minLength = possiblePaths.reduce((min, p) => {
        const pathLength = computeShortestPath(p, stepsLeft - 1);
        if (pathLength < min) {
          memoisation.set(stepsLeft + subPath, pathLength);
          return pathLength;
        }
        return min;
      }, Infinity);
      
      return totalLength + minLength;
    }, Acount);
  };
  

const solvePuzzle = (codes, steps) => codes.reduce((sum, code) => {
  let possiblePaths = [""];
  let position = "A";
  for (const char of code) {
    possiblePaths = possiblePaths.flatMap(p =>
      keypadPaths[position][char].map(next => p + next + "A")
    );
    position = char;
  }
  const minLength = possiblePaths.reduce((min, p) => {
    const length = computeShortestPath(p, steps);
    return length < min ? length : min;
  }, Infinity);
  return sum + minLength * Number.parseInt(code);
}, 0);

const keypadPaths = getAllPaths(keypadLayout);
const keyboardPaths = getAllPaths(keyboardLayout);

console.log(solvePuzzle(input.split("\n"), 2));
console.log(solvePuzzle(input.split("\n"), 25));
