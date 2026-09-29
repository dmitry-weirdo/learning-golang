-- 1633. Percentage of Users Attended a Contest

select
r.contest_id,
--count(distinct r.user_id),
round(
    100::numeric * count(distinct r.user_id) / x.total_users,
    2
) as "percentage"

from Register r
left join (
    select
    count(user_id) as total_users
    from Users u
) x on (1 = 1) -- fake join, always add the joined column to every row
group by r.contest_id, x.total_users
order by percentage desc, r.contest_id asc
