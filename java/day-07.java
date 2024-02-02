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

    List<Integer> arr = Stream
      .of(bufferedReader.readLine().replaceAll("\\s+$", "").split(" "))
      .map(Integer::parseInt)
      .collect(toList());

    bufferedReader.close();

    // print array in reverse order
    for (int i = arr.size() - 1; i >= 0; i--) {
      System.out.printf("%d ", arr.get(i));
    }
  }
}
