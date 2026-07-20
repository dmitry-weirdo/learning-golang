-- 262. Trips and Users
-- Written for PostgreSQL

/**
-- Write your PostgreSQL query statement below
select
-- u.banned as user_banned,
-- d.banned as driver_banned,
t.request_at,
u.banned as user_banned,
d.banned as driver_banned,
(u.banned = 'Yes' or d.banned = 'Yes') as someone_banned,
t.status,
(t.status = 'cancelled_by_driver' or t.status = 'cancelled_by_client') as canceled,
count(t.id)
from Trips t
join users u on ((u.users_id = t.client_id) and (u.banned = 'No')) -- if we can exclude jobs and days with someone banned
join users d on ((d.users_id = t.driver_id) and (d.banned = 'No')) -- if we can exclude jobs and days with someone banned
where (t.request_at >= '2013-10-01') and (t.request_at <= '2013-10-03')
group by t.request_at, u.banned, d.banned, t.status
**/


select
y.request_at as "Day",
--y.cancelled_jobs,
--y.total_jobs,
-- to_char(y.cancelled_jobs::numeric / y.total_jobs, 'FM0.00') as "Cancellation Rate",
round(y.cancelled_jobs::numeric / y.total_jobs, 2) as "Cancellation Rate" -- We need `round`, NOT `to_char` that will return a quoted string
-- y.cancelled_jobs / y.total_jobs as "Cancellation Rate"
from
(
    select
    x.request_at,
    (
        select
        count(t.id)
        from Trips t
        join users u on ((u.users_id = t.client_id) and (u.banned = 'No')) -- if we can exclude jobs and days with someone banned
        join users d on ((d.users_id = t.driver_id) and (d.banned = 'No')) -- if we can exclude jobs and days with someone banned
        where (t.request_at = x.request_at)
        and (t.status = 'cancelled_by_driver' or t.status = 'cancelled_by_client')
    ) as cancelled_jobs,
    (
        select
        count(t.id)
        from Trips t
        join users u on ((u.users_id = t.client_id) and (u.banned = 'No')) -- if we can exclude jobs and days with someone banned
        join users d on ((d.users_id = t.driver_id) and (d.banned = 'No')) -- if we can exclude jobs and days with someone banned
        where (t.request_at = x.request_at)
        -- and (t.status = 'cancelled_by_driver' or t.status = 'cancelled_by_client')
    ) as total_jobs
    from (
        select
        distinct
        t.request_at
        from Trips t
        where (t.request_at >= '2013-10-01') and (t.request_at <= '2013-10-03')
        order by t.request_at
    ) x
) y
where (y.total_jobs <> 0) -- avoid division by 0
