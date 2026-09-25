-- 603. Consecutive Available Seats

select
distinct
c.seat_id
from Cinema c
join Cinema c2 on (
    (c2.free = 1)
    and (abs(c.seat_id - c2.seat_id) = 1) -- previous or next set is also free
)
where (c.free = 1)
order by c.seat_id
