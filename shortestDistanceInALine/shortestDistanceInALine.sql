-- 613. Shortest Distance in a Line
-- Same as "1709. Biggest Window Between Visits"

with x as (
    select
    p.*,

    lead(p.x) over (
        order by uv.x
    ) - p.x
    as "next_diff"

    from Point p
    order by p.x
)
select
min(x.next_diff)  as "shortest" -- will ignore next_diff = null for the last row
from x
