-- 1532. The Most Recent Three Orders

select
x.customer_name,
x.customer_id,
x.order_id,
x.order_date
from (
    select
    c.name as customer_name,
    c.customer_id,
    o.order_id,
    o.order_date,
    row_number() over (
        partition by o.customer_id
        order by o.order_date desc
    ) as rank_sort -- we can also use rank() instead of row_number()
    from Orders o
    left join Customers c on (c.customer_id = o.customer_id)
    order by c.name asc, c.customer_id asc, o.order_date desc
) x
where (x.rank_sort <= 3) -- we need just the queries with row_number() <= 3
-- ordering will be inherited from x