# 如何为prometheus开发模型

**注意： 目前仅支持使用python编写的模型。**

下面展示了一个例子：读取`data.csv`，并把文件内容输出到标准输出中。

`data.csv`内容如下:

```text
id,name,value
0,python,341
1,cpp,650
2,java,810
3,csharp,821
4,golang,730
5,erlang,400
6,shell,500
7,javascript,675
8,typescript,990
9,rust,1200
```

上传`data.csv`文件。 编辑模型脚本，并命名为`sample.py`：

```python
#!/usr/env python
import os
import time

import pandas as pd

# 注意，读取数据时需要加上路径前缀
DataRootPath = './uploads/data/'


def main():
    f = pd.read_csv(DataRootPath + 'data.csv')
    for idx, series in f.iterrows():
        time.sleep(2)
        print(series['id'], series['name'], series['value'])


if __name__ == '__main__':
    main()
```

上传`sample.py`并启动。输出文件将会创建在`/runmodel/output/model-{filename}-{filetype}-{time}`目录下。

例如，当在`2021/01/02 12:34:56`时启动模型，输出文件路径为
`/runmodel/output/model-sample-py-2021-01-02-12-34-56`，可在此目录下看见输出文件`output.txt`。

`output.txt`如下：

```text
0 python 341
1 cpp 650
2 java 810
3 csharp 821
4 golang 730
5 erlang 400
6 shell 500
7 javascript 675
8 typescript 990
9 rust 1200
```

启动模型后，prometheus后端进程将会产生记录，如下所示：

```text
[INFO] 2021/01/02 12:34 Launch Model file: sample.py
[INFO] 2021/01/02 12:34 Output path check passed. Launching model...
[INFO] 2021/01/02 12:34 Model is running...
[INFO] 2021/01/02 12:34 Model launched. Create output file...
[INFO] 2021/01/02 12:34 Output file is created.
```

`Model launched`意味着模型成功运行。  
