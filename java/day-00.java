import java.io.*;
import java.util.*;

public class Solution {

  public static void main(String[] args) {
    /* Enter your code here. Read input from STDIN. Print output to STDOUT. Your class should be named Solution. */
    System.out.println("Hello, World.");

    Scanner stringScanner = new Scanner(System.in);
    
    // Process each Java Scanner String input
    while (stringScanner.hasNext()) {
      String line = stringScanner.nextLine();
      System.out.println(line);
    }
  }
}
