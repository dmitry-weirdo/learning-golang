-- 1831. Maximum Transaction Each Day

select
x.transaction_id
from (
    select
    t.transaction_id,
    cast(t.day as date) as transaction_date, -- cut just date from date + time format
    t.amount,
    rank() over (
        partition by cast(t.day as date)
        order by t.amount desc
    ) as rn
    from Transactions t
) x
where (x.rn = 1) -- within every date, only select values with max amount
order by x.transaction_id