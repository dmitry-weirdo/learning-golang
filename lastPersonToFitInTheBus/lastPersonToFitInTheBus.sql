-- 1204. Last Person to Fit in the Bus

select
x.person_name
from (
    select
    q.person_name,
    sum(q.weight) over (order by q.turn) as "total_weight" -- running sum of q.weight
    from Queue q
    order by q.turn
) x
where (x.total_weight <= 1000) -- we only need the persons whose total_weight fit into the 1000 limit
order by x.total_weight desc -- just select 1 person with max total_weight
limit 1