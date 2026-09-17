-- 1398. Customers Who Bought Products A and B but Not C

select
c.customer_id,
c.customer_name
from Customers c
where exists (
    select o.order_id from Orders o where (o.customer_id = c.customer_id) and (o.product_name = 'A')
) and exists (
    select o.order_id from Orders o where (o.customer_id = c.customer_id) and (o.product_name = 'B')
) and not exists (
    select o.order_id from Orders o where (o.customer_id = c.customer_id) and (o.product_name = 'C')
)
