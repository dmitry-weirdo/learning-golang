-- 1978. Employees Whose Manager Left the Company

select
e.employee_id
from Employees e
left join Employees m on (m.employee_id = e.manager_id)
where (e.salary < 30000) and (e.manager_id is not null) and (m.employee_id is null)
order by e.employee_id;