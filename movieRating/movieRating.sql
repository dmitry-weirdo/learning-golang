-- 1341. Movie Rating

-- todo: MAYBE selecting having count = (select max(count(...))) will run faster, but not sure.

(
    select
    -- mr.user_id,
    u.name as results --, -- first query column name will be used
    -- count(distinct mr.movie_id) as "movies_count"
    from MovieRating mr
    left join Users u on (u.user_id = mr.user_id)
    group by mr.user_id, u.name
    order by count(distinct mr.movie_id) desc, u.name
    limit 1
)
union all -- preserve duplicate rows from query 1 and query 2, i.e. if the user name and the movie name are the same, 2 rows will be returned, not merged rows. If we want to merge the duplicates, just `union` should be used.
(
    select
    m.title --,
    -- avg(mr.rating)
    from MovieRating mr
    join Movies m on (m.movie_id = mr.movie_id)
    where (mr.created_at >= '2020-02-01') and (mr.created_at <= '2020-02-29')
    group by mr.movie_id, m.title
    order by avg(mr.rating) desc, m.title
    limit 1
)
