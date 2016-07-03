package cutter

import (
	"reflect"

	"testing"
)

func TestSentences(t *testing.T) {
	var tests = []struct {
		punct bool
		text  string
		out   []string
	}{
		/**
		 * Chinese sentences.
		 **/

		{true, "你好，世界。", []string{"你好，", "世界。"}},
		{false, "你好，世界。", []string{"你好", "世界"}},

		{true, "我在写一个分句的程序，把一段话分割成一句一句的", []string{"我在写一个分句的程序，", "把一段话分割成一句一句的"}},
		{false, "我在写一个分句的程序，把一段话分割成一句一句的", []string{"我在写一个分句的程序", "把一段话分割成一句一句的"}},

		{true, "“真是渣渣！！！！！我好无语啊。。。。。都是什么啊，，，，”", []string{"真是渣渣！", "我好无语啊。", "都是什么啊，"}},
		{false, "“真是渣渣！！！！！我好无语啊。。。。。都是什么啊，，，，”", []string{"真是渣渣", "我好无语啊", "都是什么啊"}},
	}

	for k, test := range tests {
		s := Sentences(test.text, test.punct)
		if !reflect.DeepEqual(s, test.out) {
			t.Errorf("[%d]get: [% #v] ,expect: [% #v]\n", k+1, s, test.out)
			t.Fail()
		}
	}

}
