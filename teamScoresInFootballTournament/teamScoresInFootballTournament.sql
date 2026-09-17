-- 1212. Team Scores in Football Tournament
select
x.team_id,
x.team_name,
sum (x.num_points) as "num_points"
from (
    ( -- points of home (host) team
        select
        t.team_id,
        t.team_name,
        case
            when (m.host_goals > m.guest_goals) then 3
            when (m.host_goals = m.guest_goals) then 1
            else 0
        end as "num_points"
        from Matches m
        right join Teams t on (t.team_id = m.host_team) -- join the host team, don't forget that can be teams without host matches
    )
    union all -- do not remove duplicate rows (e.g win in both home and guest match)
    ( -- points of guest team
        select
        t.team_id,
        t.team_name,
        case
            when (m.guest_goals > m.host_goals) then 3
            when (m.guest_goals = m.host_goals) then 1
            else 0
        end as "num_points"
        from Matches m
        right join Teams t on (t.team_id = m.guest_team) -- join the guest team, don't forget that can be teams without guest matches
    )
) x
group by x.team_id, x.team_name
order by num_points desc, x.team_id asc