import * as fs from 'fs';
import * as path from 'path';

export const lines = (dir: string, filePath: string = './input.txt'): string[] => {
  const txt = fs.readFileSync(path.join(dir, filePath));
  return txt.toString().split('\n');
};

export const readFileAsString = (dir: string, filePath: string = './input.txt'): string[] => {
  const txt = lines(dir, filePath);
  return txt.toString().split('\n');
};

export const readFileAsNumber = (dir: string, filePath: string = './input.txt'): number[] => {
  const txt = readFileAsString(dir, filePath);
  return txt.map(t => parseInt(t, 10));
};

export const readFileAsCommaString = (
  dir: string,
  filePath: string = './input.txt',
): string[][] => {
  const txt = lines(dir, filePath);
  return txt.map(ln => ln.toString().split(','));
};

export const readFileAsCommaNumbers = (
  dir: string,
  filePath: string = './input.txt',
): number[][] => {
  const txt = readFileAsCommaString(dir, filePath);
  return txt.map(ln => ln.map(t => parseInt(t, 10)));
};

export type Coord = {
  x: number;
  y: number;
  dist?: number;
};

export enum ParameterMode {
  POSITION,
  IMMEDIATE,
}

const breakdownOpcode = (opcode: number) => {
  const asString = String(opcode);
  const digits = asString.split('');
  const len = digits.length;
  const inst = digits[len - 2] === '0' ? parseInt(digits[len - 1], 10) : 99;
  const mode1 = digits[len - 3] !== undefined ? parseInt(digits[len - 3], 10) : 0;
  const mode2 = digits[len - 4] !== undefined ? parseInt(digits[len - 4], 10) : 0;
  const mode3 = digits[len - 5] !== undefined ? parseInt(digits[len - 5], 10) : 0;
  return [inst, mode1, mode2, mode3];
};

export const evaluateOpcode = (
  opcode: number,
  a: number,
  b: number,
  c: number,
  input: number[],
  running: boolean,
  i: number,
): [number[], boolean, number] => {
  const [inst, mode1, mode2, mode3] = breakdownOpcode(opcode);
  switch (inst) {
    case 1:
      input[c] = input[a] + input[b];
      i = i + 4;
      break;
    case 2:
      input[c] = input[a] * input[b];
      i = i + 4;
      break;
    case 3:
      i = i + 2;
      break;
    case 99:
      running = false;
      break;
    default:
      console.log('ERROR OPCODE', opcode);
      break;
  }
  return [input, running, i];
};
