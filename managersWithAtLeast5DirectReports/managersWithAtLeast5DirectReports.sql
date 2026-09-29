-- 570. Managers with at Least 5 Direct Reports

select
e.name
from Employee e
where e.id in
( -- select ids of the managers who have >= 5 employee records
    select
    e.managerId
    from Employee e
    group by e.managerId
    having count(e.id) >= 5
)