-- 1789. Primary Department for Each Employee

select
distinct
e.employee_id,
coalesce(e2.department_id, e.department_id) as "department_id" -- if there is no 'Y' department, we just select the employee department (it should be a single department)
from Employee e
left join Employee e2 on (
        (e2.employee_id = e.employee_id)
    and (e2.primary_flag = 'Y')
)
