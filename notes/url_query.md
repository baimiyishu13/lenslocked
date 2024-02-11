URL查询参数

GET请求没有正文，因此不能包含POST表单值。

首先，数据必须作为URL查询参数传递查询参数是附加到"?"

例子：

---

GET https://example.com/widgets?page=3
							    |        |
							 key     value

可以添加额外的 `&`

GET https://example.com/widgets?page=3&color=green

这些可能非常有用，原因有很多。我们在这里看到的一个是想象，但另一个用途是预先填充表单
