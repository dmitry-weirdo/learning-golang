-- 1285. Find the Start and End Number of Continuous Ranges
-- Much simpler variation of "1225. Report Contiguous Dates"

with y as ( -- mark group start as 1, other rows as 0
    select
    *,
    case
        when x.log_id is distinct from lag(x.log_id) over (order by x.log_id) + 1 -- we compare the difference from "previous + 1"
        then 1 -- 1 is a group start
        else 0 -- 0 is NOT a group start
    end as "group_start"
    from Logs x
    order by x.log_id
),
z as ( -- mark groups as increasing values for group identifiers
    select
    y.*,
    sum(y.group_start) over (order by y.log_id) as "group_id"
    from y
)
-- select the result: group by group_id, get min and max values for every group
select
min(z.log_id) as "start_id",
max(z.log_id) as "end_id"
from z
group by z.group_id
order by z.group_id