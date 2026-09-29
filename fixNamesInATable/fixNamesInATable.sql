-- 1667. Fix Names in a Table

select
u.user_id,
-- initcap(u.name) as "name" -- no, we want "Marry ann", NOT "Marry Ann", i.e. capitalize only the very first letter, not in every word
upper(left(u.name, 1)) || lower(right(u.name, -1)) as "name" -- right(-1) removes characters from left
from Users u
order by u.user_id
