-- 1407. Top Travellers

-- as right join (we need all records from Users, even if there are no records in Rides)
select
u.name as "name",
coalesce(sum(r.distance), 0) as "travelled_distance" -- no records in Rides must be 0, not null
from Rides r
right join Users u on (u.id = r.user_id)
group by r.user_id, u.name
order by travelled_distance desc, u.name asc


-- as left join (we need all records from Users, even if there are no records in Rides)
select
u.name as "name",
coalesce(sum(r.distance), 0) as "travelled_distance" -- no records in Rides must be 0, not null
from Users u
left join Rides r on (u.id = r.user_id)
group by r.user_id, u.name
order by travelled_distance desc, u.name asc


