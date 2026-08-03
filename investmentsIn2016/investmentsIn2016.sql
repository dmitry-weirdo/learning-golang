-- 585. Investments in 2016

-- `exists` / `not exists` are better than `count > 0` / `count = 0` since they will fail-fast on a single found value
select
round(sum(i.tiv_2016)::numeric, 2) as "tiv_2016" -- round to 2 decimal places
-- *
from Insurance i
where
(
    exists (
        select
        i1.pid
        from Insurance i1
        where (i1.pid != i.pid) and (i1.tiv_2015 = i.tiv_2015)
    )
)
and
(
    not exists (
        select
        i2.pid
        from Insurance i2
        where (i2.pid != i.pid) and (i2.lat = i.lat) and (i2.lon = i.lon)
    )
)


-- my solution, may be non-optimal in terms of subqueries
select
round(sum(i.tiv_2016)::numeric, 2) as "tiv_2016" -- round to 2 decimal places
-- *
from Insurance i
where
(
    (
        select
        count(i1.pid)
        from Insurance i1
        where (i1.pid != i.pid) and (i1.tiv_2015 = i.tiv_2015)
    ) > 0
)
and
(
    (
        select
        count(i2.pid)
        from Insurance i2
        where (i2.pid != i.pid) and (i2.lat = i.lat) and (i2.lon = i.lon)
    ) = 0
)