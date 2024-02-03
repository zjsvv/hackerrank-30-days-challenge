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

class Result {

  /*
   * Complete the 'factorial' function below.
   *
   * The function is expected to return an INTEGER.
   * The function accepts INTEGER n as parameter.
   */

  public static int factorial(int n) {
    // f(1)=1
    // f(2)=2
    // f(3)=3*f(2)
    // f(4)=4*f(3)

    int res = 1;

    for (int i = 1; i <= n; i++) {
      res *= i;
    }

    return res;
  }
}

public class Solution {

  public static void main(String[] args) throws IOException {
    BufferedReader bufferedReader = new BufferedReader(
      new InputStreamReader(System.in)
    );
    BufferedWriter bufferedWriter = new BufferedWriter(
      new FileWriter(System.getenv("OUTPUT_PATH"))
    );

    int n = Integer.parseInt(bufferedReader.readLine().trim());

    int result = Result.factorial(n);

    bufferedWriter.write(String.valueOf(result));
    bufferedWriter.newLine();

    bufferedReader.close();
    bufferedWriter.close();
  }
}
