import { readFileAsCommaNumbers, evaluateOpcode } from './../../helpers';

const initialInput = readFileAsCommaNumbers(__dirname)[0];
let input = [...initialInput];

const runInstructions = () => {
  while (running) {
    [input, running, i] = evaluateOpcode(
      input[i],
      input[i + 1],
      input[i + 2],
      input[i + 3],
      input,
      running,
      i,
    );
  }
};

let i = 0;
let running = true;

// Part 1

// Replacement instructions
input[1] = 12;
input[2] = 2;

runInstructions();
console.log('Part 1:', input[0]);

for (let noun = 0; noun < 100; noun++) {
  for (let verb = 0; verb < 100; verb++) {
    input = [...initialInput];
    i = 0;
    input[1] = noun;
    input[2] = verb;
    running = true;
    runInstructions();
    if (input[0] === 19690720) {
      console.log('Part 2:', 100 * noun + verb);
      noun = 100;
      verb = 100;
      break;
    }
  }
}
