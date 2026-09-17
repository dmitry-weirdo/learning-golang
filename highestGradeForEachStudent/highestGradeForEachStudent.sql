-- 1112. Highest Grade For Each Student

select
e.student_id,
e.course_id,
e.grade
from (
    select
    e.*,
    row_number() over (
        partition by e.student_id
        order by e.grade desc, e.course_id asc
    ) as rn
    from Enrollments e
) e
where (e.rn = 1) -- only one row in case of tie by grade
order by e.student_id;