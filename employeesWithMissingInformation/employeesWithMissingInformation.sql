-- 1965. Employees With Missing Information
-- todo: can be also solved using FULL OUTER JOIN

(
    select
    e.employee_id
    from Employees e
    left join Salaries s on (s.employee_id = e.employee_id)
    where (s.salary is null)
)
union
(
    select
    s.employee_id
    from Salaries s
    left join Employees e on (e.employee_id = s.employee_id)
    where (e.name is null)
)
order by employee_id -- yes, we can order union by the common column name. Notably, `order by e.employee_id` will not work