-- 1571. Warehouse Manager
select
w.name as warehouse_name,
sum(w.units * p.width * p.length * p.height) as volume
from Warehouse w
left join Products p on (p.product_id = w.product_id)
group by w.name
-- order is not important