-- 602. Friend Requests II: Who Has the Most Friends

select
x.id,
count(x.id) as "num"
from (
    (
        select
        ra.requester_id as "id"
        from RequestAccepted ra
    )
    union all
    (
        select
        ra.accepter_id as "id"
        from RequestAccepted ra
    )
) x
group by x.id
order by num desc -- just select one person with the max count. Yes, we can order by alias.
limit 1