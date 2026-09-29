-- 1321. Restaurant Growth

select
*
from
(
    select
    x.visited_on,
    sum(x.daily_amount) over (
        order by x.visited_on
        range between interval '6 days' preceding and current row -- 6 previous date and current date
    )::numeric as "amount",
    round(
        sum(x.daily_amount) over (
            order by x.visited_on
            range between interval '6 days' preceding and current row -- 6 previous date and current date
        )::numeric / 7, -- average for 7 days window
        2
    ) as "average_amount"

    from ( -- first, just group sums by date
        select
        visited_on,
        sum(c.amount) as daily_amount
        from Customer c
        group by c.visited_on
    ) x

) y
where y.visited_on >= ( -- only select the days starting from the 7-th day from the earliest date
    select
    min(c.visited_on) + interval '6 days'
    from Customer c
)