select Request_at Day, round(count(if(Status != 'completed', TRUE, null)) / count(*), 2) 'Cancellation Rate' from Trips
where (Request_at between '2013-10-01' and '2013-10-03') and Client_Id in (select Users_Id from Users where Banned = 'No')
group by Request_at;