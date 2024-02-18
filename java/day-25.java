import java.io.*;
import java.util.*;

public class Solution {

  public static void main(String[] args) {
    /* Enter your code here. Read input from STDIN. Print output to STDOUT. Your class should be named Solution. */

    Scanner sc = new Scanner(System.in);
    int n = sc.nextInt();

    while (n-- > 0) {
      int num = sc.nextInt();

      if (isPrime(num)) {
        System.out.println("Prime");
      } else {
        System.out.println("Not prime");
      }
    }
  }

  private static boolean isPrime(int num) {
    if (num == 1) {
      return false;
    }

    int sqrtRoot = (int) Math.sqrt(Double.valueOf(num));

    for (int i = 2; i < sqrtRoot + 1; i++) {
      if (num % i == 0) {
        return false;
      }
    }

    return true;
  }
}
