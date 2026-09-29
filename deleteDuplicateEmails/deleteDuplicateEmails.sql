-- 196. Delete Duplicate Emails

delete
from Person p
where exists (
    select
    *
    from Person p2
    where (p2.email = p.email) and (p2.id < p.id)
)