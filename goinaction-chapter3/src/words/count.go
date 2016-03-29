// 一个目录下只能有一个包名，函数的调用，静态方法的使用都是通过包名来操作的。
package words

import "strings"

// 计算文本中的单词个数，并返回结果。
func CountWords(text string) (count int) {
    count = len(strings.Fields(text))

    // 当返回的变量命名化，且已有赋值时，在return的时候是可以不跟返回变量的
    return count
}
