-- 1050. Actors and Directors Who Cooperated At Least Three Times

select
ad.actor_id,
ad.director_id
from ActorDirector ad
group by ad.actor_id, ad.director_id
having count(*) >= 3
