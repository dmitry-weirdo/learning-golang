-- 1045. Customers Who Bought All Products

select
c.customer_id--,
-- count(distinct c.product_key)
from Customer c
group by c.customer_id
having count(distinct c.product_key) = (select count(product_key) from Product)