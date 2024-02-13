import static java.util.stream.Collectors.joining;
import static java.util.stream.Collectors.toList;

import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.function.*;
import java.util.regex.*;
import java.util.stream.*;

public class Solution {

  public static void main(String[] args) throws IOException {
    BufferedReader bufferedReader = new BufferedReader(
      new InputStreamReader(System.in)
    );

    int n = Integer.parseInt(bufferedReader.readLine().trim());

    List<Integer> a = Stream
      .of(bufferedReader.readLine().replaceAll("\\s+$", "").split(" "))
      .map(Integer::parseInt)
      .collect(toList());

    // Write your code here
    int numSwaps = 0;

    // the maximum number will be put to the final index each loop
    for (int i = 0; i < n; i++) {
      // Track number of elements swapped during a single array traversal
      int numberOfSwaps = 0;

      for (int j = 0; j < n - 1 - i; j++) {
        // Swap adjacent elements if they are in decreasing order
        if (a.get(j) > a.get(j + 1)) {
          swap(a, j, j + 1);
          numberOfSwaps++;
        }
      }

      numSwaps += numberOfSwaps;

      // If no elements were swapped during a traversal, array is sorted
      if (numberOfSwaps == 0) {
        break;
      }
    }

    System.out.println("Array is sorted in " + numSwaps + " swaps.");
    System.out.println("First Element: " + a.get(0));
    System.out.println("Last Element: " + a.get(a.size() - 1));

    bufferedReader.close();
  }

  private static void swap(List<Integer> a, int i, int j) {
    int tmp = a.get(i);
    a.set(i, a.get(j));
    a.set(j, tmp);
  }
}
