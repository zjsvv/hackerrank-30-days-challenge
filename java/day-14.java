import java.io.*;
import java.math.*;
import java.text.*;
import java.util.*;
import java.util.regex.*;

class Difference {

  private int[] elements;
  public int maximumDifference;

  // Add your code here
  Difference(int[] elements) {
    this.elements = elements;
  }

  public void computeDifference() {
    for (int i = 0; i < elements.length - 1; i++) {
      for (int j = i + 1; j < elements.length; j++) {
        int currDiff = Math.abs(elements[i] - elements[j]);
        if (currDiff > this.maximumDifference) {
          this.maximumDifference = currDiff;
        }
      }
    }
  }
} // End of Difference class

public class Solution {

  public static void main(String[] args) {
    Scanner sc = new Scanner(System.in);
    int n = sc.nextInt();
    int[] a = new int[n];
    for (int i = 0; i < n; i++) {
      a[i] = sc.nextInt();
    }
    sc.close();

    Difference difference = new Difference(a);

    difference.computeDifference();

    System.out.print(difference.maximumDifference);
  }
}
