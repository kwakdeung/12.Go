package main

import "fmt"

func main() {
	var s1 string = "�ѱ�"
	var s2 string = "�ѱ�"
	var s3 string = "Go"

	fmt.Println(s1 == s2)          // true: �� ���ڿ��� �����Ƿ� true
	fmt.Println(s1 + s3 + s2)      // �ѱ�Go�ѱ�
	fmt.Println("�ȳ��ϼ���" + s1) // �ȳ��ϼ����ѱ�
}
