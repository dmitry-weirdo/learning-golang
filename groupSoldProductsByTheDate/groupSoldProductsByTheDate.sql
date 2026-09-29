-- 1484. Group Sold Products By The Date

select
a.sell_date,
count(distinct a.product) as "num_sold", -- same product can be sold multiple times at the same date -> use distinct
string_agg(distinct a.product, ',' order by a.product) as "products" -- same product can be sold multiple times at the same date -> use distinct
from Activities a
group by a.sell_date
order by a.sell_date
