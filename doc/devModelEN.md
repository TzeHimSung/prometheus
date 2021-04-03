# How to develop models on prometheus

**Attention: only python models are supported.**

Here's a simple model shows how to read `data.csv` data file and print it's content to standard output.

`data.csv` looks like this:

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

Upload `data.csv` file. Start coding model like following `sample.py` file:

```python
#!/usr/env python
import os
import time

import pandas as pd

# Attention data path
DataRootPath = './uploads/data/'


def main():
    f = pd.read_csv(DataRootPath + 'data.csv')
    for idx, series in f.iterrows():
        time.sleep(2)
        print(series['id'], series['name'], series['value'])


if __name__ == '__main__':
    main()
```

Upload `sample.py` and launch it. Output file will be created at `/runmodel/output/model-{filename}-{filetype}-{time}`.

For example, when launching `sample.py` at time `2021/01/01 12:34:56`, output file dir will
be `/runmodel/output/model-sample-py-2021-01-01-12-34-56` and you can see `output.txt` below this dir.

`output.txt` looks like this:

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

Some log will be generated along this process, which can be seen at backend project.

```text
[INFO] 2021/04/03 23:12 Launch Model file: sample.py
[INFO] 2021/04/03 23:12 Output path check passed. Launching model...
[INFO] 2021/04/03 23:12 Model is running...
[INFO] 2021/04/03 23:12 Model launched. Create output file...
[INFO] 2021/04/03 23:12 Output file is created.
```

`Model launched` means that model script runs perfectly.