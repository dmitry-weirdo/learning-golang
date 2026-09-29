-- 1907. Count Salary Categories

(
    select
    'Low Salary' as "category",
    count(a.account_id) as "accounts_count"
    from Accounts a
    where (a.income < 20000)
)
union
(
    select
    'Average Salary' as "category",
    count(a.account_id) as "accounts_count"
    from Accounts a
    where (a.income >= 20000) and (a.income <= 50000)
)
union
(
    select
    'High Salary' as "category",
    count(a.account_id) as "accounts_count"
    from Accounts a
    where (a.income > 50000)
)
