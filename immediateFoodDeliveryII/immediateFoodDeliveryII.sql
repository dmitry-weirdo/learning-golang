-- 1174. Immediate Food Delivery II
select
round(
    100::numeric -- multiply by 100 to get the percentage
    * count(
        case
            when x.order_date = x.customer_pref_delivery_date
            then 1
            else null -- to not add in count, we must use null, and NOT 0
        end
    )
    / count(*),
    2 -- round to 2 decimal digits
) as "immediate_percentage"
from
(
    select
    d.customer_id,
    d.order_date,
    d.customer_pref_delivery_date,
    row_number() over ( -- we can also use rank, but it doesn't matter since first order_date is unique within customer_id
        partition by d.customer_id
        order by d.order_date
    ) as rank_sort
    from Delivery d
) x
where (x.rank_sort = 1) -- only select first order_date for every customer
