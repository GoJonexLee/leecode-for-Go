select distinct t1.* from stadium t1, stadium t2, stadium t3
where t1.people >= 100 and t2.people >= 100 and t3.people >= 100 and (
    (t1.id - t2.id = 1 and t1.id - t3.id = 2 and t2.id - t3.id = 1) or
    (t2.id - t1.id = 1 and t2.id - t3.id = 2 and t1.id - t3.id = 1) or
    (t3.id - t2.id = 1 and t2.id - t1.id = 1 and t3.id - t1.id = 2)
) order by t1.id

SELECT distinct s1.* FROM stadium AS s1, stadium AS s2, stadium as s3
WHERE s1.people>=100 AND s2.people>=100AND s3.people>=100 AND (
    (s1.id + 1 = s2.id AND s1.id + 2 = s3.id) OR 
    (s1.id - 1 = s2.id AND s1.id + 1 = s3.id) OR
    (s1.id - 2 = s2.id AND s1.id - 1 = s3.id)
) order by s1.date