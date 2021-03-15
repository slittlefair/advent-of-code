import { readFileAsNumber } from '../../helpers';

const input = readFileAsNumber(__dirname);
let sum = 0;
input.forEach(num => {
  sum += Math.floor(num / 3) - 2;
});
console.log('Part 1:', sum);

sum = 0;
input.forEach(num => {
  while (num > 0) {
    const newFuel = Math.floor(num / 3) - 2;
    if (newFuel <= 0) {
      return;
    }
    num = newFuel;
    sum += newFuel;
  }
});
console.log('Part 2:', sum);
