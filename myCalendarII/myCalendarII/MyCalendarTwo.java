package myCalendarII;

import java.util.TreeMap;

class MyCalendarTwo {
    // we're using treeMap for quick insertion and O(log N) search of the value
    private final TreeMap<Integer, Integer> bookingCount; // timepoint -> bookings on it. +1 on interval start, -1 on interval end
    private final int maxOverlappedBookings;

    public MyCalendarTwo() {
        bookingCount = new TreeMap<>();
        maxOverlappedBookings = 2;
    }

    public boolean book(int startTime, int endTime) {
        Integer currentStartValue = bookingCount.getOrDefault(startTime, 0);
        Integer currentEndValue = bookingCount.getOrDefault(endTime, 0);

        // Interval start increases +1
        bookingCount.put(startTime, currentStartValue + 1);

        // Interval end decreases -1
        bookingCount.put(endTime, currentEndValue - 1);

        // This is a Line Sweep algorithm, also used in "Zero Array Transformation I"
        // calc the prefix sums of all intervals
        // If prefixSum >= 2 (allowed intersections), fail the operation
        int prefixSum = 0;

        for (Integer v : bookingCount.values()) {
            prefixSum += v;

            if (prefixSum > maxOverlappedBookings) { // some interval has more bookings than allowed
                // rollback the changes for [start; end)
                bookingCount.put(startTime, currentStartValue);
                bookingCount.put(endTime, currentEndValue);

                // if start or end values are 0, remove them from the map
                if (bookingCount.get(startTime) == 0) {
                    bookingCount.remove(startTime);
                }

                if (bookingCount.get(endTime) == 0) {
                    bookingCount.remove(endTime);
                }

                return false;
            }
        }

        // no over-bookings found -> interval added successfully return true
        return true;
    }

    static void bookAndExpect(MyCalendarTwo cal, int start, int end, boolean expectedResult) {
        boolean result = cal.book(start, end);
        if (result) {
            System.out.printf("Successfully booked interval [%d, %d]. No intersections over 2 bookings found. \n", start, end);
        } else {
            System.out.printf("Failed to book interval [%d, %d] because of over 2 bookings intersection. \n", start, end);
        }

        if (result != expectedResult) {
            System.out.printf("FAILURE: Expected result: %s, actual result: %s. \n", expectedResult, result);
        }
    }

    static void test1() {
        // ["MyCalendarTwo", "book", "book", "book", "book", "book", "book"]
        //[[], [10, 20], [50, 60], [10, 40], [5, 15], [5, 10], [25, 55]]
        //Output
        //[null, true, true, true, false, true, true]

        MyCalendarTwo c = new MyCalendarTwo();
        bookAndExpect(c, 10, 20, true); // [10, 20)
        bookAndExpect(c, 50, 60, true); // [10, 20), [50, 60)
        bookAndExpect(c, 10, 40, true); // [10, 20) - intersection, [20, 40), [50, 60)
        bookAndExpect(c, 5, 15, false); // [10, 15) will intersect 3 -> fail
        bookAndExpect(c, 5, 10, true); // [5, 10), [10, 20) - intersection, [20, 40), [50, 60)
        bookAndExpect(c, 25, 55, true); // [5, 10), [10, 20) - intersection, [25, 40) - intersection, [40, 50), [50, 55) - intersection, [55, 60)
    }

    static void main() {
        // 731. My Calendar II
        test1();
    }
}