
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

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
        BufferedReader reader = new BufferedReader(new FileReader("data.json"));
        char[] buf = new char[64*1024];
        long sz = 0;
        long n = 0;
        int chunks = 0;

        System.out.println("Started Reading json file.");
        while ((n = reader.read(buf)) >= 0) {
           sz += n;
           chunks++;
        }
        System.out.printf("Bytes: %d Chunks: %d\n", sz, chunks);

        reader.close();
    }

}
