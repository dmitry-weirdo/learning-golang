-- 1211. Queries Quality and Percentage

select
q.query_name,

round(
    avg(rating::numeric / position),
    2
) as "quality",

round(
    100::numeric
    * count(case when q.rating < 3 then 1 else null end)
    / count(*),
    2
) as "poor_query_percentage"

from Queries q
group by q.query_name