var fs = require("fs");

fs.readFile('input.txt', function (err, data) { // CALLBACK FUNCTION
   if (err) return console.error(err);
   console.log(data.toString());
});

console.log("----- Ended -----");