-- 1070. Product Sales Analysis III

select
s.product_id,
s.year as first_year,
s.quantity,
s.price
from Sales s
where (s.product_id, s.year) in ( -- yes, psql supports multiple fields for `in`
    select
    s.product_id,
    min(year)
    from Sales s
    group by s.product_id
    order by s.product_id
)