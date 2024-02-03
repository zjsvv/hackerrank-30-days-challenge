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

    bufferedReader.close();

    int currLen = 0;
    int maxLen = 0;

    while (n > 0) {
      int remainder = n % 2;
      if (remainder == 1) {
        currLen++;
        if (currLen > maxLen) {
          maxLen = currLen;
        }
      } else {
        currLen = 0;
      }

      n /= 2;
    }

    System.out.println(maxLen);
  }
}
