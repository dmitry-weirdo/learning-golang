-- 1587. Bank Account Summary II
select
u.name as "name",
sum(t.amount) as "balance"
from Users u
join Transactions t on (t.account = u.account)
group by u.account, u.name
having sum(t.amount) > 10000