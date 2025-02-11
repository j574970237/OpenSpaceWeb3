package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "strings"
    "time"
)

/** 
题目：
实践 POW， 编写程序（编程语言不限）用自己的昵称 + nonce，不断修改nonce 进行 sha256 Hash 运算：
1. 直到满足 4 个 0 开头的哈希值，打印出花费的时间、Hash 的内容及Hash值。
2. 再次运算直到满足 5 个 0 开头的哈希值，打印出花费的时间、Hash 的内容及Hash值。
**/
func main() {
    nickname := "JackZhu"
    nonce := 0
    targetPrefix4 := "0000"

    // 寻找满足4个0开头的哈希值
    startTime := time.Now()
    for {
        hash := sha256.Sum256([]byte(fmt.Sprintf("%s%d", nickname, nonce)))
        hashStr := hex.EncodeToString(hash[:])
        if strings.HasPrefix(hashStr, targetPrefix4) {
            spendTime := time.Since(startTime)
            fmt.Printf("寻找满足4个0开头的哈希值:\n")
            fmt.Printf("花费的时间: %s\n", spendTime)
            fmt.Printf("Hash的内容: %s%d\n", nickname, nonce)
            fmt.Printf("Hash值: %s\n", hashStr)
            break
        }
        nonce++
    }

    // 重置随机数和目标前缀
    nonce = 0
	targetPrefix5 := "00000"

    // 寻找满足5个0开头的哈希值
	startTime = time.Now()
    for {
        hash := sha256.Sum256([]byte(fmt.Sprintf("%s%d", nickname, nonce)))
        hashStr := hex.EncodeToString(hash[:])
        if strings.HasPrefix(hashStr, targetPrefix5) {
            spendTime := time.Since(startTime)
            fmt.Printf("寻找满足5个0开头的哈希值:\n")
            fmt.Printf("花费的时间: %s\n", spendTime)
            fmt.Printf("Hash的内容: %s%d\n", nickname, nonce)
            fmt.Printf("Hash值: %s\n", hashStr)
            break
        }
        nonce++
    }
}