-- 1767. Find the Subtasks That Did Not Execute

with recursive temp as (
    -- base query
    select task_id,subtasks_count as subtask_id from Tasks

    union all

    -- generated recursive query
    -- the magic is that the next recursion step only adds to the rows generated at the previous step!
    select task_id, (subtask_id -1) as subtask_id from temp
    where subtask_id > 1
)
select
*
from temp t
--order by t.task_id, t.subtask_id
except -- excludes the rows from query2 from query1
select
*
from Executed e
