-- 608. Tree Node

select
t.id,
case
    when t.p_id is null then 'Root'
    when ch.id is null then 'Leaf' -- using alias of the left join
    else 'Inner'
end as type --,
-- ch.id,
-- ch.p_id
from Tree t
left join lateral (
    select
    id --, -- we only need ch.id from the child
    -- p_id
    from Tree ch
    where ch.p_id = t.id
    limit 1 -- only select maximum one child node
) ch on true;