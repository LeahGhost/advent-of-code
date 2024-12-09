'use strict';

const fs = require('fs');

function computeDiskChecksum() {
  const input = fs.readFileSync('../input.txt', 'utf-8').trim();
  let processingFile = true;
  let currentFileIndex = 0;
  const diskLayout = [];

  for (const size of input) {
    const blockSize = Number(size);
    if (processingFile) {
      diskLayout.push(...Array(blockSize).fill(currentFileIndex));
      currentFileIndex++;
    } else {
      diskLayout.push(...Array(blockSize).fill('.'));
    }
    processingFile = !processingFile;
  }

  for (let i = 0; i < diskLayout.length; i++) {
    if (diskLayout[i] === '.') {
      for (let j = diskLayout.length - 1; j > i; j--) {
        if (diskLayout[j] !== '.') {
          diskLayout[i] = diskLayout[j];
          diskLayout[j] = '.';
          break;
        }
      }
    }
  }

  const totalChecksum = diskLayout.reduce((sum, block, position) => {
    return block !== '.' ? sum + block * position : sum;
  }, 0);

  console.log('Checksum:', totalChecksum);
}

computeDiskChecksum();
