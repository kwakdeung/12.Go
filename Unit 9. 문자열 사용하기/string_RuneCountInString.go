package main

import "fmt"
import "unicode/utf8"

func main() {
	var s1 string = "�ѱ�"
	fmt.Println(utf8.RuneCountInString(s1)) // 2: ���ڿ��� ���� ���̸� ����
}
