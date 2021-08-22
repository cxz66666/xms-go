package bill_service

import "xms/models"

//CreateBillAndUpdateCache will add the bill to database, and refresh cache for bill, and return the id if success, return error if fail
func CreateBillAndUpdateCache(bill models.Bill) (int,error) {
	id,err:=models.AddNewBill(bill)
	if err!=nil {
		return 0,err
	}
	UpdateBasicAmount(bill.Type==models.Income,bill.UnitPrice*bill.Quantity)
	UpdateBasicCount(bill.Type==models.Income,true)
	return id,nil
}

func UpdateBasicAmount(isIncome bool,amount int)  {
	cfg:=models.GetDbConfigValue(models.BasicAmountKey)
	balance,income,outgoing :=models.GetBasicAmountFromString(cfg)
	if isIncome {
		income+=amount
	} else {
		outgoing+=amount
	}
	balance=income-outgoing
	value:=models.GenerateBasicAmountString(balance,income,outgoing)
	models.SetDbConfigValue(models.BasicAmountKey,value)
}

func UpdateBasicCount(isIncome bool,isAdding bool)  {
	cfg:=models.GetDbConfigValue(models.BasicCountKey)
	income,outgoing:=models.GetBasicCountFromString(cfg)
	if isIncome {
		if isAdding {
			income+=1
		} else {
			income-=1
		}
	} else {
		if isAdding{
			outgoing+=1
		} else {
			outgoing-=1
		}
	}
	value:=models.GenerateBasicCountString(income,outgoing)
	models.SetDbConfigValue(models.BasicCountKey,value)
}