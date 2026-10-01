-- 1596. The Most Frequently Ordered Products for Each Customer

select
x.customer_id,
x.product_id,
x.product_name
from
(
    select
    o.customer_id,
    o.product_id,
    p.product_name,
    -- count(o.product_id),
    rank() over (
        partition by o.customer_id
        order by count(o.product_id) desc
    ) as rank
    from Orders o
    join Products p on (p.product_id = o.product_id)
    group by o.customer_id, o.product_id, p.product_name
) x
where (x.rank = 1) -- only select top-ranked products for every customer