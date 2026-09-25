-- 1741. Find Total Time Spent by Each Employee
select
    e.event_day as "day",
    e.emp_id as "emp_id",
    sum(e.out_time - e.in_time) as "total_time"
from Employees e
group by e.emp_id, e.event_day
-- ordering does not matter