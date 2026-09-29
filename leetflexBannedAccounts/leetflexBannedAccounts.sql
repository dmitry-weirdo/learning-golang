-- 1747. Leetflex Banned Accounts

select
distinct
li.account_id
from LogInfo li
left join LogInfo li2 on (
        (li2.account_id = li.account_id) -- same account
    and (li2.ip_address <> li.ip_address) -- different ip
    and (li.login <= li2.logout) -- intersection logic: (start1 <= end2) and (start2 <= end1)
    and (li2.login <= li.logout)
)
where (li2.account_id is not null) -- only select accounts where the intersecting login with same account and different IP exists