-- 1264. Page Recommendations

select
distinct
l.page_id as "recommended_page"
from Likes l
where
( l.user_id in
    (
        (
            select
            f.user2_id
            from Friendship f
            where (f.user1_id = 1)
        )
        union
        (
            select
            f.user1_id
            from Friendship f
            where (f.user2_id = 1)
        )
    )
)
and not exists (
    select
    *
    from Likes l2
    where (l2.user_id = 1)
    and (l2.page_id = l.page_id)
)

/* -- this is slower than `and not exists`
    and l.page_id not in (
    select
    l.page_id
    from Likes l
    where (l.user_id = 1)
)
*/