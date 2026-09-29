-- 1164. Product Price at a Given Date

select
distinct
p.product_id,
coalesce(x.new_price, 10) as "price" -- no price in the date range in the table -> use the default price 10
from Products p
left join (
    select
    p.product_id,
    p.new_price,
    row_number() over ( -- we can also use rank, but it doesn't matter since event_date is unique within player_id
        partition by p.product_id
        order by p.change_date desc
    ) as rn
    from Products p
    where (p.change_date <= '2019-08-16') -- we're only interested in changes that happened not later than the target date
) x
on (
        (x.product_id = p.product_id)
    and (x.rn = 1) -- only select first price for every product_id
)
