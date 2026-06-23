import java.util.Arrays;

class Solution992 {
    public int subarraysWithKDistinct(int[] nums, int k) {
        // Get the leftmost valid position for each index with at most k distinct elements
        int[] leftmostWithAtMostK = findLeftmostPositions(nums, k);
      
        // Get the leftmost valid position for each index with at most k-1 distinct elements
        int[] leftmostWithAtMostKMinus1 = findLeftmostPositions(nums, k - 1);
      
        // Calculate the result using the difference
        // For each ending position i, the number of subarrays with exactly k distinct elements
        // equals (subarrays with at most k) - (subarrays with at most k-1)
        int result = 0;
        for (int i = 0; i < nums.length; i++) {
            result += leftmostWithAtMostKMinus1[i] - leftmostWithAtMostK[i];
        }
        return result;
    }

    /**
     * Finds the leftmost starting position for each ending position
     * such that the subarray has at most k distinct elements.
     * 
     * @param nums the input array
     * @param k maximum number of distinct elements allowed
     * @return array where positions[i] represents the leftmost valid starting index
     *         for subarrays ending at index i with at most k distinct elements
     */
    private int[] findLeftmostPositions(int[] nums, int k) {
        int n = nums.length;
      
        // Frequency counter for each element (array size is n+1 to handle all possible values)
        int[] frequency = new int[n + 1];
      
        // Result array storing leftmost valid position for each ending index
        int[] positions = new int[n];
      
        // Counter for number of distinct elements in current window
        int distinctCount = 0;
      
        // Two pointers: left (j) and right (i)
        int left = 0;
      
        for (int right = 0; right < n; right++) {
            // Add current element to the window
            if (++frequency[nums[right]] == 1) {
                // If this element appears for the first time, increment distinct count
                distinctCount++;
            }
          
            // Shrink window from left while we have more than k distinct elements
            while (distinctCount > k) {
                // Remove leftmost element from the window
                if (--frequency[nums[left]] == 0) {
                    // If this element's count becomes 0, decrement distinct count
                    distinctCount--;
                }
                left++;
            }
          
            // Store the leftmost valid position for current ending position
            positions[right] = left;
        }
      
        return positions;
    }

    static void main() {
        Solution992 s = new Solution992();

        int[] nums = {1, 2, 1, 2, 3};
        int k = 2;

        int result = s.subarraysWithKDistinct(nums, k);

        System.out.println("nums: " + Arrays.toString(nums));
        System.out.println("k: " + k);
        System.out.println("result: " + result);
    }
}