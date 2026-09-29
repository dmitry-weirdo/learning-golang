-- 1501. Countries You Can Safely Invest In

-- Simpler query -> we can join 2 rows by Calls.caller_id and Calls.caller_id using OR in the join clause.
-- Generally, this runs faster than the `union` version
select
c1.name as "country" --,
--c1.country_code,
--avg(c.duration) as "average_duration"
from Calls c
join Person p1 on ((p1.id = c.caller_id) or (p1.id = c.callee_id)) -- join both caller and callee as separate rows!
join Country c1 on (c1.country_code = substr(p1.phone_number, 1, 3)) -- substr uses 1-based indexes
group by c1.name, c1.country_code
having avg(c.duration) > (
    -- global average call duration
    select
    avg(c.duration)
    from Calls c
)

-- Solution using union all with separate joins on Calls.caller_id and Calls.callee_id
select
x.name as "country"--,
-- x.country_code,
-- sum(total_duration) as total_duration,
-- sum(total_calls) as total_calls,
-- sum(total_duration) / sum(total_calls) as average_duration
from
(
    (
        select
        c1.name,
        c1.country_code,
        sum(c.duration) as "total_duration",
        count(c.*) as "total_calls"
        from Calls c
        join Person p1 on (p1.id = c.caller_id)
        join Country c1 on (c1.country_code = substr(p1.phone_number, 1, 3)) -- substr uses 1-based indexes
        group by c1.name, c1.country_code
    )
    union all
    (
        select
        c1.name,
        c1.country_code,
        sum(c.duration) as "total_duration",
        count(c.*) as "total_calls"
        from Calls c
        join Person p1 on (p1.id = c.callee_id)
        join Country c1 on (c1.country_code = substr(p1.phone_number, 1, 3)) -- substr uses 1-based indexes
        group by c1.name, c1.country_code
    )
) x
group by x.name, x.country_code
having sum(total_duration) / sum(total_calls) > (
    -- global average call duration
    select
    avg(c.duration)
    from Calls c
)
