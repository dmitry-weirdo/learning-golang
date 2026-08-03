package stoneGameIII;

import java.util.Arrays;

class Solution {
    private int[] a;
    private int[] dp;
    private int lastIndex;

    public String stoneGameIII(int[] nums) {
        a = nums;

        lastIndex = nums.length - 1;

        // optimal difference in score when playing in the array range of [left; right]
        // right is always end of array, so we can use just a 1D-array
        dp = new int[nums.length];

        // check whether result for the whole list of nums is non-negative
        // DFS returns the diff between score of current player and opponent, when they both play optimally
        int dfsResult = dfs(0);

        if (dfsResult > 0) {
            return "Alice";
        } else if (dfsResult < 0) {
            return "Bob";
        } else {
            return "Tie";
        }
    }

    private int dfs(int left) {
        if (left > lastIndex) { // base case - no elements anymore
            return 0;
        }

        if (left == lastIndex) { // just 1 element left -> we must take it
            return a[left];
        }

        if (dp[left] != 0) { // result already found -> return from cache
            return dp[left];
        }

        // we take 1 from left -> opponent plays best in [left + 1; right] range
        int take1Left = a[left] - dfs(left + 1);

        // we take 2 from right -> opponent plays best in [left + 2; right] range
        int take2Left = take1Left; // values can be negative, so taking 0 is not a valid minimum
        if (left + 1 <= lastIndex) {
            take2Left = a[left] + a[left + 1] - dfs(left + 2);
        }

        // we take 3 from right -> opponent plays best in [left + 3; right] range
        int take3Left = take1Left;
        if (left + 2 <= lastIndex) { // values can be negative, so taking 0 is not a valid minimum
            take3Left = a[left] + a[left + 1] + a[left + 2] - dfs(left + 3);
        }

        // our optimal score is the best diff when we take either 1, 2 or 3 from left
        dp[left] = Math.max(Math.max(take1Left, take2Left), take3Left);

        return dp[left];
    }

    static void test(int[] nums, String expectedResult) {
        System.out.println();
        System.out.println("========================");

        System.out.printf("Nums: %s \n", Arrays.toString(nums));

        var result = new Solution().stoneGameIII(nums);
        System.out.printf("Winner: %s \n", result);
        System.out.printf("Expected result: %s \n", expectedResult);

        if (!result.equals(expectedResult)) {
            System.out.printf("FAILURE: expected result = %s, actual result = %s \n", expectedResult, result);
        }
    }

    static void test1() {
        var nums = new int[]{1, 2, 3, 7};
        var expected = "Bob";

        test(nums, expected);
    }

    static void test2() {
        var nums = new int[]{1, 2, 3, -9};
        var expected = "Alice";

        test(nums, expected);
    }

    static void test3() {
        var nums = new int[]{1, 2, 3, 6};
        var expected = "Tie";

        test(nums, expected);
    }

    static void main() {
        // 1406. Stone Game III
        test1();
        test2();
        test3();
    }
}