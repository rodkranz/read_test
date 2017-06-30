## Version

   * GO: 1.8.3
   * Python: 2.7.13
   * Node: v7.8.0
   * PHP: 5.6.30

## Generate File `data.json`
    [readTime] go run generate.go --fileSize 8192
     
## GO 
    [go] time ./goRead
    Started Reading json file.
    Bytes: 12558918857 Chunks: 191634
    ./goRead  0.06s user 6.15s system 64% cpu 9.656 total

## NODE 
    [node] time node nodeRead.js
    Started Reading json file.
    Bytes: 12558918857 Chunks: 191634
    node nodeRead.js  3.67s user 5.45s system 97% cpu 9.343 total

## PHP 
    [php] time php phpRead.php
    Started Reading json file.
    Bytes: 12558918857 Chunks: 191634
    php phpRead.php  0.68s user 4.71s system 84% cpu 6.411 total

## PYTHON
    [python] time python pythonRead.py                                                                                                                                                                                                                                   
    Started Reading json file.
    Bytes: 12558918857 Chunks: 191634
    python pythonRead.py  0.30s user 4.71s system 80% cpu 6.197 total
