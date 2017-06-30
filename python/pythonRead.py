file = '../data.json'

f = open(file, 'rb')

n_bytes = 0
n_chunks = 0

print "Started Reading json file."

while True:
    piece = f.read(1024*64)
    if not piece:
        break

    n_bytes += len(piece)
    n_chunks += 1

f.close()
print "Bytes: %d Chunks: %d" % (n_bytes, n_chunks)