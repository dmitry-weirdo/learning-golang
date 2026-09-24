-- 512. Game Play Analysis II

select
x.player_id,
x.device_id
from
(
    select
    a.player_id,
    a.device_id,
    row_number() over ( -- we can also use rank, but it doesn't matter since event_date is unique within player_id
        partition by a.player_id
        order by a.event_date
    ) as rank_sort
    from Activity a
) x
where x.rank_sort = 1