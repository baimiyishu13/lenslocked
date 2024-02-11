#### 创建 SQL 表

**创建表：指定需要存储什么**

**建立一个用户表：**

| **id** | **age** | **firstname** | **lastname** | **email** |
| ------------ | ------------- | ------------------- | ------------------ | --------------- |

```
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    age INTEGER,
    firstname VARCHAR(255),
    lastname VARCHAR(255),
    email VARCHAR(255)
);
```

遵循

```pgsql
CREATE TABLE table_name (
    field1_name field1_type constraints,
    field2_name field2_type (args) constraints,
    ...
);
```

`CREATE TABLE table_name`: 这是创建表的 SQL 命令的开头。`table_name` 是你要创建的表的名称。

* `field1_name`, `field2_name`: 这些是列的名称。
* `field1_type`, `field2_type`: 这些是列的数据类型。例如，`VARCHAR(255)`、`INTEGER`、`DATE` 等。
* `(args)`: 对于某些数据类型，你可能需要提供参数。例如，如果你使用 `VARCHAR` 类型，你可以指定字符的最大长度。这些参数是可选的，根据数据类型的不同可能会有所不同。
* `constraints`: 这些是列的约束条件，例如 `PRIMARY KEY`、`NOT NULL`、`UNIQUE`、`DEFAULT` 等。这些约束是可选的，根据需要可以添加到列的定义中。
