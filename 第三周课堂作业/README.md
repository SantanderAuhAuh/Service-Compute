### 1. 建立包目录
在之前划分好的工作空间`github-user`文件夹内，建立文件夹`noFunction`作为新项目的包。
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190911201749387.PNG?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dhbmd6aDI5Nw==,size_16,color_FFFFFF,t_70)
### 2.使用VScode编写功能文件
在VScode中创建文件`nothing.go`和`nothing_test.go`，在`nothing.go`中编辑文件，完成了两个功能函数`GetAreaOfCircle`和`GetPerimeterOfCircle`来分别求圆的面积和周长。`nothing_test.go`文件用于供GO语言`gotests`插件自带的轻量型测试框架使用。

将这个包命名为circle。

```
package circle

import (
	"math"
)

//GetAreaOfCircle compute the area of a circle
func GetAreaOfCircle(r float64) float64 {
	return math.Pi * math.Pow(r, 2)
}

//GetPerimeterOfCircle compute the perimeter of a circle
func GetPerimeterOfCircle(r float64) float64 {
	return 2 * math.Pi * r
}

```

**（1）保存时发现问题：出现`expected 'package', found 'EOF'`报错！！**
**解决：** 因为同在n`oFunction`文件夹内的`nothing_test.go`文件为空，所以在`gotests`自动检测时判定为出错。

编辑测试文件`nothing_test.go`：
通过创建个名字以 _test.go 结尾的，包含名为 TestXXX 且签名为 func (t *testing.T) 函数的文件来编写测试。 测试框架会运行每一个这样的函数；若该函数调用了像 t.Error 或 t.Fail 这样表示失败的函数，此测试即表示失败。

```
package circle

import (
	"math"
	"testing"
)

func TestCircle(t *testing.T) {
	cases := []float64{
		float64(1), float64(2), float64(5), float64(8), float64(0.25),
	}
	result1 := []float64{
		math.Pi * 1, math.Pi * 4, math.Pi * 25, math.Pi * 64, math.Pi * math.Pow(0.25, 2),
	}
	result2 := []float64{
		2 * math.Pi * 1, 2 * math.Pi * 2, 2 * math.Pi * 5, 2 * math.Pi * 8, 2 * math.Pi * 0.25,
	}
	for i, n := range cases {
		if GetAreaOfCircle(n) != result1[i] {
			t.Errorf("The area of the circle with radius of %f is not %f ", n, GetPerimeterOfCircle(n))
		}
		if GetPerimeterOfCircle(n) != result2[i] {
			t.Errorf("The perimeter of the with radius of %f is not %f", n, GetPerimeterOfCircle(n))
		}
	}
}

```

**（2）在`for n:=range cases`循环中对n应用`Get
AreaOfCircle`和`GetPerimeterOfCircle`函数报错！！其中cases是一个float64类型的数组**
**解决：** range函数遍历数组时会同时返回索引和数值，所以只用一个变量接受返回值只能得到索引。改为`for _, n := range cases `即可。
### 3. 对包进行测试
在VScode命令行中使用`cd`到达包所在的文件夹后，使用`go test`命令，进行测试。
观察测试结果。
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190911220559659.PNG)
结果显示通过测试。

尝试安装包并从外部调用：
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190912083017107.PNG)
发现使用noFunction的文件名安装后无法调用，将文件名改为circle。
重新达到相应文件夹内运行`go install ./circle`，结果如下图所示，表明包安装成功：
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190912125103320.PNG?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dhbmd6aDI5Nw==,size_16,color_FFFFFF,t_70)
新建一个`testForCircle.go`文件，其中有main函数，用于测试。

```
package main

import (
	"fmt"

	"github.com/github-user/circle"
)

func main() {
	var r float64
	fmt.Printf("Please input a radius:")
	fmt.Scan(&r)
	fmt.Printf("The area of circle with a radius of %f is %f\n", r, circle.GetAreaOfCircle(r))
}

```
运行结果：
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190912130625863.PNG)

完成本次实验，建立一个包并测试。

