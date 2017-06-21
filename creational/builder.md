# Builder 模式

Builder 模式把复杂对象的构建和对象想要表达的内容分开，这样相同的构建过程就能够用来创建不同的对象表现形式。

在Go里面，通常一个配置结构体就是用来实现这种行为，不过当你把一个结构体对象传入一个Builder方法的时候，这个方法里面会充满`if cfg.Field !=nil {...}`这样的检查代码。

## 实现

```
package car

type Speed float64

const (
    MPH Speed = 1
    KPH       = 1.60934
)

type Color string

const (
    BlueColor  Color = "blue"
    GreenColor       = "green"
    RedColor         = "red"
)

type Wheels string

const (
    SportsWheels Wheels = "sports"
    SteelWheels         = "steel"
)

type Builder interface {
    Color(Color) Builder
    Wheels(Wheels) Builder
    TopSpeed(Speed) Builder
    Build() Interface
}

type Interface interface {
    Drive() error
    Stop() error
}
```

##  使用

```
assembly := car.NewBuilder().Paint(car.RedColor)

familyCar := assembly.Wheels(car.SportsWheels).TopSpeed(50 * car.MPH).Build()
familyCar.Drive()

sportsCar := assembly.Wheels(car.SteelWheels).TopSpeed(150 * car.MPH).Build()
sportsCar.Drive()
```

## 实战

其中`builder_default_example.go`为本文档默认的例子的可执行代码；`builder_spider_example.go`为学习后的扩展应用示例。
