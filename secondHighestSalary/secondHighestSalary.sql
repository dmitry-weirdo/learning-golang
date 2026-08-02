-- 176. Second Highest Salary
-- this will need O(N * log N) on sorting
-- but maybe the index is already there and we just return (log N) on an index binary tree
select
x.salary as "SecondHighestSalary"
from (
    select
    distinct
    e.salary
    from Employee e
    order by e.salary desc
    offset 1
    limit 1
) x -- x has 0 or 1 rows
right join (select 1 as fakeColumn) y on (y.fakeColumn is not null) -- y has 1 row


-- max() returns `null` on no rows
-- this should be O(2*N),
-- but it is running in the same time as the right-join solution (index used for max() and running in O(log N) probably)
-- We're selecting the max salary that is smaller than the max salary. This nicely avoids duplicates.
select
max(e.salary) as "SecondHighestSalary"
from Employee e
where e.salary <
(
    select
    max(e.salary)
    from Employee e
)
