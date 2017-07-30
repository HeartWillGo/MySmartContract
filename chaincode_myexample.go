package main

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

//中央银行
type CenterBank struct {
	Name        string `json:"name"`        //中央银行名称
	TotalNumber int    `json:"totalnumber"` //发行货币总数额
	RestNumber  int    `json:"restnumber"`  //账户余额
	ID          int    `json:"id"`          //中央银行ID
}

//银行
type Bank struct {
	Name        string `json:"name"`        //银行名称
	TotalNumber int    `json:"totalnumber"` //接收货币总数额
	RestNumber  int    `json:"fromtype"`    //账户余额
	ID          int    `json:"id"`          //银行ID
}

//企业
type Company struct {
	Name   string `json:"name"`   //企业名称
	Number int    `json:"number"` //账户余额
	ID     int    `json:"id"`     //企业ID

}

//交易内容
type Transaction struct {
	FromType string `json:"fromtype"` //发送方角色 centerBank:0,Bank:1,Company:2
	FromID   int    `json:"fromid"`   //发送方 ID
	ToType   string `json:"totype"`   //接收方角色 Bank:1,Company:2
	ToID     int    `json:"toid"`     //接收方 ID
	Time     string `json:"time"`     //交易时间
	Number   int    `json:"number"`   //交易数额
	ID       int    `json:"id"`       //交易 ID
}

var center CenterBank

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Init")
	_, args := stub.GetFunctionAndParameters()
	var centerBankName string  // 中央银行名称
	var TotalNumber_center int //  发行货币总数额
	var RestNumber_center int  //账户余额
	var ID_center int          //中央银行ID
	var err error
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// Initialize the chaincode
	centerBankName = args[0]

	TotalNumber_center, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：TotalNumber_center")
	}
	RestNumber_center, err = strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：RestNumber_center")
	}
	ID_center, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：ID_center")
	}

	fmt.Printf("centerBankName = %d, TotalNumber_center = %d, RestNumber_center=%d,ID_center=%d\n", centerBankName, TotalNumber_center, RestNumber_center, ID_center)

	center.Name = centerBankName
	center.TotalNumber = TotalNumber_center
	center.RestNumber = RestNumber_center
	center.ID = ID_center

	jsons, errs := json.Marshal(center) //转换成JSON返回的是byte[]

	if errs != nil {
		return shim.Error(errs.Error())
	}
	// Write the state to the ledger
	err = stub.PutState(args[3], jsons)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf(" init success \n")
	return shim.Success(nil)
}

func (t *SimpleChaincode) CreateBank(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 CreateBank")

	var Name string     //  银行名称
	var TotalNumber int //  接收货币总数额
	var RestNumber int  //  账户余额
	var ID int          //  中央银行ID

	var bank Bank
	var err error

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// Initialize the chaincode
	Name = args[0]

	TotalNumber, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：TotalNumber ")
	}
	RestNumber, err = strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：RestNumber ")
	}
	ID, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：ID ")
	}

	fmt.Printf(" Name = %d, TotalNumber  = %d, RestNumber =%d,ID =%d\n", Name, TotalNumber, RestNumber, ID)

	bank.Name = Name
	bank.TotalNumber = TotalNumber
	bank.RestNumber = RestNumber
	bank.ID = ID

	jsons, errs := json.Marshal(bank) //转换成JSON返回的是byte[]
	if errs != nil {
		return shim.Error(errs.Error())
	}

	// Write the state to the ledger
	err = stub.PutState(args[3], jsons)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf(" CeateBank success \n")
	return shim.Success(nil)
}

//createCompany
func (t *SimpleChaincode) CreateCompany(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 CreateCompany")

	var Name_company string //  银行名称
	var Number int          //  账户余额
	var ID_company int      //  ID
	var err error
	var company Company

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	// Initialize the chaincode
	Name_company = args[0]

	Number, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	ID_company, err = strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：ID_company ")
	}

	fmt.Printf(" Name_company = %d, Number  = %d,ID_company =%d\n", Name_company, Number, ID_company)

	company.Name = Name_company
	company.Number = Number
	company.ID = ID_company

	jsons, errs := json.Marshal(company) //转换成JSON返回的是byte[]
	if errs != nil {
		return shim.Error(errs.Error())
	}
	// Write the state to the ledger
	err = stub.PutState(args[2], jsons)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("CreateCompany \n")

	return shim.Success(nil)
}

//issueCoin 发行货币
func (t *SimpleChaincode) IssueCoin(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 IssueCoin")

	var Number int        // 发行的数量
	var ID_trans int      //交易ID
	var trans Transaction //交易过程
	var err error
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// Initialize the chaincode

	Number, err = strconv.Atoi(args[0])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	ID_trans, err = strconv.Atoi(args[1])

	if err != nil {
		return shim.Error("Expecting integer value for asset holding：ID_trans ")
	}

	fmt.Printf("  Number  = %d ,ID_trans = %d \n", Number, ID_trans)

	trans.FromType = "0"
	trans.FromID = 0
	trans.ToType = "0"
	trans.ToID = 0

	cur_time := time.Now()
	trans.Time = cur_time.String()
	trans.Number = Number
	trans.ID = ID_trans

	center.RestNumber = center.RestNumber + Number

	jsons, errs := json.Marshal(trans) //转换成JSON返回的是byte[]
	if errs != nil {
		return shim.Error(errs.Error())
	}
	// Write the state to the ledger
	err = stub.PutState(args[1], jsons)
	if err != nil {
		return shim.Error(err.Error())
	}

	jsons_center, errs2 := json.Marshal(center) //转换成JSON返回的是byte[]
	if errs2 != nil {
		return shim.Error(errs2.Error())
	}
	// Write the state to the ledger
	err = stub.PutState("0", jsons_center)

	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf(" IssueCoin success \n")
	return shim.Success(nil)
}

//issueCoinToBank  发行货币至商业银行
func (t *SimpleChaincode) issueCoinToBank(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 IssueCoin")

	var Number int                // 发行的数量
	var To_ID int                 //接收方ID
	var ID_trans int              //交易ID
	var trans_to_bank Transaction //交易过程
	var toBank Bank               //商业银行
	var err error
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// Initialize the chaincode

	Number, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	To_ID, err = strconv.Atoi(args[0])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：To_ID  ")
	}

	ID_trans, err = strconv.Atoi(args[2])

	if err != nil {
		return shim.Error("Expecting integer value for asset holding：ID_trans ")
	}

	fmt.Printf("  Number  = %d ,To_ID =%d , ID_trans=%d\n", Number, To_ID, ID_trans)

	trans_to_bank.FromType = "0"
	trans_to_bank.FromID = 0
	trans_to_bank.ToType = "1"
	trans_to_bank.ToID = To_ID

	cur_time := time.Now()

	trans_to_bank.Time = cur_time.String()

	trans_to_bank.Number = Number
	trans_to_bank.ID = ID_trans

	center.RestNumber = center.RestNumber - Number

	toBankInfo, erro := stub.GetState(args[0])
	if erro != nil {
		return shim.Error(erro.Error())
	}

	//将byte的结果转换成struct
	err = json.Unmarshal(toBankInfo, &toBank)
	toBank.TotalNumber = Number
	toBank.RestNumber = toBank.RestNumber + Number

	fmt.Printf("  toBankInfo  = %d  \n", toBankInfo)

	jsons, errs := json.Marshal(trans_to_bank) //转换成JSON返回的是byte[]
	if errs != nil {
		return shim.Error(errs.Error())
	}
	ID_trans_string := strconv.Itoa(ID_trans)
	// Write the state to the ledger
	err = stub.PutState(ID_trans_string, jsons)
	if err != nil {
		return shim.Error(err.Error())
	}

	jsons_toBank, errs2 := json.Marshal(toBank) //转换成JSON返回的是byte[]
	if errs2 != nil {
		return shim.Error(errs2.Error())
	}
	toBankID_string := strconv.Itoa(toBank.ID)
	// Write the state to the ledger
	err = stub.PutState(toBankID_string, jsons_toBank)
	if err != nil {
		return shim.Error(err.Error())
	}

	jsons_center, errs3 := json.Marshal(center) //转换成JSON返回的是byte[]
	if errs3 != nil {
		return shim.Error(errs3.Error())
	}
	centerID_string := strconv.Itoa(center.ID)
	// Write the state to the ledger
	err = stub.PutState(centerID_string, jsons_center)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("  issueCoinToBank success \n")
	return shim.Success(nil)
}

//商业银行转账到企业  issueCoinToCp
func (t *SimpleChaincode) issueCoinToCp(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 IssueCoin")

	var Number int             // 数量
	var From_ID int            // 商业银行ID
	var To_ID int              //接收方ID
	var ID int                 //交易ID
	var bank_to_cp Transaction //交易过程
	var bankFrom Bank          //商业银行
	var cpTo Company           //企业
	var err error
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// Initialize the chaincode

	From_ID, err = strconv.Atoi(args[0])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：From_ID ")
	}
	Number, err = strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：Number ")
	}
	To_ID, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：To_ID  ")
	}

	ID, err = strconv.Atoi(args[3])

	if err != nil {
		return shim.Error("Expecting integer value for asset holding：ID_trans ")
	}

	fmt.Printf("  Number  = %d ,To_ID =%d , ID_trans=%d\n", Number, To_ID, ID)

	bank_to_cp.FromType = "1"
	bank_to_cp.FromID = From_ID
	bank_to_cp.ToType = "2"
	bank_to_cp.ToID = To_ID

	cur_time := time.Now()
	bank_to_cp.Time = cur_time.String()

	bank_to_cp.Number = Number
	bank_to_cp.ID = ID

	BankFromBytes, erro := stub.GetState(args[0])
	if erro != nil {
		return shim.Error(erro.Error())
	}

	//将byte的结果转换成struct
	err = json.Unmarshal(BankFromBytes, &bankFrom)
	bankFrom.RestNumber = bankFrom.RestNumber - Number

	jsons_bank, errs := json.Marshal(bankFrom) //转换成JSON返回的是byte[]
	if errs != nil {
		return shim.Error(errs.Error())
	}
	bankFromID_string := strconv.Itoa(bankFrom.ID)

	// Write the state to the ledger
	err = stub.PutState(bankFromID_string, jsons_bank)

	companyToBytes, erro1 := stub.GetState(args[1])
	if erro1 != nil {
		return shim.Error(erro1.Error())
	}
	//将byte的结果转换成struct
	err = json.Unmarshal(companyToBytes, &cpTo)
	cpTo.Number = cpTo.Number + Number

	jsons_cp, errs2 := json.Marshal(cpTo) //转换成JSON返回的是byte[]
	if errs2 != nil {
		return shim.Error(errs2.Error())
	}
	cpToID_string := strconv.Itoa(cpTo.ID)
	// Write the state to the ledger
	err = stub.PutState(cpToID_string, jsons_cp)

	jsons, errs3 := json.Marshal(bank_to_cp) //转换成JSON返回的是byte[]
	if errs3 != nil {
		return shim.Error(errs3.Error())
	}
	ID_string := strconv.Itoa(ID)
	// Write the state to the ledger
	err = stub.PutState(ID_string, jsons)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

//getBanks
func (t *SimpleChaincode) getBanks(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 getBanks")

	var Bank_ID string // 商业银行ID
	var bank_info Bank
	var err error
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// Initialize the chaincode

	Bank_ID = args[0]

	BankInfo, erro := stub.GetState(Bank_ID)
	if erro != nil {
		return shim.Error(erro.Error())
	}
	//将byte的结果转换成struct
	err = json.Unmarshal(BankInfo, &bank_info)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("  BankInfo  = %d  \n", BankInfo)

	return shim.Success(nil)
}

//getCompanys
func (t *SimpleChaincode) getCompanys(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 getCompanys")

	var CP_ID string // 企业ID
	var company_info Company
	var err error
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// Initialize the chaincode

	CP_ID = args[0]

	company_info_bytes, erro := stub.GetState(CP_ID)
	if erro != nil {
		return shim.Error(erro.Error())
	}

	//将byte的结果转换成struct

	err = json.Unmarshal(company_info_bytes, &company_info)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("  BankInfo  = %d  \n", company_info_bytes)

	return shim.Success(nil)
}

//getTransactions
func (t *SimpleChaincode) getTransactions(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 getCompanys")

	var trans_ID string // 企业ID
	var trans_info Transaction
	var err error
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// Initialize the chaincode

	trans_ID = args[0]

	trans_info_bytes, erro := stub.GetState(trans_ID)
	if erro != nil {
		return shim.Error(erro.Error())
	}

	//将byte的结果转换成struct

	err = json.Unmarshal(trans_info_bytes, &trans_info)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("  trans_info_bytes  = %d  \n", trans_info_bytes)

	return shim.Success(nil)
}

//getCenterBank
func (t *SimpleChaincode) getCenterBank(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 getCenterBank")

	var Center_ID string // 企业ID
	var center_info CenterBank
	var err error
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// Initialize the chaincode

	Center_ID = args[0]

	center_info_bytes, erro := stub.GetState(Center_ID)
	if erro != nil {
		return shim.Error(erro.Error())
	}

	//将byte的结果转换成struct

	err = json.Unmarshal(center_info_bytes, &center_info)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("  center_info_bytes  = %d  \n", center_info_bytes)

	return shim.Success(nil)
}

//transfer
func (t *SimpleChaincode) transfer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("ex02 getCenterBank")

	var From_ID int // 转账方ID
	var To_ID int   //接收方ID
	var number int  //转账金额
	var fromCP Company
	var toCP Company
	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	// Initialize the chaincode

	From_ID, err = strconv.Atoi(args[0])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：From_ID  ")
	}
	To_ID, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：To_ID  ")
	}
	number, err = strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding：number ")
	}

	fromID_string := strconv.Itoa(From_ID)
	from_cp_info_bytes, erro := stub.GetState(fromID_string)
	if erro != nil {
		return shim.Error(erro.Error())
	}

	//将byte的结果转换成struct

	err = json.Unmarshal(from_cp_info_bytes, &fromCP)

	fmt.Printf("  from_cp_info_bytes  = %d  \n", from_cp_info_bytes)

	To_ID_string := strconv.Itoa(To_ID)
	to_cp_info_bytes, erro1 := stub.GetState(To_ID_string)
	if erro1 != nil {
		return shim.Error(erro1.Error())
	}

	//将byte的结果转换成struct

	err = json.Unmarshal(to_cp_info_bytes, &toCP)

	fmt.Printf("  to_cp_info_bytes  = %d  \n", to_cp_info_bytes)

	from_cp_old_num := fromCP.Number
	if from_cp_old_num <= number {
		return shim.Error("money no enough")
	}

	fromCP.Number = from_cp_old_num - number

	to_cp_old_num := toCP.Number
	toCP.Number = to_cp_old_num + number

	jsons_from, errs := json.Marshal(fromCP) //转换成JSON返回的是byte[]
	if errs != nil {
		return shim.Error(errs.Error())
	}
	fromCPID_string := strconv.Itoa(fromCP.ID)
	// Write the state to the ledger
	err = stub.PutState(fromCPID_string, jsons_from)
	if err != nil {
		return shim.Error(err.Error())
	}

	jsons_to, errs2 := json.Marshal(toCP) //转换成JSON返回的是byte[]
	if errs2 != nil {
		return shim.Error(errs2.Error())
	}
	toCPID_string := strconv.Itoa(toCP.ID)
	// Write the state to the ledger
	err = stub.PutState(toCPID_string, jsons_to)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf(" transfer success \n")
	return shim.Success(nil)
}

// Deletes an entity from state
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "invoke" {
		// Make payment of X units from A to B
		return t.invoke(stub, args)
	} else if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	} else if function == "query" {
		// the old "Query" is now implemtned in invoke
		return t.query(stub, args)
	} else if function == "CreateBank" {
		// the old "Query" is now implemtned in invoke
		return t.CreateBank(stub, args)
	} else if function == "CreateCompany" {
		// the old "Query" is now implemtned in invoke
		return t.CreateCompany(stub, args)
	} else if function == "getBanks" {
		// the old "Query" is now implemtned in invoke
		return t.getBanks(stub, args)
	} else if function == "getCenterBank" {
		// the old "Query" is now implemtned in invoke
		return t.getCenterBank(stub, args)
	} else if function == "getCompanys" {
		// the old "Query" is now implemtned in invoke
		return t.getCompanys(stub, args)
	} else if function == "getTransactions" {
		// the old "Query" is now implemtned in invoke
		return t.getTransactions(stub, args)
	} else if function == "IssueCoin" {
		// the old "Query" is now implemtned in invoke
		return t.IssueCoin(stub, args)
	} else if function == "issueCoinToBank" {
		// the old "Query" is now implemtned in invoke
		return t.issueCoinToBank(stub, args)
	} else if function == "issueCoinToCp" {
		// the old "Query" is now implemtned in invoke
		return t.issueCoinToCp(stub, args)
	} else if function == "transfer" {
		// the old "Query" is now implemtned in invoke
		return t.transfer(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	return shim.Success(nil)
}

// Deletes an entity from state

// query callback representing the query of a chaincode
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A string // Entities
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, erro := stub.GetState(A)
	if erro != nil {
		return shim.Error(erro.Error())
	}
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
