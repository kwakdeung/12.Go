package main

import "fmt"

func main() {
	var s1 string = "�ѱ�"
	var s2 string = "Hello"

	fmt.Println(len(s1)) // 6: UTF-8 ���ڵ��� ����Ʈ �����̹Ƿ� 6
	fmt.Println(len(s2)) // 5: ���ĺ� 5�����̹Ƿ� 5
}
