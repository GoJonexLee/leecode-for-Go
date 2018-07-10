delete from Person where Id not in (select Id from (select min(Id) Id from Person group by Email) p);
delete p2 from Person p1 join Person p2 on p2.Email = p1.Email where p2.Id > p1.Id;
delete p2 from Person p1, Persion p2 where p1.Email = p2.Email and p2.Id > p2.Id;