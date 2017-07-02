import java.io.File;
import java.io.RandomAccessFile;
import java.nio.MappedByteBuffer;
import java.nio.channels.FileChannel;
import java.nio.channels.FileChannel.MapMode;
import java.util.ArrayList;
import java.util.List;

public class ReadFile3 {

    public static void main( String[] args ) throws Exception {
        
        long t0 = System.currentTimeMillis();

        File file                   = new File("../data.json");

        final int chunck            = 64*1024;        
        long len                    = file.length();
        long chunckCount            = 0l;
        long off                    = 0;

        RandomAccessFile raf        = new RandomAccessFile(file, "rw");
        raf.setLength(len);
        
        FileChannel chan            = raf.getChannel();
        List<MappedByteBuffer> maps = new ArrayList<MappedByteBuffer>(); 

        while (off < len)
        {
           MappedByteBuffer map = chan.map(MapMode.READ_ONLY, off, chunck);
           off += map.capacity();
           maps.add(map);
           chunckCount++;
        }

        raf.close();
        long t1 = System.currentTimeMillis();

        System.out.println("Took: " + (t1 - t0) + "ms" + " Chunks: " + chunckCount + " Bytes: " + len );
    }
    
}