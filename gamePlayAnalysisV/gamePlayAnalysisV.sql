-- 1097. Game Play Analysis V

with q1 as (
    select
    x.player_id,
    x.event_date as first_date, -- earliest date
    y.event_date as second_date -- earliest date + 1 day
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
    left join Activity y on ((y.player_id = x.player_id) and (y.event_date = x.event_date + 1)) -- join the second date only if it it + 1 day to the first date
    where x.rank_sort = 1
)
select
    q1.first_date as "install_dt",
    count(q1.player_id) as "installs",
    round(
        count(q1.second_date)::numeric / count(q1.player_id), -- divide players with 2nd day login by total players (all within one date)
        2
    ) as "Day1_retention"
from q1
group by q1.first_date
order by q1.first_date; -- ordering does not matter