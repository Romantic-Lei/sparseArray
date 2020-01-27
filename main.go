package main
import (
	"fmt"
	"os"
	"bufio"
	"io"
	"encoding/json"
)

type ValNode struct {
	Row int 
	Col int
	Val int
}

func main() {
	// 1. 先创建一个原始数组
	var chessMap [11][11] int
	chessMap[1][2] = 1; // 黑子
	chessMap[2][3] = 2; // 蓝子

	// 2. 输出看看原始的数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	// 3. 转成稀疏数组
	// 思路 =》 
	// 1. 遍历chessMap ，发过发现有一个元素的值不为0，就创建一个node结构体，将其放到对应的结构体中
	// 2. 将其放入到对应的切片即可
	var sparseArr []ValNode

	// 标准的一个稀疏数组还有一个 记录元素的二维数组的规模(行，列和默认值)
	// 创建一个 ValNode 值节点
	valNode := ValNode{
		Row : 11,
		Col : 11,
		Val : 0,
	}

	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				// 创建一个 ValNode 值节点
				valNode := ValNode{
					Row : i,
					Col : j,
					Val : v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	// 输出稀疏数组
	// fmt.Println("当前的稀疏数组是：：：：：：：")
	// for i, valNode := range sparseArr {
	// 	fmt.Printf("%d: %d, %d, %d \n", i, valNode.row, valNode.col, valNode.val)
	// }

	WriteFile("d:\\ chess.data", sparseArr)

	// 将这个稀疏数组存盘
	// 将稀疏数组恢复成原始的数组

	// 1. 打开存盘的文件

	// 2. 这里使用稀疏数组恢复
	sparseArr1 := ReadFile("d:\\ chess.data")
	// 先创建一个原始的数组
	var chessMap2 [11][11]int

	// 遍历 sparseArr, 遍历文件每一行
	for i, valNode := range sparseArr1 {
		if i != 0 {
			// 跳过第一行记录值
			chessMap2[valNode.Row][valNode.Col] = valNode.Val
		}
	}

	// 查看是否恢复
	fmt.Println("恢复之后的数据")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}

func WriteFile(filePath string, sparseArr []ValNode) {
	// 创建一个新文件
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("os.OpenFile err = %s", err)
		return 
	}
	// 及时关闭句柄
	defer file.Close()

	write := bufio.NewWriter(file)
	for _, val := range sparseArr {

		valNode := ValNode{
			Row : val.Row,
			Col : val.Col,
			Val : val.Val,
		}

		data, err := json.Marshal(&valNode)
		if (err != nil) {
			fmt.Println("json.Marshal err =", err)
			return 
		}

		// str := fmt.Sprintf("%d: %d, %d, %d \n", i, valNode.row, valNode.col, valNode.val)
		// 写文件
		str := string(data) + "\n"
		write.WriteString(str)
	}

	// 将缓存中的数据写入到内存
	write.Flush()
}

func ReadFile(filePath string) (sparseArray []ValNode) {
	var sparseArr []ValNode
	var valNode ValNode
	// var sparseArr []byte
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("os.Open err = ", err)
		return 
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	// 循环读取文件
	for {
		// 换行作为结束
		// str, bool, err := reader.ReadLine()
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			// io.EOF表示文件的末尾
			// fmt.Println("reader.ReaderString err =", err)
			break
		}
		// 将文件保存到切片中
		// []ValNode(str)
		err = json.Unmarshal([]byte(str), &valNode)
		if err != nil {
			return 
		}
		
		sparseArr = append(sparseArr, valNode)
		// sparseArr = append(sparseArr, byte(str))
		// fmt.Print("~~~~~~~~~", str)
	}
	return sparseArr
}