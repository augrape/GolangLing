package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 求元素和
func sumArr(a [10]int) int {
	var sum int = 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

func main() {
	// 若想做一个真正的随机数，要种子
	// seed()种子默认是1
	//rand.Seed(1)
	rand.Seed(time.Now().Unix())

	var b [10]int
	for i := 0; i < len(b); i++ {
		// 产生一个0到1000随机数
		b[i] = rand.Intn(1000)
	}
	sum := sumArr(b)
	fmt.Printf("sum=%d\n", sum)
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
