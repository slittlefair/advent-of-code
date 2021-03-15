"use strict";
exports.__esModule = true;
var helpers_1 = require("../../helpers");
var input = helpers_1.readFileAsNumber(__dirname);
var sum = 0;
input.forEach(function (num) {
    sum += Math.floor(num / 3) - 2;
});
console.log('Part 1:', sum);
sum = 0;
input.forEach(function (num) {
    while (num > 0) {
        var newFuel = Math.floor(num / 3) - 2;
        if (newFuel <= 0) {
            return;
        }
        num = newFuel;
        sum += newFuel;
    }
});
console.log('Part 2:', sum);
