/**
  Copyright xuehuiit Corp. 2018 All Rights Reserved.

  http://www.xuehuiit.com

  QQ 411321681

  go语言学习-基础语法

 */
package study

import (
	"fmt"
	"strings"
)

func TestString() {
	string_sub()
	string_index()
}

//查找字符串
func string_index() {

	str := "my name is iding "
	flag := strings.Index(str, "na")

	fmt.Printf(" 查找的位置是    %d \n", flag)

}

//截取字符串
func string_sub() {

	str := "my name is iding "
	childstr := str[1:6]

	fmt.Printf(" 截取之后的字符串为：  %s  \n", childstr)
}

/**
p("Contains: ", s.Contains("source", "es")) //是否包含 true
p("Count: ", s.Count("source", "t")) //字符串出现字符的次数 2
p("HasPrefix: ", s.HasPrefix("source", "te")) //判断字符串首部 true
p("HasSuffix: ", s.HasSuffix("source", "st")) //判断字符串结尾 true
p("Index: ", s.Index("source", "e")) //查询字符串位置 1
p("Join: ", s.Join([]string{"a", "b"}, "-"))//字符串数组 连接 a-b
p("Repeat: ", s.Repeat("a", 5)) //重复一个字符串 aaaaa
p("Replace: ", s.Replace("foo", "o", "0", -1)) //字符串替换 指定起始位置为小于0,则全部替换 f00
p("Replace: ", s.Replace("foo", "o", "0", 1)) //字符串替换 指定起始位置1 f0o
p("Split: ", s.Split("a-b-c-d-e", "-")) //字符串切割 [a b c d e]
p("ToLower: ", s.ToLower("TEST")) //字符串 小写转换 source
p("ToUpper: ", s.ToUpper("source")) //字符串 大写转换 TEST
p()
p("Len: ", len("hello")) //字符串长度
p("Char:", "hello"[1]) //标取字符串中的字符，类型为byte
*/
