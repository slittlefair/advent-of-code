const lower = 240298;
const upper = 784956;

let part1 = 0;
let part2 = 0;
for (let i = lower; i <= upper; i++) {
  let isValid = true;
  // Check increasing
  for (let j = 0; j < 5; j++) {
    if (('' + i)[j] > ('' + i)[j + 1]) {
      isValid = false;
      break;
    }
  }
  if (!isValid) {
    continue;
  }

  // Check double
  const doubles: Record<string, number> = {};
  for (let j = 0; j < 6; j++) {
    const num = ('' + i)[j];
    if (doubles[num] === undefined) {
      doubles[num] = 1;
    } else {
      doubles[num] = doubles[num] + 1;
    }
  }
  let part1Valid = false;
  let part2Valid = false;
  Object.values(doubles).forEach(v => {
    if (v > 1) {
      part1Valid = true;
    }
    if (v === 2) {
      part2Valid = true;
    }
  });
  if (part1Valid) {
    part1++;
  }
  if (part2Valid) {
    part2++;
  }
}

console.log('Part 1:', part1);
console.log('Part 2:', part2);
