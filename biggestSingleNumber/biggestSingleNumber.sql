-- 619. Biggest Single Number

select (
    select
    mn.num
    from MyNumbers mn
    group by mn.num
    having count(mn.num) = 1 -- select only the single numbers
    order by mn.num desc
    limit 1 -- select only the 1st biggest number
) as "num" -- to return `null` when there are no single numbers