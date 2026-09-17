-- 1607. Sellers With No Sales

select
s.seller_name
from Seller s
left join Orders o on (
    (o.seller_id = s.seller_id)
    and (extract(year FROM o.sale_date) = 2020)
)
where (o.order_id is null) -- no orders
order by s.seller_name

-- since extract can prevent index on Orders.sale_date from working, let's use the explicit comparisons
-- Not much faster
select
s.seller_name
from Seller s
left join Orders o on (
    (o.seller_id = s.seller_id)
    and (o.sale_date >= DATE '2020-01-01')
    and (o.sale_date <  DATE '2021-01-01')
    -- and (extract(year FROM o.sale_date) = 2020) -- extract can prevent the usage of the index o.sale_date
)
where o.order_id is null
order by s.seller_name