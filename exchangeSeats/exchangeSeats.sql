-- 626. Exchange Seats

select
s.id,
--s.student,

case
    when s.id % 2 = 0 then -- even -> take prev value
        coalesce(
            lag(s.student) over (order by s.id),
            s.student -- leave the current row for the first row
        )
    else -- odd -> take next value
        coalesce(
            lead(s.student) over (order by s.id),
            s.student -- leave the current row for the last row
        )

end as "student"--,

/*
coalesce(
    lag(s.student) over (order by s.id),
    s.student -- leave the current row for the first row
) as prev,
coalesce(
    lead(s.student) over (order by s.id),
    s.student -- leave the current row for the last row
) as next
*/
from Seat s


