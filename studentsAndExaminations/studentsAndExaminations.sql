-- 1280. Students and Examinations

-- we need cross join (no `on`) to get all combinations of Students x Subjects
select
s.student_id,
s.student_name,
su.subject_name,
count(e.student_id) as attended_exams
from Subjects su
cross join Students s -- cross join does not need any `on` clause
left join Examinations e on (e.student_id = s.student_id and e.subject_name = su.subject_name)
group by s.student_id, s.student_name, su.subject_name
order by s.student_id, su.subject_name -- no need to sort by s.student_name