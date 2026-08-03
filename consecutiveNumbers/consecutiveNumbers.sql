-- 180. Consecutive Numbers

-- compare with next and next.next
select
distinct x.num as "ConsecutiveNums"
from (
    select
    --l.id,
    l.num,
    lead(num) over (order by id) as next_num_1,
    lead(num, 2) over (order by id) as next_num_2
    from Logs l
) x
where (x.num = x.next_num_1) and (x.num = x.next_num_2) -- value equals value.next and value.next.next

-- compare with next and prev
select
distinct x.num as "ConsecutiveNums"
from (
    select
    --l.id,
    l.num,
    lag(num) over (order by id) as prev_num_1,
    lead(num) over (order by id) as next_num_1
    from Logs l
) x
where (x.num = x.prev_num_1) and (x.num = x.next_num_1) -- value equals value.prev and value.next