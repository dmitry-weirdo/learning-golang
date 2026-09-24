-- 550. Game Play Analysis IV
select
round(
    q1.target_players::numeric / q2.total_players, -- cast to numeric to have a double division instead of integer division
    2 -- round to 2 decimal digits
) as fraction
from
( -- q1 - select count of players who have 2nd game the next date after the 1st game
    select
    count(x.player_id) as target_players
    --x.event_date as first_date, -- earliest date
    --y.event_date as second_date -- earliest date + 1 day
    from
    (
        select
        a.player_id,
        a.device_id,
        a.event_date,
        row_number() over ( -- we can also use rank, but it doesn't matter since event_date is unique within player_id
            partition by a.player_id
            order by a.event_date
        ) as rank_sort
        from Activity a
    ) x
    join Activity y on ((y.player_id = x.player_id) and (y.event_date = x.event_date + 1)) -- join the second date only if it it + 1 day to the first date
    where x.rank_sort = 1
) q1,
( -- q2 - select total players
    select
    count(distinct a.player_id) as total_players
    from Activity a
) q2