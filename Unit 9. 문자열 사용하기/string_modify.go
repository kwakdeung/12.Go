package main

import "fmt"

func main() {
	var s1 string = "Hello, world!\n"

	s1 = "abcdefg" // �ٸ� ���ڿ��� ������ �� ����

	fmt.Println(s1[0]) // 97: ASCII �ڵ� a, �迭 ���·� �� ���ڿ� ����

	s1[0] = 'z' // ������ ����
}
