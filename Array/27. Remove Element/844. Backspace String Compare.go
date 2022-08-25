// 思路1：栈
func backspaceCompare(s string, t string) bool {
    return build(s) == build(t)
}
// 新建一个build函数，内部实现函数创建
func build(str string) string {
    s := make([]byte, 0)     // 用str[i]索引时用byte类型，直接索引v用rune
    // s := []byte{}         // 官方题解的写法
    // s := make([]rune, 0)  // 此处用rune时，第11行报错：cannot use str[i] (type byte) as type rune in append (solution.go)
    for i := range str {     // 省略索引值，只使用索引i的for range写法，s要定义为byte
        if str[i] != '#' {
            s = append(s, str[i])
        } else if len(s) > 0 {
            s = s[:len(s)-1]  // 相当于 s 出栈一个元素 (字符)
        } 
    }
    /* 另一种省略索引的for range写法，s要定义为rune
    for _, v := range str {
        if v != '#' {
            s = append(s, v)
        } else if len(s) > 0 {
            s = s[:len(s)-1]
        }
    } */

// 思路2：双指针
func backspaceCompare(s string, t string) bool {
    i, j := len(s)-1, len(t)-1
    skipS, skipT := 0, 0
    for i >= 0 || j >= 0 {  // 注意>=
        for i >= 0 {
            if s[i] == '#' {
                skipS++
                i--       // 忘记写i--，导致遇到 '#' 循环一直无法终止，导致runtime error
            } else if skipS > 0 {
                skipS--
                i--
            } else {
                break     // 不加break会继续循环，不会跳出
            }
        }
        for j >= 0 {
            if t[j] == '#' {
                skipT++        
                j--       // 忘记写j--，导致遇到 '#' 循环一直无法终止，导致runtime error
            } else if skipT > 0 {
                skipT--
                j--
            } else {
                break     // 不加break会继续循环，不会跳出
            }
        }
        if i >= 0 && j >= 0 {  // 注意 i，j都 >=，代表两个元素可以比较
            if s[i] != t[j] {
                return false
            } 
        } else if i >= 0 || j >= 0 {  // else if后用 ||，代表有一个<0，长度不等一定是false
            return false
        }
        i--
        j--                   // 注意在判断所有除了false的情况都需要i--， j--，所以写在if外
    }
    return true
}