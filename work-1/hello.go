package main

import (
	"fmt"
)

func func1() {
	fmt.Println("我是 func1")
}

func func2() {
	fmt.Println("我是 func2")
}

func defer_call() {
	// 被defer的语句最后被执行，最后被defer的语句，最先被执行，通常用于释放资源。
	defer func() { fmt.Println("我是 前") }()
	defer func() { fmt.Println("我是 中") }()
	defer func() { fmt.Println("我是 后") }()
	defer func1()
	defer func2()
	panic("Error")
}

func main() {
	defer_call()
}

// 问题解决：Failed to download metadata for repo 'appstream': Cannot prepare internal mirrorlist.

// cd /etc/yum.repos.d/
// sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-*
// sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-*
// wget -O /etc/yum.repos.d/CentOS-Base.repo https://mirrors.aliyun.com/repo/Centos-vault-8.5.2111.repo
// yum clean all
// yum makecache

// root@7bab0efafad4
// docker commit -m "go tools" -a "Docker Newbee" 7bab0efafad4 work1
