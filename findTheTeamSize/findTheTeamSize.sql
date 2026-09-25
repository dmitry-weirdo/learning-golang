-- 1303. Find the Team Size
select
e.employee_id,
x.team_size
from Employee e
join (
    select
    e.team_id,
    count(e.employee_id) as team_size
    from Employee e
    group by e.team_id
) x
on (x.team_id = e.team_id)