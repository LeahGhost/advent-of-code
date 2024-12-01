const fs = require('fs');

// Calculate the similarity score
function calculateSimilarityScore(leftList, rightList) {
    let similarityScore = 0;

    // Iterate over each number in the left list
    leftList.forEach(left => {
        let count = 0;
        
        // Count how many times the number appears in the right list
        rightList.forEach(right => {
            if (left === right) {
                count++;
            }
        });

        // Add the product of the number and its count to the similarity score
        similarityScore += left * count;
    });

    return similarityScore;
}

// Function to read the input from the file
function readInputFile(filePath) {
    const data = fs.readFileSync(filePath, 'utf-8');
    const lines = data.split('\n').filter(line => line.trim() !== '');
    const leftList = [];
    const rightList = [];

    lines.forEach(line => {
        const [left, right] = line.split(' ').map(Number);
        leftList.push(left);
        rightList.push(right);
    });

    return { leftList, rightList };
}

function main() {
    const filePath = './input.txt';
    const { leftList, rightList } = readInputFile(filePath);
    const similarityScore = calculateSimilarityScore(leftList, rightList);
    console.log('Similarity Score:', similarityScore);
}

main();
