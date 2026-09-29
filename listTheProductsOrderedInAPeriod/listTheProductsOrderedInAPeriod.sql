-- 1327. List the Products Ordered in a Period

select
p.product_name as "product_name",
sum(o.unit) as "unit"
from Orders o
join Products p on (p.product_id = o.product_id)
where (o.order_date >= '2020-02-01') and (o.order_date <= '2020-02-29')
group by o.product_id, p.product_name
having sum(o.unit) >= 100