-- 1890. The Latest Login in 2020

-- instead of using extract(year), we can just compare the dates
-- PROBABLY it will work better if there is an index on the `Logins.time_stamp` field.
select
l.user_id,
max(l.time_stamp) as "last_stamp"
from Logins l
where (l.time_stamp >= '2020-01-01')
and (l.time_stamp < '2021-01-01')
group by l.user_id