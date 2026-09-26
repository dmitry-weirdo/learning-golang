-- 1076. Project Employees II

with x as ( -- yes, we can use aliases for CTE!
    select
    p.project_id,
    count(distinct p.employee_id) as employee_count
    from Project p
    group by p.project_id
)
select
x.project_id
from x
where x.employee_count = (
    select max(x.employee_count) from x
)
