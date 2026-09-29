-- 1934. Confirmation Rate
-- Very similar to "1251. Average Selling Price".

select
distinct
s.user_id,
coalesce(x.confirmation_rate, 0) as "confirmation_rate" -- null rate must be returned as 0
from Signups s
left join
(
    select
    c.user_id,
    -- count(case when c.action = 'confirmed' then 1 else null end), as "confirmed_count"
    -- count(*) as "total_count"
    round(
        count(case when c.action = 'confirmed' then 1 else null end)::numeric / count(*),
        2
    ) as "confirmation_rate"
    from Confirmations c
    group by c.user_id
) x
on (x.user_id = s.user_id)