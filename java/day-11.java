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

    List<List<Integer>> arr = new ArrayList<>();

    IntStream
      .range(0, 6)
      .forEach(i -> {
        try {
          arr.add(
            Stream
              .of(bufferedReader.readLine().replaceAll("\\s+$", "").split(" "))
              .map(Integer::parseInt)
              .collect(toList())
          );
        } catch (IOException ex) {
          throw new RuntimeException(ex);
        }
      });

    bufferedReader.close();

    Integer maxSum = Integer.MIN_VALUE;

    for (int x = 0; x < 4; x++) {
      for (int y = 0; y < 4; y++) {
        Integer currSum = 0;

        for (int dx = 0; dx < 3; dx++) {
          for (int dy = 0; dy < 3; dy++) {
            if (dx == 1 && (dy == 0 || dy == 2)) {
              continue;
            }
            currSum += arr.get(x + dx).get(y + dy);
          }
        }

        if (currSum > maxSum) {
          maxSum = currSum;
        }
      }
    }

    System.out.println(maxSum);
  }
}
