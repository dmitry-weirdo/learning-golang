-- Q3. Monthly Transactions I
-- Write your PostgreSQL query statement below

-- nicer solution if we know that sum can take arguments
select
to_char(t.trans_date, 'YYYY-MM') as month,
t.country as country,
count(t.id) as trans_count,
sum( -- sum can take filters!
    (t.state = 'approved')::int -- We need to cast boolean to int, i.e. true will sum up as 1
) as approved_count,
sum(t.amount) as trans_total_amount,
sum(
    case
        when t.state = 'approved'
        then t.amount
        else 0
    end
) as approved_total_amount
from Transactions t
group by
t.country,
to_char(t.trans_date, 'YYYY-MM')

-- same stuff but for approved_count, use count() instead of sum()
select
to_char(t.trans_date, 'YYYY-MM') as month,
t.country as country,
count(t.id) as trans_count,
count(
    case
        when t.state = 'approved'
        then 1
        else null -- count() counts the values that are not null
    end
) as approved_count,
sum(t.amount) as trans_total_amount,
sum(
    case
        when t.state = 'approved'
        then t.amount
        else 0
    end
) as approved_total_amount
from Transactions t
group by
t.country,
to_char(t.trans_date, 'YYYY-MM')

-- ugly and slow solution, but it's working
select
x.month,
x.country,
(
    select count(*) from Transactions t
    where
    ((t.country = x.country) or (t.country is null and x.country is null)) -- country can be null, we still need to join
    and (to_char(t.trans_date, 'YYYY-MM') = x.month)
) as trans_count,
(
    select count(*) from Transactions t
    where
    ((t.country = x.country) or (t.country is null and x.country is null)) -- country can be null, we still need to join
    and (to_char(t.trans_date, 'YYYY-MM') = x.month)
    and (t.state = 'approved')
) as approved_count,
(
    select coalesce(sum(t.amount), 0) from Transactions t
    where
    ((t.country = x.country) or (t.country is null and x.country is null)) -- country can be null, we still need to join
    and (to_char(t.trans_date, 'YYYY-MM') = x.month)
) as trans_total_amount,
(
    select coalesce(sum(t.amount), 0) from Transactions t
    where
    ((t.country = x.country) or (t.country is null and x.country is null)) -- country can be null, we still need to join
    and (to_char(t.trans_date, 'YYYY-MM') = x.month)
    and (t.state = 'approved')
) as approved_total_amount
from (
    select
    distinct
    to_char(t.trans_date, 'YYYY-MM') as month, -- cast '2018-12-18` date to `2018-12`
    t.country
    from Transactions t
    --order by t.country, month -- ordering is not necessary
) x


