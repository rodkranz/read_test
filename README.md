## Version

   * GO: 1.8.3
   * Node: 7.8.0
   * PHP: 5.6.30
   * Python: 2.7.13

## Generate File `data.json`
    [readTime] go run generate.go --fileSize 8192
     
## GO 
    [go] time ./goRead
    Started Reading json file.
    Bytes: 12558918857 Chunks: 191634
    ./goRead  0.06s user 5.53s system 80% cpu 6.904 total

## NODE 
    [node] time node nodeRead.js
    Started Reading json file.
    Bytes: 12558918857 Chunks: 191634
    node nodeRead.js  4.05s user 7.62s system 91% cpu 12.706 total

## PHP 
    [php] time php phpRead.php
    Started Reading json file.
    Bytes: 12558918857 Chunks: 191634
    php phpRead.php  0.68s user 5.97s system 84% cpu 7.914 total
    
## PYTHON
    [python] time python pythonRead.py
    Started Reading json file.
    Bytes: 12558918857 Chunks: 191634
    python pythonRead.py  0.37s user 6.41s system 64% cpu 10.471 total

## Results

| Language | Version |   User   | System  |  CPU  |   Time    |
|----------|---------|----------|---------|-------|-----------|
| GO       | 1.8.3   |**0.06s** |**5.53s**|  80%  | **6.904** | 
| Node     | 7.8.0   |  4.05s   |  7.62s  |  91%  |  12.706   |
| PHP      | 5.6.30  |  0.68s   |  5.97s  |  84%  |   7.914   | 
| Python   | 2.7.13  |  0.37s   |  6.41s  |**64%**|  10.471   | 
