create function getHthHighestSalary(N int) returns int
begin
    set N = N - 1;
    return (
        select distinct Salary from Employee group by Salary
        order by Salary desc limit 1 offset N
    );
end

create function getHthHighestSalary(N int) returns int
begin 
    return (
        select max(Salary) from Employee e1
        where N - 1 =
        (select count(distinct e2.Salary)) from Employee e2
        where e2.Salary > e1.Salary
    );
end

create function getHthHighestSalary(N int) returns int
begin
    return (
        select max(Salary) from Employee e1
        where N = 
        (select count(distinct e2.Salary)) from Employee e2
        where e2.Salary >= e2.Salary
    );
end