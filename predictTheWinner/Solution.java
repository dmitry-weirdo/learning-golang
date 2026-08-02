package predictTheWinner;

import java.util.Arrays;

class Solution {
    private int[] a;
    private int[][] dp;

    // For 877. Stone Game, just use a different function name.
    // public boolean stoneGame(int[] nums) { // 877. Stone Game
    public boolean predictTheWinner(int[] nums) { // 486. Predict the Winner
        a = nums;

        // optimal difference in score when playing in the array range of [left; right]
        dp = new int[nums.length][nums.length];

        // check whether result for the whole list of nums is non-negative
        // DFS returns the diff between score of current player and opponent, when they both play optimally
        return dfs(0, nums.length - 1) >= 0;
    }

    private int dfs(int left, int right) {
        if (left > right) { // base case - no elements anymore
            return 0;
        }

        if (left == right) {
            return a[left];
        }

        if (dp[left][right] != 0) { // result already found -> return from cache
            return dp[left][right];
        }

        // we take from left -> opponent plays best in [left + 1; right] range
        int takeLeft = a[left] - dfs(left + 1, right);

        // we take from right -> opponent plays best in [left; right - 1] range
        int takeRight = a[right] - dfs(left, right - 1);

        // our optimal score is the best diff when we take either from left or from right
        dp[left][right] = Math.max(takeLeft, takeRight);

        return dp[left][right];
    }

    static void test(int[] nums, boolean expectedResult) {
        System.out.println();
        System.out.println("========================");

        System.out.printf("Nums: %s \n", Arrays.toString(nums));

        var result = new Solution().predictTheWinner(nums);
        System.out.printf("Player 1 will win: %s \n", result);
        System.out.printf("Expected result: %s \n", expectedResult);

        if (result != expectedResult) {
            System.out.printf("FAILURE: expected result = %s, actual result = %s \n", expectedResult, result);
        }
    }

    static void test1() {
        var nums = new int[]{1, 5, 2};
        var expected = false;

        test(nums, expected);
    }

    static void test2() {
        var nums = new int[]{1, 5, 233, 7};
        var expected = true;

        test(nums, expected);
    }

    static void main() {
        // 486. Predict the Winner
        // 877. Stone Game - exactly same solution is working
        test1();
        test2();
    }
}