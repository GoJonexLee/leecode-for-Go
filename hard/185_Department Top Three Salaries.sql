select d.Name Department, e.Name Employee, e.Salary from Employee e join Department d on e.DepartmentId = d.Id
where (select count(distinct Salary) from Employee where Salary > e.Salary and DepartmentId = d.Id) < 3
order by d.Name, e.Salary desc;

select d.Name Department, e.Name Employee, e.Salary from
(select e1.Name, e1.Salary, e1.DepartmentId from Employee e1 join Employee e2 on e1.DepartmentId = e2.DepartmentId and d1.Salary <= d2.Salary group by e1.Id
having count(distinct e2.Salary) <= 3) e
join Department d on e.DepartmentId = d.Id
order by d.Name, e.Salary desc;

select d.Name Department, e.Name Employee, e.Salary from Employee e, Department d
where (select count(distinct Salary) from Employee where Salary > e.Salary and DepartmentId = d.Id) in (0, 1, 2)
and e.DepartmentId = d.Id
order by d.Name, e.Salary desc;