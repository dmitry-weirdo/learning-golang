-- 1270. All People Report to the Given Manager

select
e3.employee_id
from Employees e1
left join Employees e2 on (e2.manager_id = e1.employee_id)
left join Employees e3 on (e3.manager_id = e2.employee_id)
where ((e1.manager_id = 1) or (e2.manager_id = 1) or (e3.manager_id = 1)) -- any of the 3 levels is the manager
and (e3.employee_id <> 1) -- exclude the top manager itself, he also has manager = 1