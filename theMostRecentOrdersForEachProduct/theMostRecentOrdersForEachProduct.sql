-- 1549. The Most Recent Orders for Each Product

select
x.product_name,
x.product_id,
x.order_id,
x.order_date
from
(
    select
    p.product_name,
    p.product_id,
    o.order_id,
    o.order_date,
    rank() over (
        partition by p.product_id
        order by o.order_date desc
    ) as rn
    from Products p
    join Orders o on (o.product_id = p.product_id) -- we don't select products without orders, therefore join (NOT left join)
) x
where (x.rn = 1) -- only select the most recent date for every product
order by x.product_name asc, x.product_id asc, x.order_id asc