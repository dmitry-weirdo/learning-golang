-- 1555. Bank Account Summary
select
q1.user_id,
--q2.user_id,
q1.user_name,
--coalesce(q2.amount_plus, 0) as amount_plus,
--coalesce(q1.amount_minus, 0) as amount_minus,
coalesce(q1.credit) + coalesce(q2.amount_plus, 0) - coalesce(q1.amount_minus, 0) as "credit",
case
    when coalesce(q1.credit) + coalesce(q2.amount_plus, 0) - coalesce(q1.amount_minus, 0) < 0
    then 'Yes'
    else 'No'
end as "credit_limit_breached"
from
( -- select transactions where user is paid_from, i.e. when he gets this amount as minus to his account
    select
    u.user_id,
    u.user_name,
    u.credit,
    sum(tf.amount) as "amount_minus"
    from Users u
    left join Transactions tf on (tf.paid_by = u.user_id)
    group by u.user_id, u.user_name, u.credit
) q1,
( -- select transactions where user is paid_to, i.e. when he gets this amount as plus to his account
    select
    u.user_id,
    sum(tt.amount) as "amount_plus"
    from Users u
    left join Transactions tt on (tt.paid_to = u.user_id)
    group by u.user_id
) q2
where (q1.user_id = q2.user_id)
