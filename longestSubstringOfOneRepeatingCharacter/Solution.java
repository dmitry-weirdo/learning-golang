package longestSubstringOfOneRepeatingCharacter;

import java.util.Arrays;

class Solution {
    public int[] longestRepeating(String s, String queryCharacters, int[] queryIndices) {
        // s is lowercase English chars, so char will be ok, no Unicode

        SegmentTree tree = new SegmentTreeAsArray();
        tree.build(s.toCharArray());

        var queriesCount = queryIndices.length;

        int[] result = new int[queriesCount];

        for (int i = 0; i < queriesCount; i++) {
            var index = queryIndices[i];
            var character = queryCharacters.charAt(i); // char

            tree.update(index, character);

            var longestSingleCharacterSubstringLength = tree.getLongestSingleCharacterSubstringLength(0, s.length() - 1);

            result[i] = longestSingleCharacterSubstringLength;
        }

        return result;
    }

    private interface SegmentTree {
        void build(char[] nums);

        void update(int index, char value);

        int getLongestSingleCharacterSubstringLength(int left, int right);
    }

    private static class SegmentTreeAsArray implements SegmentTree {
        // array implementation of the binary tree
        // Array keeps the sums.
        // Navigating between parent and leftChild and rightChild is executed via arrayIndexes (standard formulae).
        // Ranges (left; right) are also inexplicitly calculated in array indexes.

        // a[0] is not used.
        // a[1] is root.
        // root i to children: leftChild = 2*i, rightChild = 2*i+1
        // child i to parent: parent = i/2

        // see https://cp-algorithms.com/data_structures/segment_tree.html#memory-efficient-implementation

        private static final int ROOT_INDEX = 1;
        private static final int ROOT_LEFT_RANGE = 0;

        private int n; // initial array length
        private int rootRightRange;

        // these are arrays of values for every segment
        private char[] chars; // just the character in the leaf segment // todo: this is not necessary for this task, it cannot be aggregatet above the leaf nodes
        private char[] leftChars; // leftmost character of the segment // todo: name startChars?
        private char[] rightChars; // leftmost character of the segment // todo: name endChars?

        private int[] longestPrefixLengths; // length of the longest single-char prefix substring starting from segment left
        private int[] longestSuffixLengths; // length of the longest single-char suffix substring ending on segment right
        private int[] maxSubstringLengths; // length of the longest single-char substring within the segment (it can be in any place of the segment)

        @Override
        public void build(char[] nums) {
            // build the array for the complete array
            n = nums.length;

            // every node has 2 children, so the actual used number of nodes will be within 2 * n.
            // However, the segment tree is NOT like heap, it has empty places not always at the left of the last level.
            // Therefore, using just 2 * n elements is not enough.

            // see https://cp-algorithms.com/data_structures/segment_tree.html#memory-efficient-implementation
            // for explanation of how we can change index calculation to just use 2 * n - 1 elements (element a[0] is not used).

            // For example, for n = 10, i = 12 range [5;6] ->
            // left child will be 2 * i = 24,
            // right child will be 2 * i + 1 = 25
            // This will be out of range of 2 * n - 1 = 19.
//            a = new int[4 * n];

            // for modified getLeftChildIndex, getRightChildIndex, we can allocate just 2 * n - 1 elements (element a[0] is not used).
            initArrays(2 * n);
//            chars = new char[2 * n];

            rootRightRange = n - 1;

            buildArray(nums, ROOT_INDEX, ROOT_LEFT_RANGE, rootRightRange);

            System.out.printf("Segment tree array: %s \n", Arrays.toString(chars));
        }

        private void initArrays(int size) {
            chars = new char[size];
            leftChars = new char[size];
            rightChars = new char[size];
            longestPrefixLengths = new int[size];
            longestSuffixLengths = new int[size];
            maxSubstringLengths = new int[size];
        }

        private void buildArray(char[] nums, int i, int left, int right) { // i is root index, left and right are segment indexes in the original array
            // todo: it's possible to build bottom-to-top iteratively,
            // but algorithm at https://leetcode.com/problems/range-sum-query-mutable/editorial/?envType=problem-list-v2&envId=segment-tree
            // is incorrect.

//            System.out.printf("Build array: i = %s, [left; right] = [%s; %s]. \n", i, left, right);

            // a[0] does not have a value.
            // a[1] is root.
            // Root i to children: leftChild = 2*i, rightChild = 2*i+1
            // Child i to parent: parent = i/2
            if (left == right) {
                // reached the leaf node -> set a[left] = a[right] to this node
                buildValuesForTheLeaf(i, nums[left]);
//                chars[i] = nums[left];

                return;
            }

            int mid = (left + right) / 2;

            // leftChild is sum[left; mid]
            int leftChildIndex = getLeftChildIndex(i);
            buildArray(nums, leftChildIndex, left, mid);

            // rightChild is sum[mid + 1; right]
            int rightChildIndex = getRightChildIndex(i, mid, left);
            buildArray(nums, rightChildIndex, mid + 1, right);

            // root is aggregate of leftChild and rightChild
            // update node (parent) aggregate data
//            chars[i] = chars[leftChildIndex] + chars[rightChildIndex];

            // todo: think about this length calculation
            int leftSegmentLength = mid - left + 1; // mid belongs to left
            int rightSegmentLength = right - mid; // mid does NOT belong to right

            updateParentNode(i, leftChildIndex, rightChildIndex, leftSegmentLength, rightSegmentLength);
//            chars[i] = chars[leftChildIndex] + chars[rightChildIndex];
        }

        private void buildValuesForTheLeaf(int i, char value) {
            // no, i != left
            // i is the segment-tree index,
            // indexInNums is index in the original array/string

//            if (i != indexInNums) {
//                throw new IllegalStateException(String.format("buildValuesForTheLeaf: i == %d, indexInNums == %d, i <> indexInNums.", i, indexInNums));
//            }

            chars[i] = value;
            leftChars[i] = value; // just 1 character, it's both left and right
            rightChars[i] = value; // just 1 character, it's both left and right

            // all lengths are 1 since it's just 1 single character
            longestPrefixLengths[i] = 1;
            longestSuffixLengths[i] = 1;
            maxSubstringLengths[i] = 1;
        }

        private void updateLeafNode(int i, char value) {
            chars[i] = value;
            leftChars[i] = value; // just 1 character, it's both left and right
            rightChars[i] = value; // just 1 character, it's both left and right

            // all lengths are always 1 -> nothing to  change
//            longestPrefixLengths[i] = 1;
//            longestSuffixLengths[i] = 1;
//            maxSubstringLengths[i] = 1;
        }

        private void updateParentNode(
                int i, // i - parent index 
                int leftChildIndex,
                int rightChildIndex,
                int leftLen, // length of the left segment
                int rightLen // length of the right segment
        ) {
            // parent.leftmostChar is left.leftmostChar
            // parent.rightmostChar is right.rightmostChar
            leftChars[i] = leftChars[leftChildIndex];
            rightChars[i] = rightChars[rightChildIndex];

            // initially, parent.longestPrefix = left.longestPrefix
            longestPrefixLengths[i] = longestPrefixLengths[leftChildIndex];
            if (
                    (longestPrefixLengths[leftChildIndex] == leftLen) && // left.longestPrefix is the complete left segment
                            (rightChars[leftChildIndex] == leftChars[rightChildIndex]) // left.lastChar = right.firstChar
            ) {
                // join the complete left segment and right.longestPrefix, since they are the same character
                longestPrefixLengths[i] = longestPrefixLengths[leftChildIndex] + longestPrefixLengths[rightChildIndex];
            }

            // initially, parent.longestSuffix = right.longestSuffix
            longestSuffixLengths[i] = longestSuffixLengths[rightChildIndex];
            if (
                    (longestSuffixLengths[rightChildIndex] == rightLen) // right.longestSuffix is the complete right segment
                            && (rightChars[leftChildIndex] == leftChars[rightChildIndex]) // left.lastChar = right.firstChar
            ) {
                // join the complete right segment and left.longestSuffix, since they are the same character
                longestSuffixLengths[i] = longestSuffixLengths[rightChildIndex] + longestSuffixLengths[leftChildIndex];
            }

            // initially, maxParent = max(maxLeft, maxRight)
            maxSubstringLengths[i] = Math.max(maxSubstringLengths[leftChildIndex], maxSubstringLengths[rightChildIndex]);
            if (rightChars[leftChildIndex] == leftChars[rightChildIndex]) {
                // If left ends on same char that right starts,
                // combine left.longestEndingSuffix and right.longestStartingSuffix

                // And this is a candidate for longest single-char substring
                int leftSuffixPlusRightPrefix = longestSuffixLengths[leftChildIndex] + longestPrefixLengths[rightChildIndex];

                maxSubstringLengths[i] = Math.max(maxSubstringLengths[i], leftSuffixPlusRightPrefix);
            }
        }

        @Override
        public void update(int index, char value) {
            // root is a sum[0; n - 1]
            updateArray(ROOT_INDEX, ROOT_LEFT_RANGE, rootRightRange, index, value);
        }

        private void updateArray(int i, int leftRange, int rightRange, int index, char value) {
            if (leftRange == rightRange) { // reached the leaf node -> update its value if this is the correct node for [index; index]
                if (leftRange == index) {
                    //chars[i] = value;
                    updateLeafNode(i, value);
                } else { // this must never happen if update is called with a valid index
                    System.out.printf("FAILURE: reached leaf node [%s;%s], but its index != %s", leftRange, rightRange, index);
                }

                return;
            }

            var midRange = (leftRange + rightRange) / 2;

            int leftChildIndex = getLeftChildIndex(i);
            int rightChildIndex = getRightChildIndex(i, midRange, leftRange);

            if (index > midRange) { // go right, right is [midRange+1, rightRange]
                updateArray(rightChildIndex, midRange + 1, rightRange, index, value);
            } else { // go left, left is [leftRange; midRange]
                updateArray(leftChildIndex, leftRange, midRange, index, value);
            }

            // todo: think about this length calculation
            int leftSegmentLength = midRange - leftRange + 1; // mid belongs to left
            int rightSegmentLength = rightRange - midRange; // mid does NOT belong to right

            // update node (parent) aggregate data
//            chars[i] = chars[leftChildIndex] + chars[rightChildIndex];
            updateParentNode(i, leftChildIndex, rightChildIndex, leftSegmentLength, rightSegmentLength);
        }

        @Override
        public int getLongestSingleCharacterSubstringLength(int left, int right) {
            // root corresponds to range [0; n - 1] (aggregates this range)
            return getLongestSingleCharacterSubstringLengthInArray(ROOT_INDEX, ROOT_LEFT_RANGE, rootRightRange, left, right);
        }

        private int getLongestSingleCharacterSubstringLengthInArray(int i, int leftRange, int rightRange, int left, int right) {
            // The trick is that the range of the current node a[i]
            // is passed as parameters to this function [leftRange; rightRange].
            // [left; right] is the range that we are looking for.
            if ((leftRange == left) && (rightRange == right)) {
//                return chars[i];

                return maxSubstringLengths[i];
            }

            var midRange = (leftRange + rightRange) / 2;

            int leftChildIndex = getLeftChildIndex(i);
            int rightChildIndex = getRightChildIndex(i, midRange, leftRange);

            if (left > midRange) {
                // both left and right are > than midRange -> go to the right subtree that is [midRange + 1; rightRange]
                return getLongestSingleCharacterSubstringLengthInArray(rightChildIndex, midRange + 1, rightRange, left, right);
            } else if (right <= midRange) {
                // both left and right are <= than midRange -> go to the left subtree that is [leftRange; midRange]
                return getLongestSingleCharacterSubstringLengthInArray(leftChildIndex, leftRange, midRange, left, right);
            } else {
                // left <= midRange < right
                // -> then the range [left; midRange] comes from the left subtree,
                // and range [midRange + 1; right] comes from the right subtree.
                return getLongestSingleCharacterSubstringLengthInArray(leftChildIndex, leftRange, midRange, left, midRange) +
                        getLongestSingleCharacterSubstringLengthInArray(rightChildIndex, midRange + 1, rightRange, midRange + 1, right);
            }
        }

        private static int getLeftChildIndex(int i) {
            // see https://cp-algorithms.com/data_structures/segment_tree.html#memory-efficient-implementation

            // We post the left part for array range [left; mid] right after the parent index [i],
            // i.e. left of parent [i] starts at [i + 1]
            // This part will serve for (mid - left + 1) array elements.
            // It will use 2 * (mid - left + 1) - 1 nodes.

            return i + 1; // this requires an array of size 2 * n

//            return 2 * i; // this requires an array of size 4 * n
        }

        private static int getRightChildIndex(int i, int midRange, int leftRange) {
            // see https://cp-algorithms.com/data_structures/segment_tree.html#memory-efficient-implementation

            // We post the right part for array range [mid + 1; right] right after the left part.
            // left start = i + 1
            // elements in left part = 2 * (mid - left + 1) - 1
            // i.e. right start = i + 1 + 2 * (mid - left + 1) - 1 = i + 2 * (mid - left + 1)
            return i + 2 * (midRange - leftRange + 1); // this requires an array of size 2 * n

//            return 2 * i + 1; // this requires an array of size 4 * n
        }
    }

    static void main() {
        // 2213. Longest Substring of One Repeating Character

        // Segment tree array implementation copied from "307. Range Sum Query - Mutable"
        // What needs to be changed for this ticket -> we're not tracking the sum,
        // we're tracking multiple properties for the longest single-character calculation,
        // i.e. instead of a[] with sums, we have multiple arrays (every property of node[i] is in propertyArray[i]).

        // Also, the tree is built on a string, not int[].

        var s = new Solution();
        int[] result = s.longestRepeating("babacc", "bcb", new int[]{1, 3, 3});
        System.out.println("Result: " + Arrays.toString(result));

        // todo: add tests
        // todo: understand about lengths

        // todo: try to rewrite the SegmentTree in Go, and check whether it will be faster. The current solution is slow, around 140 ms :(
    }
}