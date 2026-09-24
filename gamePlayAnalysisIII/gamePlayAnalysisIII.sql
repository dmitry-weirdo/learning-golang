-- 534. Game Play Analysis III

select
a.player_id,
a.event_date,
sum(a.games_played) over ( -- will sum up to the current row inclusive
    partition by a.player_id
    order by a.event_date
) as games_played_so_far
from Activity a