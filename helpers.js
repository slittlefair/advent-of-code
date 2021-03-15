"use strict";
exports.__esModule = true;
var fs = require("fs");
var path = require("path");
exports.lines = function (dir, filePath) {
    if (filePath === void 0) { filePath = './input.txt'; }
    var txt = fs.readFileSync(path.join(dir, filePath));
    return txt.toString().split('\n');
};
exports.readFileAsString = function (dir, filePath) {
    if (filePath === void 0) { filePath = './input.txt'; }
    var txt = exports.lines(dir, filePath);
    return txt.toString().split('\n');
};
exports.readFileAsNumber = function (dir, filePath) {
    if (filePath === void 0) { filePath = './input.txt'; }
    var txt = exports.readFileAsString(dir, filePath);
    return txt.map(function (t) { return parseInt(t, 10); });
};
exports.readFileAsCommaString = function (dir, filePath) {
    if (filePath === void 0) { filePath = './input.txt'; }
    var txt = exports.lines(dir, filePath);
    return txt.map(function (ln) { return ln.toString().split(','); });
};
exports.readFileAsCommaNumbers = function (dir, filePath) {
    if (filePath === void 0) { filePath = './input.txt'; }
    var txt = exports.readFileAsCommaString(dir, filePath);
    return txt.map(function (ln) { return ln.map(function (t) { return parseInt(t, 10); }); });
};
exports.evaluateOpcode = function (opcode, a, b, c, input, running) {
    switch (opcode) {
        case 1:
            input[c] = input[a] + input[b];
            break;
        case 2:
            input[c] = input[a] * input[b];
            break;
        case 99:
            running = false;
            break;
        default:
            console.log('ERROR OPCODE', opcode);
            break;
    }
    return [input, running];
};
