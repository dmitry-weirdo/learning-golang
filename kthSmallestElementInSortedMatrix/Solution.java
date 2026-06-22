public class Solution {
    public static int kthSmallest(int[][] matrix, int k) {
        int n = matrix.length;
        int left = matrix[0][0];
        int right = matrix[n - 1][n - 1];
        int firstTrueIndex = -1;

        while (left <= right) {
            int mid = left + (right - left) / 2;
            if (countLessEqual(matrix, mid, n) >= k) {
                firstTrueIndex = mid;
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }

        return firstTrueIndex;
    }

    /**
     * Counts elements <= target using staircase search from bottom-left.
     */
    private static int countLessEqual(int[][] matrix, int target, int n) {
        int count = 0;
        int row = n - 1;
        int col = 0;

        System.out.printf("Searching for value %s \n", target);
        System.out.printf("i, j: [%s, %s] \n", row, col);

        while (row >= 0 && col < n) {
            if (matrix[row][col] <= target) {
                count += row + 1;
                col++;
            } else {
                row--;
            }

            System.out.printf("i, j: [%s, %s] \n", row, col);
        }

        System.out.printf("Returning count = %s \n", count);
        System.out.println();

        return count;
    }

    static void main() {
//        test1();

        test2();
    }

    private static void test2() {
        int[] row0 = { -5, -4 };
        int[] row1 = { -5, -4 };
        int[][] matrix = { row0, row1 };

        int k = 2;

        int result = kthSmallest(matrix, k);

        System.out.printf("%s-th value: %s  \n", k, result);
    }

    private static void test1() {
        int[] row0 = { 1, 2 };
        int[] row1 = { 1, 3 };
        int[][] matrix = { row0, row1 };

        int k = 3;

        kthSmallest(matrix, k);
    }
}