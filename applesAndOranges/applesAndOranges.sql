-- 1445. Apples & Oranges

select
s.sale_date,
coalesce(s.sold_num, 0) - coalesce(s2.sold_num, 0) as "diff"
from Sales s
left join Sales s2 on (
    (s2.sale_date = s.sale_date)
    and (s2.fruit = 'oranges') -- right table is only oranges
)
where (s.fruit = 'apples') -- left table is only apples