-- 184. Department Highest Salary

select
x.department_name as "Department",
x.employee_name as "Employee",
x.salary as "Salary"
from (
    select
    d.name as department_name,
    e.name as employee_name,
    e.salary as salary,
    dense_rank() over (
        partition by d.id
        order by e.salary desc
    ) as employee_rank
    from Department d
    join Employee e on (e.departmentId = d.id) -- NOT `left join` since we don't want any result rows for departments without employees
    order by d.id, e.salary desc
) x
where (x.employee_rank = 1) -- only take employees with highest salaries