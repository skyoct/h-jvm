# h-jvm(hydrogen java virtual machine)
一个简单的Java虚拟机（完美运行各种姿势Hello World）

#### 简单使用
下载h-jvm
执行 ./h-jvm -jre （jre路径） 类名称
例如：./h-jvm -jre /Users/october/WorkSpace/jre/ Fibonacci

![run_example](./img/run_example.png)

编译好的下载：https://h-jvm-1252354013.cos.ap-nanjing.myqcloud.com/h-jvm.zip


#### 目前已经可以完成的功能
实现方法的调用，可以运行基本运算

以数组的方式计算斐波那契：

代码：

![code](./img/fib_array.png)


运行效果
![result](./img/fib_array_result.png)

计算斐波那契数运行如下：

代码：

![code](./img/fib_code.png)


运行效果
![result](./img/run_result.png)


#### 当前计划实现功能
* ~~类加载~~
* ~~类解析~~
* ~~运行时数据区域~~
* ~~指令集和简单解释器~~
* 异常处理
* 简易GC
* 简易线程

#### 参考资料
* 自己动手写Java虚拟机
* 深入理解Java虚拟机


