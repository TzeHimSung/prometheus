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

Upload `data.csv` file. Start coding model like following one:

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