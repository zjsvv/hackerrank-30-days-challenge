import java.io.*;
import java.util.*;

public class Solution {

  private static int calculate(String[] returnedDate, String[] dueDate) {
    int d1 = Integer.parseInt(returnedDate[0]);
    int m1 = Integer.parseInt(returnedDate[1]);
    int y1 = Integer.parseInt(returnedDate[2]);
    int d2 = Integer.parseInt(dueDate[0]);
    int m2 = Integer.parseInt(dueDate[1]);
    int y2 = Integer.parseInt(dueDate[2]);

    if (y1 < y2) {
      return 0;
    } else if (y1 > y2) {
      return 10000;
    }

    // y1==y2 here
    if (m1 < m2) {
      return 0;
    } else if (m1 > m2) {
      return 500 * (m1 - m2);
    }

    // y1==y2 and m1==m2 here
    if (d1 < d2) {
      return 0;
    } else if (d1 > d2) {
      return 15 * (d1 - d2);
    }

    return 0;
  }

  public static void main(String[] args) {
    Scanner scanner = new Scanner(System.in);
    String returnedDateStr = scanner.nextLine();
    String dueDateStr = scanner.nextLine();

    String[] returnedDate = returnedDateStr.split(" ");
    String[] dueDate = dueDateStr.split(" ");

    System.out.println(calculate(returnedDate, dueDate));

    scanner.close();
  }
}
