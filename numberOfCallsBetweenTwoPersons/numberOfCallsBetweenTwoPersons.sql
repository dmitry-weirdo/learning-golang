-- 1699. Number of Calls Between Two Persons

select
least(c.from_id, c.to_id) as "person1",
greatest(c.from_id, c.to_id) as "person2",
count(c.*) as "call_count",
sum(c.duration) as "total_duration"
from Calls c
group by least(c.from_id, c.to_id), greatest(c.from_id, c.to_id)
-- ordering does not matter