package util

import (
	"bufio"
	"fmt"
	"os"
)

type FamilyAccount struct {
	// reader用于读取字符串
	reader *bufio.Reader

	// Declarate a variable to store the user's input
	key string

	// 变量detail用于存放收支明细
	detail string

	// 变量income用于存储收入金额
	income float64

	// 变量espense用于存储支出金额
	expense float64

	// 变量balance用于存储账户余额
	balance float64

	// 变量note用于存储收支说明
	note string

	// err用于存放error
	err error

	// flag用于标志是否有收支登记
	flag bool

	// choice用于标志是否退出
	choice string
}

// New construct a new FamilyAccount
func New() *FamilyAccount {

	return &FamilyAccount{
		reader:  bufio.NewReader(os.Stdin),
		key:     "",
		detail:  fmt.Sprintf("%-2s\t%-10s\t%-10s\t%-2s\n", "收支", "收支金额", "账户余额", "说明"),
		balance: 10000.00,
		note:    "",
	}
}

// ShowMenu shows the menu in a loop
func (famacc *FamilyAccount) ShowMenu() {
loop:
	for {
		fmt.Println("\n--------------------家庭收支记账软件--------------------")
		fmt.Println("                       1 收支明细")
		fmt.Println("                       2 登记收入")
		fmt.Println("                       3 登记支出")
		fmt.Println("                       4 退出软件")
		fmt.Print("请选择（1-4）：")

		fmt.Scanln(&famacc.key)

		switch famacc.key {
		case "1":
			famacc.ShowDetail()
		case "2":
			famacc.RecordIncome()
		case "3":
			famacc.RecordExpense()
		case "4":
			if famacc.Exit() {
				break loop
			}
			continue
		}
	}
}

// ShowDetail shows the details of the family account
func (famacc *FamilyAccount) ShowDetail() {
	fmt.Println("\n--------------------当前收支明细--------------------")

	if famacc.flag {
		fmt.Println(famacc.detail)
	} else {
		fmt.Println("没有收支登记！来一笔把！")
	}
}

// RecordIncome records income
func (famacc *FamilyAccount) RecordIncome() {
	fmt.Println("\n--------------------登记收入--------------------")

	fmt.Print("请输入收入金额：")
	fmt.Scanln(&famacc.income)
	famacc.balance += famacc.income

	fmt.Print("请输入收入说明：")
	famacc.note, famacc.err = famacc.reader.ReadString('\n')
	if famacc.err != nil {
		fmt.Println(famacc.err)
		return
	}

	fmt.Println("\n--------------------登记完成--------------------")
	fmt.Printf("\t本次收入金额： %v\n\t本次收入说明： %s\n", famacc.income, famacc.note)
	famacc.detail += fmt.Sprintf("%-2s\t%-10.2f\t%-10.2f\t%-s", "收入", famacc.income, famacc.balance, famacc.note)

	famacc.flag = true
}

// RecordExpense records expense
func (famacc *FamilyAccount) RecordExpense() {
	fmt.Println("\n--------------------登记支出--------------------")

	fmt.Print("请输入支出金额：")
	fmt.Scanln(&famacc.expense)

	if famacc.expense > famacc.balance {
		fmt.Println("余额不足！")
		return
	}

	famacc.balance -= famacc.expense

	fmt.Print("请输入支出说明：")
	famacc.note, famacc.err = famacc.reader.ReadString('\n')
	if famacc.err != nil {
		fmt.Println(famacc.err)
		return
	}
	fmt.Println("\n--------------------登记完成--------------------")
	fmt.Printf("\t本次支出金额： %v\n\t本次支出说明： %s\n", famacc.expense, famacc.note)
	famacc.detail += fmt.Sprintf("%-2s\t%-10.2f\t%-10.2f\t%-s", "支出", famacc.expense, famacc.balance, famacc.note)
	famacc.flag = true
}

// Exit exits the program
func (famacc *FamilyAccount) Exit() bool {
	for {
		fmt.Print("确认退出y/n?:")
		fmt.Scanln(&famacc.choice)
		switch famacc.choice {
		case "y":
			fmt.Println("\n您已退出家庭记账软件的使用...")
			return true
		case "n":
			return false
		default:
			fmt.Println("输入错误!")
			continue
		}
	}
}
