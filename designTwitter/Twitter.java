package designTwitter;

// import java.util.Deque;

import java.util.*;

class Twitter {
    record Tweet(int tweetId, int authorId, int order) {}
    record TweetAndNextTweet(Tweet tweet, Iterator<Tweet> iterator) {}

    private static final int COUNT = 10;
    private int order = 0;

    // tweets by user (every author has its list of 10 tweets)
    private final Map<Integer, Deque<Tweet>> tweetsByUser;

    private final Map<Integer, Set<Integer>> followeesByUser;

    public Twitter() {
        tweetsByUser = new HashMap<>();
        followeesByUser = new HashMap<>();
    }

    public void postTweet(int userId, int tweetId) {
        order++;

        var tweet = new Tweet(tweetId, userId, order);

        var userTweets = getTweetsByUser(userId);
        userTweets.addFirst(tweet);

        while (userTweets.size() > COUNT) {
            userTweets.removeLast();
        }
    }

    private Deque<Tweet> getTweetsByUser(int userId) {
        // todo: maybe there is a dequeue that automatically controls its size
        tweetsByUser.computeIfAbsent(userId, _ -> new LinkedList<>());

        return tweetsByUser.get(userId);
    }

    public List<Integer> getNewsFeed(int userId) {
        return getNewsFeedByTopUsers(userId);

/*
        PriorityQueue<Tweet> pq = new PriorityQueue<>(
                Comparator.comparingInt(t -> t.order)
        );

        var followees = getFolloweesByUser(userId); // will include the user himself

        // todo: if there are more than 10 followers, we can just select the top 10 followers with latest tweets and then execute the "merge K sorted" lists logic with a max-heap

        // todo: maybe foreach
        for (var followee : followees) {
            var tweets = getTweetsByUser(followee);
            pq.addAll(tweets);

            // we only need 10 in the pq
            while (pq.size() > COUNT) {
                pq.poll();
            }
        }

        // convert to list and reverse
        List<Integer> result = new ArrayList<>();

        while (!pq.isEmpty()) {
            result.add(pq.poll().tweetId);
        }

        // pq was min-heap "from oldest to newest", feed must be "from newest to oldest".
        Collections.reverse(result);

        return result;
*/
    }

    private List<Integer> getNewsFeedByTopUsers(int userId) {
        Comparator<TweetAndNextTweet> comparator = Comparator.comparingInt((TweetAndNextTweet tn) -> tn.tweet().order()).reversed();
        PriorityQueue<TweetAndNextTweet> pqFirstTweetsByUser = new PriorityQueue<>(comparator);

        // add users with top 10 first tweets. These 10 tweets will be better than any of the tweets of the remaining followees.
        // If there are less than 10 tweets (i.e. less than 10 followee users with tweets), fewer elements will be collected,
        // but we will still collect the following tweets by getting next tweets of these users.
        var followees = getFolloweesByUser(userId); // will include the user himself

        for (var followee : followees) {
            Deque<Tweet> tweets = getTweetsByUser(followee);
            Iterator<Tweet> iterator = tweets.iterator();

            if (iterator.hasNext()) { // followee can have no tweets
                TweetAndNextTweet tweetAndNextTweet = new TweetAndNextTweet(iterator.next(), iterator);
                pqFirstTweetsByUser.add(tweetAndNextTweet);
            }
        }

        // take just the top 10 (or fewer) followees from the pq (10 best users by their latest tweets)
        PriorityQueue<TweetAndNextTweet> pq = new PriorityQueue<>(comparator);

        while ((pq.size() < COUNT) && !pqFirstTweetsByUser.isEmpty()) {
            pq.add(pqFirstTweetsByUser.poll());
        }

        List<Integer> result = new ArrayList<>();

        while ((result.size() < COUNT) && (!pq.isEmpty())) {
            TweetAndNextTweet tn = pq.poll();

            result.add(tn.tweet.tweetId);

            if (tn.iterator.hasNext()) {
                // if exists -> add the next tweet from this followee
                // it will take its place in the heap
                Tweet nextTweet = tn.iterator.next();
                TweetAndNextTweet tweetAndNextTweet = new TweetAndNextTweet(nextTweet, tn.iterator);

                pq.add(tweetAndNextTweet);
            }
        }

        return result;
    }

    public void follow(int userId, int followeeId) {
        // todo: throw if userId == followeeId
        getFolloweesByUser(userId).add(followeeId);
    }

    public void unfollow(int userId, int followeeId) {
        // todo: throw if userId == followeeId
        getFolloweesByUser(userId).remove(followeeId);
    }

    private Set<Integer> getFolloweesByUser(int userId) {
        followeesByUser.computeIfAbsent(userId, _ -> new HashSet<>());

        // put the user himself to his/her followees
        followeesByUser.get(userId).add(userId);

        return followeesByUser.get(userId);
    }

    private static void getNewsFeed(Twitter twitter, int userId) {
        List<Integer> feed = twitter.getNewsFeed(userId);

        System.out.printf("Feed of user %d: %s\n", userId, feed);
    }

    static void main() {
        // 355. Design Twitter

        Twitter twitter = new Twitter();
        twitter.postTweet(1, 10); // User 1 posts a new tweet with id = 10.
        twitter.postTweet(2, 20); // User 2 posts a new tweet with id = 20.

        getNewsFeed(twitter, 1); // User 1's news feed should only contain their own tweets -> [10].
        getNewsFeed(twitter, 2); // User 2's news feed should only contain their own tweets -> [20].


//        twitter.getNewsFeed(1);   // User 1's news feed should only contain their own tweets -> [10].
//        twitter.getNewsFeed(2);   // User 2's news feed should only contain their own tweets -> [20].

        twitter.follow(1, 2);     // User 1 follows user 2.
        getNewsFeed(twitter, 1); // User 1's news feed should contain both tweets from user 1 and user 2 -> [20, 10].
        getNewsFeed(twitter, 2); // User 2's news feed should still only contain their own tweets -> [20].

//       twitter.getNewsFeed(1);   // User 1's news feed should contain both tweets from user 1 and user 2 -> [20, 10].
//        twitter.getNewsFeed(2);   // User 2's news feed should still only contain their own tweets -> [20].

        twitter.unfollow(1, 2);   // User 1 unfollows user 2.
        getNewsFeed(twitter, 1); // User 1's news feed should only contain their own tweets -> [10].

//        twitter.getNewsFeed(1);
    }
}

/**
 * Your Twitter object will be instantiated and called as such:
 * Twitter obj = new Twitter();
 * obj.postTweet(userId,tweetId);
 * List<Integer> param_2 = obj.getNewsFeed(userId);
 * obj.follow(followerId,followeeId);
 * obj.unfollow(followerId,followeeId);
 */