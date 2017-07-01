import java.io.File;
import java.io.RandomAccessFile;
import java.nio.MappedByteBuffer;
import java.nio.channels.FileChannel;
import java.nio.channels.FileChannel.MapMode;
import java.util.ArrayList;
import java.util.List;

public class ReadFile3 {

    public static void main( String[] args ) throws Exception {
        //12GB
        long len = 12L * 1024 * 1024 * 1024;
        
        File file = new File("../data.json");

        RandomAccessFile raf = new RandomAccessFile(file, "rw");
        raf.setLength(len);
        FileChannel chan = raf.getChannel();

        long t0 = System.currentTimeMillis();
        
        List<MappedByteBuffer> maps = new ArrayList<MappedByteBuffer>(); 
        long chunckCount = 0l;
        long off = 0;
        
        while (off < len)
        {
           long chunk = Math.min(len - off, Integer.MAX_VALUE);
           MappedByteBuffer map;
           map = chan.map(MapMode.READ_WRITE, off, chunk);
           off += map.capacity();
           maps.add(map);
           chunckCount++;
        }
        raf.close();

        long t1 = System.currentTimeMillis();

        System.out.println("took: " + (t1 - t0) + "ms" + " Chunks " + chunckCount);
    }
    
}