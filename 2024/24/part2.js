const fs = require('fs');

function readFile(path) {
    return fs.readFileSync(path, 'utf-8').split('\n').map(line => line.trim());
}

function findMatch(a, b, gate, configs) {
    const pattern1 = `${a} ${gate} ${b} -> `;
    const pattern2 = `${b} ${gate} ${a} -> `;
    return configs.find(config => config.includes(pattern1) || config.includes(pattern2))?.split(' -> ')[1];
}

function swapWires(a, b, configs) {
    return configs.map(config => {
        const [inputs, output] = config.split(' -> ');
        if (output === a) return `${inputs} -> ${b}`;
        if (output === b) return `${inputs} -> ${a}`;
        return config;
    });
}

function processAdders(configs) {
    let carryWire = null;
    const swaps = [];
    let bit = 0;

    while (bit < 45) {
        const x = `x${bit.toString().padStart(2, '0')}`;
        const y = `y${bit.toString().padStart(2, '0')}`;
        const z = `z${bit.toString().padStart(2, '0')}`;

        if (bit === 0) {
            carryWire = findMatch(x, y, 'AND', configs);
        } else {
            const xorGate = findMatch(x, y, 'XOR', configs);
            const andGate = findMatch(x, y, 'AND', configs);
            const cinXorGate = findMatch(xorGate, carryWire, 'XOR', configs);

            if (!cinXorGate) {
                swaps.push(xorGate, andGate);
                configs = swapWires(xorGate, andGate, configs);
                bit = 0;
                continue;
            }

            if (cinXorGate !== z) {
                swaps.push(cinXorGate, z);
                configs = swapWires(cinXorGate, z, configs);
                bit = 0;
                continue;
            }

            const cinAndGate = findMatch(xorGate, carryWire, 'AND', configs);
            carryWire = findMatch(andGate, cinAndGate, 'OR', configs);
        }

        bit++;
    }

    return swaps;
}

function solve(input) {
    const divider = input.indexOf('');
    const configs = input.slice(divider + 1);
    const result = processAdders(configs);
    console.log(result.sort().join(','));
}

const input = readFile('input.txt');
solve(input);
