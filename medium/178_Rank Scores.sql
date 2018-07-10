select Score, (select count(distinct Score) from Scores where Score >= s.Score) Rank from Scores s order by Score desc;
select Score, (select count(*) from (select distinct Score s from Scores) t where s >= Score) Rank from Scores order by Score desc;
select s.Score, count(distinct t.Score) Rank from Scores s join Scores t on s.Score <= t.Score group by s.Id order by s.Score desc;