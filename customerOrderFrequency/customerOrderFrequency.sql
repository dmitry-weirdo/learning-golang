-- 1511. Customer Order Frequency

with x as (
    select
    o.customer_id,
    c.name,
    sum(o.quantity * p.price) as "total_price",
    date_trunc('month', o.order_date) as "month" -- contains both year and month
    -- date_trunc('month', o.order_date) = '2020-06-01 00:00:00' as "june_2020", -- test the month comparison
    -- date_trunc('month', o.order_date) = '2020-07-01 00:00:00' as "july_2020"
    from Orders o
    join Product p on (p.product_id = o.product_id)
    join Customers c on (c.customer_id = o.customer_id)
    where (o.order_date >= '2020-06-01') and (o.order_date < '2020-08-01') -- only select the required months
    group by o.customer_id, c.name, month
    having sum(o.quantity * p.price) >= 100 -- only select months where customer spent >= 100
    order by o.customer_id, month
)
-- we only select customers that have target rows for 2 months
select
x.customer_id,
x.name
--count(x.customer_id)
from x
group by x.customer_id, x.name
having count(x.customer_id) = 2
