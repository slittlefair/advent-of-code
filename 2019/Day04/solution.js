var lower = 240298;
var upper = 784956;
var part1 = 0;
var part2 = 0;
var _loop_1 = function (i) {
    var isValid = true;
    // Check increasing
    for (var j = 0; j < 5; j++) {
        if (('' + i)[j] > ('' + i)[j + 1]) {
            isValid = false;
            break;
        }
    }
    if (!isValid) {
        return "continue";
    }
    // Check double
    var doubles = {};
    for (var j = 0; j < 6; j++) {
        var num = ('' + i)[j];
        if (doubles[num] === undefined) {
            doubles[num] = 1;
        }
        else {
            doubles[num] = doubles[num] + 1;
        }
    }
    var part1Valid = false;
    var part2Valid = false;
    Object.values(doubles).forEach(function (v) {
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
};
for (var i = lower; i <= upper; i++) {
    _loop_1(i);
}
console.log('Part 1:', part1);
console.log('Part 2:', part2);
