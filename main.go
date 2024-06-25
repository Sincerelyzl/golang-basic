package main

import (
	"fmt"
	"struct/models"
)

func main() {

	// item1 := Item{
	// 	ItemName:   "GreatSword",
	// 	Price:      10000,
	// 	Rarity:     "Epic",
	// 	PowerLevel: 120,
	// }

	// item1.changeItemName("SuperGreatSword")
	// jsondata, _ := json.MarshalIndent(&item1, "", "\t")
	// fmt.Println(string(jsondata))

	box := models.Box{Length: 2, Width: 3, Height: 4}
	fmt.Println(box.CalVolume())
	fmt.Println(box.CalSurface())

	user := models.User{Name: "KengLarb", Address: "888 Nongkai Larb"}
	user.UpdateAddress("777 Slot Machine")
	fmt.Println(user.Address)

	student := models.Student{Name: "keng", Score: 0}
	student.AddScore(20)
	fmt.Println(student.Score)

	book := models.CreatBook("Jojo", "Akira Pormungtai")
	fmt.Println(book.Summary())

	cart := models.Cart{}
	product := models.Product{Name: "Mama : ", Price: 555.305}
	fmt.Println(cart.Product)
	cart.AddProduct(product)
	fmt.Println(cart.Product)

	bankAcc1 := models.CreateBankAcc("laguna", "123456", 5000.65)
	bankAcc2 := models.CreateBankAcc("sincelz", "987654", 8000.66)
	fmt.Println(bankAcc1)
	fmt.Println(bankAcc2)

	bankAcc1.Transfer("987654", 4500)

	// 	bankAcc1.Deposit(5000)
	// 	cash := bankAcc2.Withdraw(-10000000)
	fmt.Println("Acc1CurrentMoney : ", bankAcc1.Money)
	fmt.Println("Acc2CurrentMoney : ", bankAcc2.Money)
	// 	fmt.Println(cash)
}

type Item struct {
	ItemName   string `json:"itemName"`
	Price      uint   `json:"price"`
	Rarity     string `json:"rarity"`
	PowerLevel uint   `json:"powerLevel"`
}

func (i *Item) changeItemName(name string) {
	oldItemName := i.ItemName
	i.ItemName = name
	fmt.Println("Oldname : ", oldItemName)
	fmt.Println("Newname : ", i.ItemName)
}
