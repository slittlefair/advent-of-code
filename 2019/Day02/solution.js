"use strict";
exports.__esModule = true;
var helpers_1 = require("./../../helpers");
var initialInput = helpers_1.readFileAsCommaNumbers(__dirname)[0];
var input = initialInput.slice();
var runInstructions = function () {
    var _a;
    while (running) {
        _a = helpers_1.evaluateOpcode(input[i], input[i + 1], input[i + 2], input[i + 3], input, running), input = _a[0], running = _a[1];
        i = i + 4;
    }
};
var i = 0;
var running = true;
// Part 1
// Replacement instructions
input[1] = 12;
input[2] = 2;
runInstructions();
console.log('Part 1:', input[0]);
for (var noun = 0; noun < 100; noun++) {
    for (var verb = 0; verb < 100; verb++) {
        input = initialInput.slice();
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
