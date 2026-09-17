-- 607. Sales Person
select
sp.name
from SalesPerson sp
where not exists (
    select
    order_id
    from Orders o
    join Company c on (c.com_id = o.com_id and c.name = 'RED')
    where (o.sales_id = sp.sales_id)
)
