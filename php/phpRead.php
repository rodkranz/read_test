<?php
$fileName = 'data.json';
$handle = fopen($fileName, "rb");
if (FALSE === $handle) {
    exit("Error to read [file=${fileName}].");
}

echo "Started Reading json file. \n";

$nChunks = 0;
$nBytes = 0;
while (!feof($handle)) {
    $nBytes += strlen(fread($handle, 64*1024));
    $nChunks++;
}
fclose($handle);

echo "Bytes: " . $nBytes . " Chunks: " . $nChunks . "\n";
exit(0);