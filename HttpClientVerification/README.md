# README

## 


![](http://o7d2h0gjo.bkt.clouddn.com/2017-12-25-15142095283938.png)

![](http://o7d2h0gjo.bkt.clouddn.com/2017-12-25-15142096539053.png)

![](http://o7d2h0gjo.bkt.clouddn.com/2017-12-25-15142096602737.png)



## Examples

```
➜  goroutine ./goroutine -h
Usage of ./goroutine:
  -n    using native method or not
➜  goroutine time ./goroutine -n
./goroutine -n  0.00s user 0.00s system 0% cpu 5.465 total
➜  goroutine time ./goroutine   
./goroutine  0.00s user 0.00s system 0% cpu 0.744 total

```

## Discussion
As we have expected, the asynchronous method will be more effective, and it takes only 0.744s. But the synchronous method will cost about 5.465s. So, we can easily make a conclusion that the asynchronous method is much more effective than the synchronous one. 


