
import java.io.BufferedInputStream;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;

/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

/**
 *
 * @author massahud
 */
public class ReadFile {

    public static void main(String[] args) throws FileNotFoundException, IOException {
        BufferedInputStream is = new BufferedInputStream(new FileInputStream("../data.json"));
        byte[] buf = new byte[64*1024];
        long sz = 0;
        long n = 0;
        int chunks = 0;

        System.out.println("Started Reading json file.");
        long time = System.currentTimeMillis();
        while ((n = is.read(buf)) >= 0) {
            sz += n;
            chunks++;
        }
        System.out.println(System.currentTimeMillis()-time);

        System.out.printf("Bytes: %d Chunks: %d\n", sz, chunks);
        is.close();
    }

}
