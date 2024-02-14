import java.io.*;
import java.util.*;

class Printer<T> {

  // Generic function to print elements of an array
  public void printArray(T[] array) {
    for (T element : array) {
      System.out.println(element);
    }
  }
}

public class Solution {

  public static void main(String[] args) {
    /* Enter your code here. Read input from STDIN. Print output to STDOUT. Your class should be named Solution. */
    Scanner scanner = new Scanner(System.in);
    int n = scanner.nextInt();
    Integer[] intArray = new Integer[n];
    for (int i = 0; i < n; i++) {
      intArray[i] = scanner.nextInt();
    }

    n = scanner.nextInt();
    String[] stringArray = new String[n];
    for (int i = 0; i < n; i++) {
      stringArray[i] = scanner.next();
    }

    Printer<Integer> intPrinter = new Printer<>();
    Printer<String> stringPrinter = new Printer<>();
    intPrinter.printArray(intArray);
    stringPrinter.printArray(stringArray);
    if (Printer.class.getDeclaredMethods().length > 1) {
      System.out.println(
        "The Printer class should only have 1 method named printArray."
      );
    }
  }
}
