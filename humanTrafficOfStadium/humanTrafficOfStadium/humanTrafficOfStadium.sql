-- 601. Human Traffic of Stadium
select
id,
visit_date,
people
from (
    select
    lag(id) over (order by id) as prev_id_1,
    lag(id, 2) over (order by id) as prev_id_2,
    lead(id) over (order by id) as next_id_1,
    lead(id, 2) over (order by id) as next_id_2,
    *
    from Stadium s
    where s.people >= 100
) x
where (
    ((x.next_id_1 = x.id + 1) and (x.next_id_2 = x.id + 2)) -- first -> 2 next are consecutive
    or
    ((x.prev_id_1 = x.id - 1) and (x.prev_id_2 = x.id - 2)) -- previous -> 2 previous are consecutive
    or
    ((x.prev_id_1 = x.id - 1) and (x.next_id_1 = x.id + 1)) -- middle -> previous and next are consecutive
)
order by visit_date