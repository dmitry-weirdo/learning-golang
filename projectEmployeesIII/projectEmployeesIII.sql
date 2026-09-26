-- 1077. Project Employees III

select
x.project_id,
x.employee_id
from (
    select
    p.project_id,
    p.employee_id,
    rank() over (
        partition by p.project_id
        order by e.experience_years desc
    ) as rn
    from Project p
    join Employee e on (e.employee_id = p.employee_id)
) x
where (x.rn = 1) -- only take most experienced employees for every project