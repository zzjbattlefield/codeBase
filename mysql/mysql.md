# CURD
## 增删改查技巧
### 忽略错误 "ignore" 
使用ignore可以使Mysql忽略错误语句继续执行  

栗子:  
```mysql
insert ignore into "table" 
```

### DUPLICATE KEY UPDATE 
实现不存在就插入,存在就更新  

栗子:  
```mysql
INSERT INTO "table" (id,emp,ip) VALUES 
(5,8004,"192.168.1.1"),
(6,8005,"192.168.1.2")
ON DUPLICATE KEY UPDATE ip = VALUES(ip)
当要插入的数据在表中已经存在的时候(字段为唯一性约束)就会用插入的ip去更新原来表中的数据
```

### 表连接修改
可以将UPDATE语句中的WHERE子查询改成表连接  

```mysql
UPDATE t_emp e JOIN t_dept d ON e.deptno = d.deptno AND d.name='SALES'
SET e.sal=10000,d.name='销售部';
将t_emp表中所有叫销售部的员工工资改成10000,并且修改t_dept中的部门名称
```

### 表连接删除
DELETE语句也可以使用表连接  

```mysql
DELETE e,d FROM t_emp e join t_dept d on e.deptno = d.deptno AND d.dname='销售部';
DELETE 后跟随的表都是会删除的表上面会把部门表中的销售部和员工表中销售部员工全部删除
```
# 数据库使用技巧 

## 表的主键使用数字ID还是UUID 

### UUID的好处 
- 使用UUID生成的主键值全局唯一
- 跨服务器合并数据更方便

### UUID的缺点
- UUID占用16个字节,比4字节int和8字节bigint更占用存储空间
- UUID是字符串类型,查询速度慢
- UUID不是顺序增长,作为主键，随机IO大

### 主键自动增长的优点
- int和bigint存储空间占用小
- MySql索引数字类型效率比字符串高
- 自动增长的主键值io写入连续性好

#### 总结：目前无论什么情况都不建议使用UUID做主键

## 优化SQL语句 

> - 尽量不要SELECT * 
>> ```mysql
>>    SELECT * FROM t_emp;
>> ```

>- 谨慎使用模糊查询
>> ```mysql
>>    SELECT ename FROM t_emp WHERE ename LIKE '%S%';
>> ```
>>如果ename有索引，like子句左边有%则会跳过索引执行全表扫描
>> ```mysql
>>    SELECT ename FROM t_emp WHERE ename LIKE 'S%';
>> ```
>> 如果左边没有%则可以使用索引

>- 对ORDER BY排序的字段设置索引 

>- 少用IS NULL和IS NOT NULL
>> ```mysql
>> SELECT ename FROM t_emp WHERE comm IS NOT NULL;
>> ```
>>此时会跳过索引，和NULL值相关的判断都不会走索引，比如判断奖金不为空值可以优化成
>>```mysql
>> SELECT ename FROM t_emp WHERE comm >=0;
>>```

>- 尽量少用 !=运算符
>>```mysql
>> SELECT ename FROM t_emp WHERE deptno!=20
>>```
>>可以优化成
>>```mysql
>> SELECT ename FROM t_emp WHERE deptno>20 AND deptno<20
>>```

>- 尽量少用OR运算符
>> ```mysql
>> SELECT ename FROM t_emp WHERE deptno=20 OR deptno=30
>>```
>>逻辑或也会使mysql跳过索引，OR运算符后的SQL会执行全表扫描,可以优化成
>>```mysql
>> SELECT ename FROM t_emp WHERE deptno=20 
>> UNION ALL
>> SELECT ename FROM t_emp WHERE deptno=30;
>>```

>- 尽量少用IN和NOT IN运算符
>>```mysql
>> SELECT ename FROM t_emp WHERE deptno IN (20,30);
>> SELECT ename FROM t_emp WHERE deptno=20 
>> UNION ALL
>> SELECT ename FROM t_emp WHERE deptno=30;
>>```

>- 在表达式的左侧使用运算符和函数都会使索引失效
>>```mysql
>> SELECT ename FROM t_emp WHERE salary*12 = 100000;
>>```
>>改成
>>```mysql
>> SELECT ename FROM t_emp WHERE salary = 100000/12;
>>```
>>```mysql
>> SELECT ename FROM t_emp WHERE year(hiredate)>=2000;
>>```
>>改成
>>```mysql
>> SELECT ename FROM t_emp WHERE hiredate>='2000-01-01 00:00:00';
>>```