-- 1495. Friendly Movies Streamed Last Month

select
c.title
from Content c
where (c.kids_content = 'Y')
and (c.content_type = 'Movies')
and exists(
    select
    tvp.content_id
    from TVProgram tvp
    where (tvp.content_id = c.content_id)
    and (tvp.program_date >= '2020-06-01') -- All days in June 2026
    and (tvp.program_date <= '2020-06-30')
)