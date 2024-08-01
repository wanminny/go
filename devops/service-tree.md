# 路径要求 g.p.a 和Host 命名规范类似;
- g : group 
- p : project
- a : application 
- a代表app尽可能贴近应用的名称
- p代表project，按理说应该都不相同，但不强制
- g代表group是不可以相同的
- path要求
    - g的path 都为0
    - p的path为 /gid
    - a的path 为 /gid/pid

```shell
+----+-------+--------+------------+
| id | level | path   | node_name  |
+----+-------+--------+------------+
| 10 |     1 | 0      | ad         |
|  1 |     1 | 0      | inf        |
|  6 |     2 | /1     | cicd       |
|  2 |     2 | /1     | monitor    |
| 11 |     2 | /10    | engine     |
|  5 |     3 | /1/2   | kafka      |
|  4 |     3 | /1/2   | m3db       |
|  3 |     3 | /1/2   | prometheus |
|  8 |     3 | /1/6   | deploy     |
|  9 |     3 | /1/6   | gray       |
|  7 |     3 | /1/6   | jenkins    |
| 12 |     3 | /10/11 | schedule   |
| 13 |     3 | /10/11 | worker     |
+----+-------+--------+------------+
```

# 