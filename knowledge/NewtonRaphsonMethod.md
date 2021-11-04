## **牛顿迭代法**

今天意外的看到了牛顿迭代法，对这个理论感到很好奇，特意去网上找了下资料，牛顿迭代法（Newton's method）又称为牛顿-拉夫逊（拉弗森）方法（Newton-Raphson method），它是牛顿在17世纪提出的一种在实数域和复数域上近似求解方程的方法。

- 多数方程不存在求根公式，因此求精确根非常困难，甚至不可能，从而寻找方程的近似根就显得特别重要。牛顿迭代法使用函数 的泰勒级数的前面几项来寻找方程 的根。牛顿迭代法是求方程根的重要方法之一，其最大优点是在方程 的单根附近具有平方收敛，而且该法还可以用来求方程的重根、复根，此时线性收敛，但是可通过一些方法变成超线性收敛。
- 上面的描述过于偏学术化，我们知道有些一元多次方程的最终解可能非常难求，如果直接求解的话，可能根本就没有解方程的办法，但是我们可以利用牛顿迭代法本质上可以求出方程的近似的一个或者多个解。

## **原理**

我们设方程函数![[公式]](https://www.zhihu.com/equation?tex=f%28x%29+%3D+m),改方程可以转化为![[公式]](https://www.zhihu.com/equation?tex=g%28x%29+%3D+f%28x%29+-+m+%3D+0) 我们只需要求出函数![[公式]](https://www.zhihu.com/equation?tex=g%28x%29+%3D+0)的解，就可以求出![[公式]](https://www.zhihu.com/equation?tex=f%28x%29+%3D+m)的解。

## **牛顿迭代公式**

设![[公式]](https://www.zhihu.com/equation?tex=r) 是![[公式]](https://www.zhihu.com/equation?tex=f%28x%29+%3D+0)的根，选取![[公式]](https://www.zhihu.com/equation?tex=x_%7B0%7D)作为![[公式]](https://www.zhihu.com/equation?tex=r)的初始近似值，则我们可以过点![[公式]](https://www.zhihu.com/equation?tex=%28x_%7B0%7D%2Cf%28x_%7B0%7D%29%29)做曲线![[公式]](https://www.zhihu.com/equation?tex=y+%3D+f%28x%29)的切线![[公式]](https://www.zhihu.com/equation?tex=L),我们知道切线与![[公式]](https://www.zhihu.com/equation?tex=x)轴有交点，我们已知切线![[公式]](https://www.zhihu.com/equation?tex=L)的方程为![[公式]](https://www.zhihu.com/equation?tex=L+%3A+y+%3D+f%28x_%7B0%7D%29+%2B+f%5E%7B%27%7D%28x_%7B0%7D%29%28x+-+x_%7B0%7D%29)我们求的它与![[公式]](https://www.zhihu.com/equation?tex=x)轴的交点为![[公式]](https://www.zhihu.com/equation?tex=x_%7B1%7D+%3D+x_%7B0%7D+-+%5Cfrac%7Bf%28x_%7B0%7D%29%7D%7Bf%5E%7B%27%7D%28x_%7B0%7D%29%7D). 我们在以![[公式]](https://www.zhihu.com/equation?tex=%28x_%7B1%7D%2Cf%28%7Bx_%7B1%7D%7D%29%29)斜率为![[公式]](https://www.zhihu.com/equation?tex=f%5E%7B%27%7D%28x_%7B1%7D%29)做斜线，求出与![[公式]](https://www.zhihu.com/equation?tex=x)轴的交点，重复以上过程直到![[公式]](https://www.zhihu.com/equation?tex=f%28x_%7Bn%7D%29)无限接近于0即可。其中第n次的迭代公式为：

![[公式]](https://www.zhihu.com/equation?tex=x_%7Bn%2B1%7D+%3D+x_%7Bn%7D+-+%5Cfrac%7Bf%28x_%7Bn%7D%29%7D%7Bf%5E%7B%27%7D%28x_%7Bn%7D%29%7D+%5C%5C)

## **题目**

以函数![[公式]](https://www.zhihu.com/equation?tex=f%28x%29+%3D+%28x-2%29%5E%7B2%7D)为例，

1. 我们可以任意取一点A![[公式]](https://www.zhihu.com/equation?tex=%284%2C4%29),在曲线上做A的切线，求得切线与![[公式]](https://www.zhihu.com/equation?tex=x)轴的交点为B。

![img](https://pic2.zhimg.com/80/v2-58d311ca11ebfbc07cabb4b5ebb18d91_1440w.jpg)



1. 在曲线上做C点的切线，交X轴与D点，在D点做X轴的垂线，交曲线于E点。我们可以看到D点比B点更加接近方程![[公式]](https://www.zhihu.com/equation?tex=f%28x%29+%3D+%28x+-+2%29+%2A+%28x+-+2%29+%3D+0) 的根![[公式]](https://www.zhihu.com/equation?tex=%EF%BC%88x+%3D+2%EF%BC%89).

![img](https://pic2.zhimg.com/80/v2-6ca2f9fe76d6194e10a455175600a17d_1440w.jpg)



1. 在曲线上做E点的切线，交X轴与F点，在F点做X轴的垂线，交曲线于G点。可以看到G点比D点更加接近方程的根.

![img](https://pic4.zhimg.com/80/v2-0953fb68d519c576ff942abcb1f474a7_1440w.jpg)



1. 按照这个方式一直迭代即可得到函数![[公式]](https://www.zhihu.com/equation?tex=f%28x%29+%3D+0)的近似解。

## **牛顿法求平方根**

我们对实数`n`求其开方，即![[公式]](https://www.zhihu.com/equation?tex=f%28x%29+%3D+x%5E%7B2%7D+-+n+%3D+0)得算法平方根。我们可以根据上述方法得到迭代`n`次的公式为:

![[公式]](https://www.zhihu.com/equation?tex=x_%7Bn%2B1%7D+%3D+x_%7Bn%7D+-+%5Cfrac%7Bf%28x_%7Bn%7D%29%7D%7Bf%5E%7B%27%7D%28x_%7Bn%7D%29%7D+%3D+x_%7Bn%7D+-+%5Cfrac%7Bx_%7Bn%7D%5E%7B2%7D-n%7D%7B2x_%7Bn%7D%7D+%3D+%5Cfrac%7B1%7D%7B2%7D%28x_%7Bn%7D+%2B+%5Cfrac%7Bn%7D%7Bx_%7Bn%7D%7D%29+%5C%5C)

- 以下为实现代码，初始时设![[公式]](https://www.zhihu.com/equation?tex=x_%7B0%7D+%3D+n).

```text
double sqrt(double c) {
    if (c < 0) {
        return -1;
    }
     
    double e = 1e-15;
    double x = c;
    double y = (x + c / x) / 2;
    while (abs(x - y) > e) {
        x = y;
        y = (x + c / x) / 2;
    }
    return x;
}

(define (sqrt-stream n)
    (letrec ([f (lambda (x)
                    (let ([next (/ (+ x (/ n x)) 2.0)])
                         (cons next (lambda () (f next)))))])
            (lambda () (f n))))
                         
(define (approx-sqrt n e)
    (stream-until (lambda (x) (< (* e 1.0) (abs (- (* n 1.0) (* x x))))) (sqrt-stream n)))
```

