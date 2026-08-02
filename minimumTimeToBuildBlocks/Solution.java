package minimumTimeToBuildBlocks;

import java.util.Arrays;
import java.util.PriorityQueue;

class Solution {

    public int minBuildTime(int[] blocks, int split) {
        // We're trying to minimize the splits for the current max block,
        // but this is a _current_ max block, not the initial max block.

        // Executing 2 blocks is the equivalent to the block of [ max(block1, block2) + split time ].

        // Create a min-heap to store block build times
        // We use this to always process the smallest times first
        PriorityQueue<Integer> minHeap = new PriorityQueue<>();

        // Add all block build times to the heap
        for (int blockTime : blocks) {
            minHeap.offer(blockTime);
        }

        // todo: remove this
 /*       System.out.println("Polling heap");
        while (!minHeap.isEmpty()) {
            System.out.printf("%s ", minHeap.poll());
        }
        System.out.println();

        if (true) {
            return -1;
        }
*/
        // Keep merging blocks until only one remains
        // This simulates the process of workers splitting and building
        while (minHeap.size() > 1) {
            int firstSmallest = minHeap.poll();
            // Remove the smallest build time (this block gets built in parallel)

            // Take the second smallest and add split time
            // This represents the time for a worker to split and then build
            int secondSmallest = minHeap.poll();
            minHeap.offer(secondSmallest + split);

            System.out.printf("2 smallest: %d, %d. Added %d to the heap. \n", firstSmallest, secondSmallest, secondSmallest + split);
        }

        // Return the final remaining time, which is the minimum total build time
        return minHeap.poll();
    }

    public int minBuildTime_bad(int[] blocks, int split) {
        // this is incorrect since sometimes we need to split multiple workers,
        // instead of always using one worker for the max job of the initial blocks.

        // [1, 1, 1, 1], split = 100
        // split into 2 workers -> time 100
        // split 2 workers into 4 workers -> time 200
        // run all 4 jobs in parallel -> time 201

        if (blocks.length == 0) {
            // this should never happen
            return 0;
        }

        if (blocks.length == 1) {
            // corner-case -> no need to split, just run one job with one worker
            return blocks[0];
        }

        // the logic is - after every split, we execute the max block to minimize its split
        // we split the worker and execute the max job from the current time

        // For the last 2 blocks, we have 2 split workers and just execute these 2 jobs
        // from the current timestamp

        // we need to know the max block
        Arrays.sort(blocks);

        // minimum time for every block
        var times = new int[blocks.length];

        // calculate times for all the blocks
        var time = split; // time including all splits

        for (int i = blocks.length - 1; i >= 1; i--) {
            // last 2 blocks -> we have 2 workers, do not split them anymore
            if (i == 1) {
                times[0] = time + blocks[0];
                times[1] = time + blocks[1];
                break;
            }

            // from the 2 current split workers, 1 of them executes the maximum block
            times[i] = time + blocks[i];

            // split the remaining worker into 2 workers
            time += split;
        }

        // return the max time of every blocks
        var max = 0;

        for (int t : times) {
            max = Math.max(max, t);
        }

        return max;
    }

    static void test(int[] blocks, int split, int expectedResult) {
        System.out.println();
        System.out.println("========================");

        System.out.printf("Blocks: %s \n", Arrays.toString(blocks));
        System.out.printf("Split time: %s \n", split);

        var result = new Solution().minBuildTime(blocks, split);
        System.out.printf("Min time required: %s \n", result);
        System.out.printf("Expected result: %s \n", expectedResult);

        if (result != expectedResult) {
            System.out.printf("FAILURE: expected result = %s, actual result = %s \n", expectedResult, result);
        }
    }

    static void test1() {
        var blocks = new int[]{1};
        var split = 1;
        var expectedResult = 1;

        test(blocks, split, expectedResult);
    }

    static void test2() {
        var blocks = new int[]{1, 2};
        var split = 5;
        var expectedResult = 7;

        test(blocks, split, expectedResult);
    }

    static void test3() {
        var blocks = new int[]{1, 2, 3};
        var split = 1;
        var expectedResult = 4;

        test(blocks, split, expectedResult);
    }

    static void test4() {
        var blocks = new int[]{1, 1, 1, 1};
        var split = 100;
//        var expectedResult = 301;
        var expectedResult = 201; // ?????

        test(blocks, split, expectedResult);
    }

    static void main() {
        // 1199. Minimum Time to Build Blocks
        test1();
        test2();
        test3();
        test4();
    }
}