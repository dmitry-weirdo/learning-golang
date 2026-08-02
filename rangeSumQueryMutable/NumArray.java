package rangeSumQueryMutable;

class NumArray {

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

    public NumArray(int[] nums) {
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

    public void update(int index, int val) {
        updateTree(root, index, val);
    }

    private void updateTree(TreeNode node, int index, int val) {
        if (node.left == node.right) { // reached the leaf node -> update its value if this is the correct node for [index; index]
            if (node.left == index) {
                node.sum = val;
            } else { // this must never happen if update is called with a valid index
                System.out.printf("FAILURE: reached leaf node [%s;%s], but its index != %s", node.left, node.right, index);
            }

            return;
        }

        var mid = (node.left + node.right) / 2;
        if (index > mid) { // go right, right is [mid+1, node.right]
            updateTree(node.rightChild, index, val);
        } else { // go left, left is [node.left; mid]
            updateTree(node.leftChild, index, val);
        }

        // update node (parent) sum
        node.sum = node.leftChild.sum + node.rightChild.sum;
    }

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

    static void sumRange(NumArray na, int left, int right, int expectedResult) {
        var result = na.sumRange(left, right);

        System.out.printf("Sum of indexes [%d; %d] = %d \n", left, right, result);
        System.out.printf("Expected result: %s \n", expectedResult);

        if (result != expectedResult) {
            System.out.printf("FAILURE: Expected result: %s, actual result: %s. \n", expectedResult, result);
        }
    }

    static void test() {
        int[] nums = {1, 3, 5};
        NumArray na = new NumArray(nums);

        sumRange(na, 0, 2, 9);
        na.update(1, 2); // change index 1 to 2 -> [1, 2, 5]
        sumRange(na, 0, 2, 8);
    }

    static void main() {
        // 307. Range Sum Query - Mutable
        test();
    }
}

/**
 * Your NumArray object will be instantiated and called as such:
 * NumArray obj = new NumArray(nums);
 * obj.update(index,val);
 * int param_2 = obj.sumRange(left,right);
 */