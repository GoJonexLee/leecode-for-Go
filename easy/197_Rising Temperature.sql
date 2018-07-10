select w1.Id from Weather w1, Weather w2 where w1.Temperature > w2.Temperature and datediff(w1.RecordDate, w2.RecordDate) = 1;
select w1.Id from Weather w1, Weather w2 where w1.Temperature > w2.Temperature and to_days(w1.RecordDate) = to_days(w2.RecordDate) + 1;
select w1.Id from Weather w1, Weather w2 where w1.Temperature > w2.Temperature and subdate(w1.RecordDate, 1) = w2.RecordDate;