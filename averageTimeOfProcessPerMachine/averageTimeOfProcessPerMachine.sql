-- 1661. Average Time of Process per Machine

select
a.machine_id,
round(
    avg(a2.timestamp - a.timestamp)::numeric, -- average of (endTime - startTime) for all the jobs of this machine. Need to case to `numeric` to make `round` work
    3 -- round to 3 decimal places
) as "processing_time"
from Activity a -- 'start' event
join Activity a2 on ( -- join 'start' event onto its corresponding 'end' event
        (a2.process_id = a.process_id) -- same process
    and (a2.machine_id = a.machine_id) -- same machine
    and (a2.activity_type = 'end')
)
where (a.activity_type = 'start')
group by a.machine_id