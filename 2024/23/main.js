import { readFileSync } from "fs";

const graph = new Map();
readFileSync("./input.txt", "utf-8").trim().split("\n").forEach(l => {
  const [a, b] = l.split("-");
  (graph.get(a) || graph.set(a, new Set()).get(a)).add(b);
  (graph.get(b) || graph.set(b, new Set()).get(b)).add(a);
});

const findConnections = () => [...graph.entries()].reduce((set, [node, neighbors]) => {
  if (node.startsWith("t")) neighbors.forEach(n1 => neighbors.forEach(n2 => {
    if (n1 !== n2 && graph.get(n1).has(n2)) set.add([node, n1, n2].sort().join("-"));
  }));
  return set;
}, new Set()).size;

const findMaxClique = () => {
  const cache = new Map();
  const maxClique = (clique, candidates) => cache.get(clique.join(",")) || 
    cache.set(clique.join(","), candidates.reduce((best, c) => {
      const newClique = maxClique([...clique, c].sort(), candidates.filter(n => graph.get(c).has(n)));
      return newClique.length > best.length ? newClique : best;
    }, clique)).get(clique.join(","));
  
  return [...graph.keys()].reduce((best, node) => {
    const clique = maxClique([node], [...graph.get(node)]);
    return clique.length > best.length ? clique : best;
  }, []).join(",");
};

console.log(findConnections());
console.log(findMaxClique());
