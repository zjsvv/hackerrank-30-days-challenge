import java.io.*;
import java.util.*;

interface AdvancedArithmetic {
  int divisorSum(int n);
}

class Calculator implements AdvancedArithmetic {

  public int divisorSum(int n) {
    if (n == 1) {
      return 1;
    }

    int res = 0;

    int sqrtRoot = (int) Math.sqrt(Double.valueOf(n));

    for (int i = 1; i <= sqrtRoot; i++) {
      if (n % i == 0) {
        res += i + n / i;
      }
    }

    if (sqrtRoot * sqrtRoot == n) {
      res -= sqrtRoot;
    }

    return res;
  }
}

class Solution {

  public static void main(String[] args) {
    Scanner scan = new Scanner(System.in);
    int n = scan.nextInt();
    scan.close();

    AdvancedArithmetic myCalculator = new Calculator();
    int sum = myCalculator.divisorSum(n);
    System.out.println(
      "I implemented: " + myCalculator.getClass().getInterfaces()[0].getName()
    );
    System.out.println(sum);
  }
}
