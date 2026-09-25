-- 2686. Immediate Food Delivery III
select
d.order_date,
round(
    100::numeric -- multiply by 100 to get the percentage
    * count(
        case
            when d.order_date = d.customer_pref_delivery_date
            then 1
            else null -- to not add a row into count(), we must use null, and NOT 0
        end
    )
    / count(*),
    2 -- round to 2 decimal digits
) as "immediate_percentage"
from Delivery d
group by d.order_date
order by d.order_date
