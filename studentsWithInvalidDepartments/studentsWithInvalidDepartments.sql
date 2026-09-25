-- 1350. Students With Invalid Departments

select
s.id,
s.name
from Students s
left join Departments d on (d.id = s.department_id)
where (d.id is null)
