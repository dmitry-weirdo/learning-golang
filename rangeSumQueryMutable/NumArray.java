package rangeSumQueryMutable;

import java.util.Arrays;

class NumArray {

    private interface SegmentTree {
        void build(int[] nums);

        void update(int index, int value);

        int sumRange(int left, int right);
    }

    private static class SegmentTreeAsTree implements SegmentTree {
        // tree implementation of Segment Tree
        // see https://neetcode.io/courses/advanced-algorithms/8

        // since we need non-final fields (we change values of nodes) -> we cannot use record
        private class TreeNode {
            TreeNode leftChild;
            TreeNode rightChild;
            int left; // left range, inclusive
            int right; // right range, inclusive
            int sum; // sum of [left; right]

            public TreeNode(TreeNode leftChild, TreeNode rightChild, int left, int right, int sum) {
                this.leftChild = leftChild;
                this.rightChild = rightChild;
                this.left = left;
                this.right = right;
                this.sum = sum;
            }
        }

        private TreeNode root;

        @Override
        public void build(int[] nums) {
            // build the tree for the complete array
            root = buildTree(nums, 0, nums.length - 1);
        }

        private TreeNode buildTree(int[] nums, int left, int right) {
            if (left == right) { // reached the leaf node -> create it (no children, sum = a[sum])
                return new TreeNode(null, null, left, right, nums[left]);
            }

            int mid = (left + right) / 2;

            // leftChild is sum [left; mid]
            var leftChild = buildTree(nums, left, mid);

            // rightChild is sum [mid + 1; right]
            var rightChild = buildTree(nums, mid + 1, right);

            var sum = leftChild.sum + rightChild.sum;

            // create root
            return new TreeNode(leftChild, rightChild, left, right, sum);
        }

        @Override
        public void update(int index, int value) {
            updateTree(root, index, value);
        }

        private void updateTree(TreeNode node, int index, int value) {
            if (node.left == node.right) { // reached the leaf node -> update its value if this is the correct node for [index; index]
                if (node.left == index) {
                    node.sum = value;
                } else { // this must never happen if update is called with a valid index
                    System.out.printf("FAILURE: reached leaf node [%s;%s], but its index != %s", node.left, node.right, index);
                }

                return;
            }

            var mid = (node.left + node.right) / 2;
            if (index > mid) { // go right, right is [mid+1, node.right]
                updateTree(node.rightChild, index, value);
            } else { // go left, left is [node.left; mid]
                updateTree(node.leftChild, index, value);
            }

            // update node (parent) sum
            node.sum = node.leftChild.sum + node.rightChild.sum;
        }

        @Override
        public int sumRange(int left, int right) {
            return sumRangeInTree(root, left, right);
        }

        private int sumRangeInTree(TreeNode node, int left, int right) {
            if ((node.left == left) && (node.right == right)) {
                // we hit exactly the range [left; right] -> return sum of this node, it's already pre-calculated
                return node.sum;
            }

            var mid = (node.left + node.right) / 2;

            if (left > mid) {
                // both left and right are > than mid -> go to the right subtree that is [mid + 1; node.right]
                return sumRangeInTree(node.rightChild, left, right);
            } else if (right <= mid) {
                // both left and right are <= than mid -> go to the left subtree that is [node.left; mid]
                return sumRangeInTree(node.leftChild, left, right);
            } else {
                // left <= mid < right
                // -> then the range [left; mid] comes from the left subtree,
                // and range [mid + 1; right] comes from the right subtree.
                return sumRangeInTree(node.leftChild, left, mid) +
                        sumRangeInTree(node.rightChild, mid + 1, right);
            }
        }
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

        private int[] a;
        private int n; // initial array length
        private int rootRightRange;

        @Override
        public void build(int[] nums) {
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
            a = new int[4 * n];

            rootRightRange = n - 1;

            buildArray(nums, ROOT_INDEX, ROOT_LEFT_RANGE, rootRightRange);

            System.out.printf("Segment tree array: %s \n", Arrays.toString(a));
        }

        private void buildArray(int[] nums, int i, int left, int right) {
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
                a[i] = nums[left];
                return;
            }

            int mid = (left + right) / 2;

            // leftChild is sum [left; mid]
            int leftChildIndex = 2 * i;
            buildArray(nums, leftChildIndex, left, mid);

            // rightChild is sum[mid + 1; right]
            int rightChildIndex = 2 * i + 1;
            buildArray(nums, rightChildIndex, mid + 1, right);

            // root is sum of leftChild and rightChild
            a[i] = a[leftChildIndex] + a[rightChildIndex];
        }

        @Override
        public void update(int index, int value) {
            // root is a sum[0; n - 1]
            updateArray(ROOT_INDEX, ROOT_LEFT_RANGE, rootRightRange, index, value);
        }

        private void updateArray(int i, int leftRange, int rightRange, int index, int value) {
            if (leftRange == rightRange) { // reached the leaf node -> update its value if this is the correct node for [index; index]
                if (leftRange == index) {
                    a[i] = value;
                } else { // this must never happen if update is called with a valid index
                    System.out.printf("FAILURE: reached leaf node [%s;%s], but its index != %s", leftRange, rightRange, index);
                }

                return;
            }

            var midRange = (leftRange + rightRange) / 2;

            int leftChildIndex = 2 * i;
            int rightChildIndex = 2 * i + 1;

            if (index > midRange) { // go right, right is [midRange+1, rightRange]
                updateArray(rightChildIndex, midRange + 1, rightRange, index, value);
            } else { // go left, left is [leftRange; midRange]
                updateArray(leftChildIndex, leftRange, midRange, index, value);
            }

            // update node (parent) sum
            a[i] = a[leftChildIndex] + a[rightChildIndex];
        }

        @Override
        public int sumRange(int left, int right) {
            // root is a sum[0; n - 1]
            return sumRangeInArray(ROOT_INDEX, ROOT_LEFT_RANGE, rootRightRange, left, right);
        }

        private int sumRangeInArray(int i, int leftRange, int rightRange, int left, int right) {
            // The trick is that the range of the current node a[i]
            // is passed as parameters to this function [leftRange; rightRange].
            // [left; right] is the range that we are looking for.
            if ((leftRange == left) && (rightRange == right)) {
                return a[i];
            }

            var midRange = (leftRange + rightRange) / 2;

            int leftChildIndex = 2 * i;
            int rightChildIndex = 2 * i + 1;

            if (left > midRange) {
                // both left and right are > than midRange -> go to the right subtree that is [midRange + 1; rightRange]
                return sumRangeInArray(rightChildIndex, midRange + 1, rightRange, left, right);
            } else if (right <= midRange) {
                // both left and right are <= than midRange -> go to the left subtree that is [leftRange; midRange]
                return sumRangeInArray(leftChildIndex, leftRange, midRange, left, right);
            } else {
                // left <= midRange < right
                // -> then the range [left; midRange] comes from the left subtree,
                // and range [midRange + 1; right] comes from the right subtree.
                return sumRangeInArray(leftChildIndex, leftRange, midRange, left, midRange) +
                        sumRangeInArray(rightChildIndex, midRange + 1, rightRange, midRange + 1, right);
            }
        }
    }

    private SegmentTree implementation;

    public NumArray(int[] nums) {
        boolean tree = false;

        if (tree) {
            implementation = new SegmentTreeAsTree();
        } else {
            implementation = new SegmentTreeAsArray();
        }

        implementation.build(nums);
    }

    public void update(int index, int val) {
        implementation.update(index, val);
    }

    public int sumRange(int left, int right) {
        return implementation.sumRange(left, right);
    }

    static void sumRange(NumArray na, int left, int right, int expectedResult) {
        var result = na.sumRange(left, right);

        System.out.printf("Sum of indexes [%d; %d] = %d \n", left, right, result);
        System.out.printf("Expected result: %s \n", expectedResult);

        if (result != expectedResult) {
            System.out.printf("FAILURE: Expected result: %s, actual result: %s. \n", expectedResult, result);
        }
    }

    static void test1() {
        int[] nums = {1, 3, 5};
        NumArray na = new NumArray(nums);

        sumRange(na, 0, 2, 9);
        na.update(1, 2); // change index 1 to 2 -> [1, 2, 5]
        sumRange(na, 0, 2, 8);
    }

    static void test2() {
        // array of size 10 will NOT fit into 2 * n array if using the standard indexing
        int[] nums = {-28, -39, 53, 65, 11, -56, -65, -39, -43, 97};

        NumArray na = new NumArray(nums);
        sumRange(na, 0, 9, -44);
    }

    static void main() {
        // 307. Range Sum Query - Mutable
        test1();
        test2();
    }
}

