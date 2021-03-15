import { readFileAsCommaString, Coord } from './../../helpers';

const input = readFileAsCommaString(__dirname);

const wire1Coords: Coord[] = [];
let smallestDistance = -1;
let smallestIndex = -1;

const addCoord = (coord: Coord) => {
  const dupes = wire1Coords.filter(co => co.x === coord.x && co.y === coord.y);
  if (dupes.length === 0) {
    wire1Coords.push(coord);
  }
};

const checkDist = (coord: Coord) => {
  wire1Coords.forEach(co => {
    if (co.x === coord.x && co.y === coord.y) {
      const dist = Math.abs(coord.x - 0) + Math.abs(coord.y - 0);
      if (dist < smallestDistance || smallestDistance < 0) {
        smallestDistance = dist;
      }
      const index = coord.dist + co.dist;
      if (index < smallestIndex || smallestIndex < 0) {
        smallestIndex = index;
      }
    }
  });
};

const evaluatePath = (pathIndex: number, method: (co: Coord) => void) => {
  let lastCoord: Coord = { x: 0, y: 0, dist: 0 };
  let count = 1;
  input[pathIndex].forEach(inst => {
    const dir = inst[0];
    const num = parseInt(inst.substr(1), 10);
    switch (dir) {
      case 'U':
        for (let i = 1; i <= num; i++) {
          lastCoord = {
            x: lastCoord.x,
            y: lastCoord.y + 1,
            dist: count,
          };
          method(lastCoord);
          count++;
        }
        break;
      case 'R':
        for (let i = 1; i <= num; i++) {
          lastCoord = {
            x: lastCoord.x + 1,
            y: lastCoord.y,
            dist: count,
          };
          method(lastCoord);
          count++;
        }
        break;
      case 'D':
        for (let i = 1; i <= num; i++) {
          lastCoord = {
            x: lastCoord.x,
            y: lastCoord.y - 1,
            dist: count,
          };
          method(lastCoord);
          count++;
        }
        break;
      case 'L':
        for (let i = 1; i <= num; i++) {
          lastCoord = {
            x: lastCoord.x - 1,
            y: lastCoord.y,
            dist: count,
          };
          method(lastCoord);
          count++;
        }
        break;
      default:
        console.log('ERROR');
        break;
    }
  });
};

evaluatePath(0, addCoord);
evaluatePath(1, checkDist);

console.log('Part 1:', smallestDistance);
console.log('Part 2:', smallestIndex);
