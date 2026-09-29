-- 185. Department Top Three Salaries

select
x.department as "Department",
x.employee as "Employee",
x.salary as "Salary"
from (
    select
    d.name as department,
    e.name as employee,
    e.salary as salary,
    dense_rank() over (
        partition by e.departmentId
        order by e.salary desc
    ) as rank
    from Employee e
    join Department d on (d.id = e.departmentId)
) x
where (x.rank <= 3)