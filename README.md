## feather

- `feather`，一种比 csv 更快，体积更小的数据格式;
- Reading feather files in go;
- 类似于`pip install feather-format`;
- fork from `https://github.com/sglyon/feather` and update,rebuild, use mod;

## 更新说明

- fork ,重新更新包，构建；
- update,rebuild, use mod;

## Feather file format:

- https://github.com/wesm/feather
- Here is the general structure of the file:

```go
<4-byte magic number "FEA1">
<ARRAY 0>
<ARRAY 1>
...
<ARRAY n>
<METADATA>
<uint32: metadata size>
<4-byte magic number "FEA1">
```

## TODO

- docs
- Other dtypes:
  - Binary
  - Date
  - Time
  - Datetime
  - timestamp

# feather arrow格式，请看下面
https://arrow.apache.org/docs/

