import java.io.*;
import java.util.*;

public class Solution {

  public static void main(String[] args) {
    /* Enter your code here. Read input from STDIN. Print output to STDOUT. Your class should be named Solution. */

    Scanner scan = new Scanner(System.in);

    int numberOfStrings;
    numberOfStrings = Integer.parseInt(scan.nextLine());

    for (int i = 0; i < numberOfStrings; i++) {
      String currStr = scan.nextLine();

      String oddChars = "";
      String evenChars = "";

      for (int j = 0; j < currStr.length(); j++) {
        if (j % 2 == 1) {
          oddChars += currStr.charAt(j);
        } else {
          evenChars += currStr.charAt(j);
        }
      }

      System.out.printf("%s %s\n", evenChars, oddChars);
    }
  }
}
