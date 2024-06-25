package models

import "fmt"

type BankAccount struct {
	Name, AccNum string
	Money        float64
}

var databaseAccount []*BankAccount

func (mineAcc *BankAccount) Transfer(receiveAccNum string, amouth float64) {
	if !validationAmouth(amouth) {
		return
	}
	if amouth > mineAcc.Money {
		return
	}
	if receiveAccNum == mineAcc.AccNum {
		return
	}

	targetAcc := getBankAccFromAccNum(receiveAccNum)

	if targetAcc == nil {
		return
	}

	mineAcc.Money -= amouth
	targetAcc.Money += amouth

}

func (acc *BankAccount) Deposit(amouth float64) {
	if !validationAmouth(amouth) {
		return
	}
	acc.Money += amouth
}

func (acc *BankAccount) Withdraw(amouth float64) float64 {
	if !validationAmouth(amouth) {
		return 0.00
	}
	if amouth <= acc.Money {
		acc.Money -= amouth
		return amouth
	}
	fmt.Println("มึงจน ไอควาย")
	return 0.0
}

func validationAmouth(amouth float64) bool {
	return amouth > 0
}

func getBankAccFromAccNum(accNum string) *BankAccount {
	for i := range databaseAccount {
		targetBankAcc := databaseAccount[i]
		if targetBankAcc.AccNum == accNum {
			return targetBankAcc
		}
	}
	return nil
}

func CreateBankAcc(name, accNum string, money float64) *BankAccount {
	if getBankAccFromAccNum(accNum) == nil {
		bankAcc := BankAccount{
			Name:   name,
			AccNum: accNum,
			Money:  money,
		}
		databaseAccount = append(databaseAccount, &bankAcc)
		return &bankAcc
	}
	return nil
}
