import java.io.*;
import java.util.*;

public class Solution {

  public static void main(String[] args) {
    /* Enter your code here. Read input from STDIN. Print output to STDOUT. Your class should be named Solution. */

    Scanner stringScanner = new Scanner(System.in);

    // hash table
    Hashtable<String, String> m = new Hashtable<String, String>();

    // get n from stdin
    Integer n = Integer.parseInt(stringScanner.nextLine());

    // read n line from stdin and update hash table
    for (int i = 0; i < n; i++) {
      String line = stringScanner.nextLine();

      String[] res = line.split(" ");

      m.put(res[0], res[1]);
    }

    // read the rest of lines from stdin
    while (stringScanner.hasNext()) {
      String name = stringScanner.nextLine();

      String number = m.get(name);
      if (number == null) {
        System.out.println("Not found");
      } else {
        System.out.println(name + "=" + number);
      }
    }
  }
}
