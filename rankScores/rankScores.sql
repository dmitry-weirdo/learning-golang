-- 178. Rank Scores

-- dense_rank() will do "no gaps after tie" logic

select
    s.score,
    dense_rank() over (order by s.score desc) as "rank"
from Scores s
-- order by s.score desc -- todo: this shouldn't work without this order by, but still the query is sorted even without this `order by`
-- notably, if we write `order by s.id`, it will be order by id, ranks will be set correctly by `order by s.score desc`.