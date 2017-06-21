# 工厂方法模式

工厂方法创建对象的设计模式可以在不必指定对象具体类型的情况下创建对象。

## 实现

这里的参考实现演示了如何创建拥有不同后端实现的数据存储对象。

### 类型

```
package data
import "io"

type Store interface {
  Open(string) (io.ReadWriteCloser, error)
}
```

### 不同的实现

```
package data

type StorageType int

const (
    DiskStorage StorageType = 1 << iota
    TempStorage
    MemoryStorage
)

func NewStore(t StorageType) Store {
    switch t {
    case MemoryStorage:
        return newMemoryStorage( /*...*/ )
    case DiskStorage:
        return newDiskStorage( /*...*/ )
    default:
        return newTempStorage( /*...*/ )
    }
}
```

用法
----

利用工厂方法，用户可以指定他们期望的存储类型。

```
s, _ := data.NewStore(data.MemoryStorage)
f, _ := s.Open("file")

n, _ := f.Write([]byte("data"))
defer f.Close()
```
