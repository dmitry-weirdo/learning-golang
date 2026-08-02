package myCalendarI;

import java.util.TreeMap;

public class MyCalendar {

    // maps interval start to interval end
    private TreeMap<Integer, Integer> calendar;

    public MyCalendar() {
        calendar = new TreeMap<>();
    }

    public boolean book(int startTime, int endTime) {
        // The intervals [s1; e1] and [s2; e2] do not intersect
        // when (e1 <= s2) or (e2 <= s1).

        // So we search the last interval starting BEFORE or equal to start,
        // and the first interval starting AFTER or equal to start.
        // Search in the balanced BST (TreeMap is a red-black tree) - O(log N) operation
        Integer prevStart = calendar.floorKey(startTime);
        Integer nextStart = calendar.ceilingKey(startTime);

        // Now we check that both of these intervals either do not exist
        // or they don't intersect with [start; end].
        // (prevEnd <= start) and (end <= nextStart)

        boolean prevDoesNotIntersect;
        if (prevStart == null) { // no interval before start
            prevDoesNotIntersect = true;
        } else {
            Integer prevEnd = calendar.get(prevStart);
            prevDoesNotIntersect = (prevEnd <= startTime);
        }

        boolean nextDoesNotIntersect;
        if (nextStart == null) { // no interval after end
            nextDoesNotIntersect = true;
        } else {
            nextDoesNotIntersect = (endTime <= nextStart);
        }

        if (prevDoesNotIntersect && nextDoesNotIntersect) {
            // no intersections with given interval -> add it to the tree - O(log N) operation
            calendar.put(startTime, endTime);
        }

        return prevDoesNotIntersect && nextDoesNotIntersect;
    }

    static void bookAndExpect(MyCalendar cal, int start, int end, boolean expectedResult) {
        boolean result = cal.book(start, end);
        if (result) {
            System.out.printf("Successfully booked interval [%d, %d]. No intersections found. \n", start, end);
        } else {
            System.out.printf("Failed to book interval [%d, %d] because of the intersection. \n", start, end);
        }

        if (result != expectedResult) {
            System.out.printf("FAILURE: Expected result: %s, actual result: %s. \n", expectedResult, result);
        }
    }

    static void test() {
        MyCalendar c = new MyCalendar();
        bookAndExpect(c, 10, 20, true);
        bookAndExpect(c, 15, 25, false);
        bookAndExpect(c, 20, 30, true);
    }

    static void main() {
        // 729. My Calendar I
        test();
    }
}
