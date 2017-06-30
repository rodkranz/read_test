var fs = require('fs');
var stream = fs.createReadStream('../data.json', 'UTF-8');

var nBytes = 0;
var nChunks = 0;

stream.once('data', function () {
  console.log('Started Reading json file.');
});

stream.on('data', function (chunk) {
  nChunks++;
  nBytes += chunk.length
});

stream.on('end', function () {
  console.log("Bytes:", nBytes, "Chunks:", nChunks);
});