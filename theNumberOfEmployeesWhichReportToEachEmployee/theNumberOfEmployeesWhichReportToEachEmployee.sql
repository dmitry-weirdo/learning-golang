-- 1731. The Number of Employees Which Report to Each Employee

select
x.employee_id,
e.name,
x.reports_count,
x.average_age
from Employees e
join
(
    select
    e.reports_to as "employee_id",
    count(e.employee_id) as "reports_count",
    round(avg(e.age)) as "average_age"
    from Employees e
    where (e.reports_to is not null)
    group by e.reports_to
) x
on (x.employee_id = e.employee_id)
order by x.employee_id