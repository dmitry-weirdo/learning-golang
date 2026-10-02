-- 1709. Biggest Window Between Visits

with x as (
    select
    uv.*,

    /* -- debug purposes - show the next date
    coalesce(
        lead(uv.visit_date) over (
            partition by uv.user_id
            order by uv.visit_date
        ),
        '2021-1-1' -- for the last date, return today as next date
    ) as "next_date",
    */

    coalesce(
        lead(uv.visit_date) over (
            partition by uv.user_id
            order by uv.visit_date
        ),
        '2021-1-1' -- for the last date, return today as next date
    ) - uv.visit_date
    as "next_date_diff"

    from UserVisits uv
    order by uv.user_id, uv.visit_date
)
select
x.user_id,
max(x.next_date_diff) as "biggest_window" -- for every user, select only the max date diff
from x
group by x.user_id
order by x.user_id