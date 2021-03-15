"use strict";
exports.__esModule = true;
var helpers_1 = require("./../../helpers");
var input = helpers_1.readFileAsCommaString(__dirname);
var wire1Coords = [];
var smallestDistance = -1;
var smallestIndex = -1;
var addCoord = function (coord) {
    var dupes = wire1Coords.filter(function (co) { return co.x === coord.x && co.y === coord.y; });
    if (dupes.length === 0) {
        wire1Coords.push(coord);
    }
};
var checkDist = function (coord) {
    wire1Coords.forEach(function (co) {
        if (co.x === coord.x && co.y === coord.y) {
            var dist = Math.abs(coord.x - 0) + Math.abs(coord.y - 0);
            if (dist < smallestDistance || smallestDistance < 0) {
                smallestDistance = dist;
            }
            var index = coord.dist + co.dist;
            if (index < smallestIndex || smallestIndex < 0) {
                smallestIndex = index;
            }
        }
    });
};
var evaluatePath = function (pathIndex, method) {
    var lastCoord = { x: 0, y: 0, dist: 0 };
    var count = 1;
    input[pathIndex].forEach(function (inst) {
        var dir = inst[0];
        var num = parseInt(inst.substr(1), 10);
        switch (dir) {
            case 'U':
                for (var i = 1; i <= num; i++) {
                    lastCoord = {
                        x: lastCoord.x,
                        y: lastCoord.y + 1,
                        dist: count
                    };
                    method(lastCoord);
                    count++;
                }
                break;
            case 'R':
                for (var i = 1; i <= num; i++) {
                    lastCoord = {
                        x: lastCoord.x + 1,
                        y: lastCoord.y,
                        dist: count
                    };
                    method(lastCoord);
                    count++;
                }
                break;
            case 'D':
                for (var i = 1; i <= num; i++) {
                    lastCoord = {
                        x: lastCoord.x,
                        y: lastCoord.y - 1,
                        dist: count
                    };
                    method(lastCoord);
                    count++;
                }
                break;
            case 'L':
                for (var i = 1; i <= num; i++) {
                    lastCoord = {
                        x: lastCoord.x - 1,
                        y: lastCoord.y,
                        dist: count
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
