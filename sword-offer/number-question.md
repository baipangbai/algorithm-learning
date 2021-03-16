> 这篇的算法题都是偏数学计算，会公式会模型就很容易解决，如果不会一般推导的话很难。
>
> 下面是将这类的题目都放在一起，方便一起查看，看看是否可以便于理解。以后碰到变形的题可以直接解决

1. 求解 1~n整数中1出现的次数。例如输入12，1～12这些整数中包含1的数字有1、10、11、12，1一共出现了5次

解题思路：将1～n中的个位、十位、百位、...的1出现次数相加，就是1出现的总次数

设数字是n是x位数，记n的第i位为ni，则n可以写成nxnx-1...n2n1

ni为当前位置设为cur，nxnx-1...ni+1为高位设为high，ni-1...n2n1为低位

cur 也就是ni当前位置有三种情况分别为1，0，非0和1

cur为0的情况

2304：cur为0，0010~2219 =>  229-000+1 = 230 => high * digit = 23 * 10 

cur为1的情况：

2314：cur为1，0010~2314 => 234 - 000 + 1 = 235 => hight * digit + low + 1 

cur 不为0和1的情况：

2324: cur为2，0010~2319 => 239 - 000 + 1 => 240 => (hight + 1) * digit 

```go
func countDigitOne(n int) int {
  low := 0
  high := n / 10
  cur := n % 10
  digit := 1
  res := 0
  for cur != 0 || high != 0 {
    if cur == 0 {
      res += high * digit
    } else if cur == 1 {
      res += high * digit + low + 1
    } else {
      res += (high+1) * digit
    }
    low += cur * digit
    cur = high % 10
    high = high/10
    digit *= 10
  }
  return res
}
```



