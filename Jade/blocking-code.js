var fs = require("fs");

var data = fs.readFileSync('input.txt'); // WAIT UNTIL READ FILE FINISH

console.log(data.toString());
console.log("----- Ended -----");