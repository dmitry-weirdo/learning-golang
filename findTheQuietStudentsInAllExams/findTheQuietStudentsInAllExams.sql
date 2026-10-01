-- 1412. Find the Quiet Students in All Exams

select
s.*
from Student s
left join
(
    select
    distinct
    e.student_id
    from Exam e
    left join (
        -- for every exam, select min and max scores
        select
        e.exam_id,
        min(e.score) as exam_min_score,
        max(e.score) as exam_max_score
        from Exam e
        group by e.exam_id
    ) x
    on (x.exam_id = e.exam_id)
    where (e.score = x.exam_min_score) or (e.score = x.exam_max_score) -- get student scores with min and max
    --where (e.score != x.exam_min_score) and (e.score != x.exam_max_score) -- exclude student scores with min and max
) y
on (y.student_id = s.student_id)
where (y.student_id is null) -- for the student, there is no exam where he/she scored the min or max score
and exists ( -- only select students with at least one exam
    select
    e.student_id
    from Exam e
    where (e.student_id = s.student_id)
)
order by s.student_id