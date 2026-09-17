-- 1440. Evaluate Boolean Expression

select
e.left_operand,
e.operator,
e.right_operand,
case
    when (e.operator = '>') then (lv.value > rv.value)
    when (e.operator = '<') then (lv.value < rv.value)
    when (e.operator = '=') then (lv.value = rv.value)
end::text as "value" -- !!! we need the result as string, NOT as boolean
--lv.*,
--rv.*
from Expressions e
join Variables lv on (lv.name = e.left_operand)
join Variables rv on (rv.name = e.right_operand)