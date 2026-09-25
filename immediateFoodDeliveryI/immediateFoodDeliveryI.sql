-- 1173. Immediate Food Delivery I
select
round(
    q1.immediate_count::numeric * 100 / q2.total_count, -- multiply by 100 to get the percentage
    2
) as "immediate_percentage"
from
(
    select
    count(*) as immediate_count
    from Delivery d
    where (d.order_date = d.customer_pref_delivery_date )
) q1,
(
    select
    count(*) as total_count
    from Delivery d
) q2
