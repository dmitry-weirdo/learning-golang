-- 1251. Average Selling Price
-- Very similar to "1934. Confirmation Rate"

select
distinct
p.product_id,
coalesce(x.average_price, 0) as "average_price"
from Prices p
left join
(
    select
    us.product_id,

    round(
        sum(us.units * p.price)::numeric / sum(us.units),
        2
    ) as "average_price"

    from UnitsSold us
    join Prices p on (
            (p.product_id = us.product_id)
        and (p.start_date <= us.purchase_date)
        and (p.end_date >= us.purchase_date)
    )
    group by us.product_id
) x
on (x.product_id = p.product_id)
-- ordering is not important