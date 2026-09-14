-- 197. Rising Temperature
select
x.id as "Id"
from (
    select
    w.id,
    w.temperature,
    w.recordDate,
    lag(temperature) over (order by w.recordDate) as prev_temperature,
    lag(recordDate) over (order by w.recordDate) as prev_date
    from Weather w
    order by w.id
) x
where (x.temperature > x.prev_temperature)
and (x.prev_date = x.recordDate - 1); -- psql magic of subtracting 1 day from a date field