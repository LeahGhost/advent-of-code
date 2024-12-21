const { readFileSync } = require('fs');

const mapData = readFileSync("./input.txt", "utf-8").trim().split("\n").map(line => line.split(""));

function discoverRoute(mapData) {
  const startPoint = { x: 0, y: 0 };
  const endPoint = { x: 0, y: 0 };
  let moveCount = 0;
  let positionsToExplore = [startPoint];
  const visitedLocations = new Map();

  mapData.forEach((line, rowIndex) => {
    line.forEach((cell, colIndex) => {
      if (cell === "S") Object.assign(startPoint, { x: colIndex, y: rowIndex });
      if (cell === "E") Object.assign(endPoint, { x: colIndex, y: rowIndex });
    });
  });

  while (positionsToExplore.length > 0) {
    const nextPositions = [];
    for (const currentPos of positionsToExplore) {
      const positionKey = `${currentPos.x},${currentPos.y}`;
      if (mapData[currentPos.y][currentPos.x] === "#" || visitedLocations.has(positionKey)) continue;
      visitedLocations.set(positionKey, moveCount);
      if (currentPos.x === endPoint.x && currentPos.y === endPoint.y) return computeTreasure(visitedLocations);
      nextPositions.push(
        { x: currentPos.x + 1, y: currentPos.y },
        { x: currentPos.x - 1, y: currentPos.y },
        { x: currentPos.x, y: currentPos.y + 1 },
        { x: currentPos.x, y: currentPos.y - 1 }
      );
    }
    positionsToExplore = nextPositions;
    moveCount++;
  }

  function computeTreasure(visitedLocations) {
    let silverTreasure = 0, goldTreasure = 0;
    visitedLocations.forEach((steps, locationKey) => {
      const [x, y] = locationKey.split(",").map(Number);
      for (let rowOffset = -20; rowOffset <= 20; rowOffset++) {
        for (let colOffset = -20; colOffset <= 20; colOffset++) {
          if (Math.abs(rowOffset) + Math.abs(colOffset) > 20 || mapData[y + rowOffset]?.[x + colOffset] === "#") continue;
          const distance = visitedLocations.get(`${x + colOffset},${y + rowOffset}`) - Math.abs(rowOffset) - Math.abs(colOffset) - steps;
          if (distance >= 100) {
            if (Math.abs(rowOffset) + Math.abs(colOffset) === 2) silverTreasure++;
            goldTreasure++;
          }
        }
      }
    });
    return `${silverTreasure}\n${goldTreasure}`;
  }
}

console.log(discoverRoute(mapData));
