-- 1225. Report Contiguous Dates

with x as ( -- union the intervals into one table, filter out only the required dates
    (
        select
        fail_date as "date",
        'failed' as "status"
        from Failed f
        where (f.fail_date >= '2019-01-01') and (f.fail_date < '2020-01-01')
    )
    union
    (
        select
        s.success_date as "date",
        'succeeded' as "status"
        from Succeeded s
        where (s.success_date >= '2019-01-01') and (s.success_date < '2020-01-01')
    )
    order by "date"
),
y as ( -- mark group start as 1, other rows as 0
    select
    *,
    -- sum( -- cannot use nested window functions :(
        case
            when x.status is distinct from lag(x.status) over (order by x.date)
            then 1 -- 1 is a group start
            else 0 -- 0 is NOT a group start
        end as "group_start"
    --) over (order by x.date)
    from x
    order by x.date
),
z as ( -- mark groups as increasing values for group identifiers
    select
    y.*,
    sum(y.group_start) over (order by y.date) as "group_id"
    from y
)
-- select the result: group by group_id, get min and max dates for every group
select
z.status as "period_state",
min(z.date) as "start_date",
max(z.date) as "end_date"
from z
group by z.group_id, z.status
