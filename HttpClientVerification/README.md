
![](https://jersey.github.io/documentation/latest/images/rx-client-problem.png)

![](https://jersey.github.io/documentation/latest/images/rx-client-sync-approach.png)

![](https://jersey.github.io/documentation/latest/images/rx-client-async-approach.png)

## 测试

```
指令提示：
➜  goroutine ./goroutine -h
Usage of ./goroutine:
  -n    Whether to use native method

使用native方法：
➜  goroutine time ./goroutine -n
./goroutine -n  0.00s user 0.00s system 0% cpu 5.368 total

不使用native方法：
➜  goroutine time ./goroutine   
./goroutine  0.00s user 0.00s system 0% cpu 0.735 total

```

## 分析
异步方法只需要0.735s，而同步方法的成本约为5.368s；异步方法比同步方法更有效。


