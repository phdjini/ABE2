package main

import (
	"github.com/fentec-project/gofe/abe"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SimpleChaincode struct {}

type ValInfo struct {
	Id    string  'jason:"id"'
	Val   string  'jason:"val"'
	Val   string  'jason:"policy"'
}

type KeyInfo struct {
	Id    string  'jason:"pubkey"'
	Val   string  'jason:"userkeys"'
}

func (t *SimpleChaincode) Init(stub shim.chaincodeStubInterface) pb.Response{
	fmt.Println("ABE init")
	return shim.Success(nil)
}
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response{
	fmt.Println("ABE Invoke")
	function, args :=stub.GetFunctionAndParameters()
	if function == "makeIdAndVal"{return t.invoke(stub, args)}
	else if function=="query"{return t.query(stub, args)}
	else if function =="queryById"{	return t.queryById(stub, args)	}
	return shim.Error("Invalid invoke function name")
}

func (t *SimpleChaincode) makeIdAndVal(stub shim.ChaincodeStubInterface, arg []string) pb.Response{
	id :=args[0] //id
	val:=args[1] //message
	policy:=args[2] //policy
	attribute:=args[3] //attribute
	fmt.Println("log>Input Id Value  : " +id)
	fmt.Println("log>Input message   : " +Val)
	fmt.Println("log>Input policy    : " +policy)
	fmt.Println("log>Input attribute : " +attribute)
	// encrypt  message
	a := abe.NewFAME()
	pubKey, secKey, _ := a.GenerateMasterKeys()
	msp, err := abe.BooleanToMSP(policy, false)
	if (err!=nil) {		fmt.Printf("Error in policy\n")	}
	cipher, _ := a.Encrypt(val, msp, pubKey)

	// generate attriute user keys
	user_attributes := toArray(attributes)
	userkeys, _ := a.GenerateAttribKeys(user_attributes, secKey)

	//put state DB
	valInfo :=&ValInfo{id,val, policy}
	valInfoBytes,err :=json.Marshal(valInfo)
	if err !=nil{	return shim.Error(error.Error())}
	err=stub.PutState(id,valInfoBytes)
	if err!=nil{return shim.Error(err.Error())	}
	fmt.Println("putstate complete")
	return shim.Success(nil)
}
/*
func (t *SimpleChaincode) queryById(stub shim.ChaincodeStubInterface, arg []string) pb.Response{
	fmt.Println("queryById 호출")
	id :=args[0]
	queryString :=fmt.Sprintf("{\"selector\" : {\"id\":\"%s\"}}",id)
	queryResults,err :=getQueryResultForQueryString(stub, queryString)

	if err !=nil{	return shim.Error(err,Error())}
	return shim.Success(queryResults)
}
*/
/*
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, arg []string) pb.Response{
	fmt.Println("query 호출")
	queryString :="{\"selector\" : {}}""
	fmt.Println("queryString" + queryString)
	queryResults,err :=getQueryResultForQueryString(stub, queryString)
	//	demsg, err := a.Decrypt(cipher, userkeys, pubKey)
	//	if (err!=nil) { fmt.Printf("You do not have rights!!!")	}
	//	else { fmt.Printf("Decrypted Message: %s\n",demsg) }

	if err !=nil{	return shim.Error(err,Error())	}
	return shim.Success(queryResults)
}
*/
/*
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error){
	fmt.Println("-getQueryResultForQueryString queryString :\n%s\n", queryString)
	resultsIterator, err:=stub.GetQueryResult(queryString)
	if err !=nil{	return nil, err	}
	defer resultsIterator.Close()
	buffer, err :=constructQueryResponseFromIterator(resultsIterator)
	if err !=nil{	return nil, err	}
	fmt.Println(-getQueryResultForQueryString queryString :\n%s\n", buffer.String())
	return buffer.Bytes(), nil
}
*/
/*
func constructQueryResponseFromIterator(resultsIterator Shim.StateQueryIteratorInterface) 26(* bytes.Buffer,error){
	fmt.Println("call constructQueryResponseFromIterator")
	var buffer bytes.buffer
	buffer.WriteString("[]")
	bArrayMemberAlreadyWritten :=false
	for resultsIterator.HasNext(){
		queryResponse, err :=resultIterator.Next()
		if err !=nil{	return nil, err	}
		if bArrayMemberAlreadyWritten == true {	buffer.WriteString(",")		}
		fmt.Println("what")
		fmt.Println(string(queryResponse.Value))
		fmt.Println("what")
		buffer.WriteString("{\"key\":")
		buffer.WriteString("\"")
		buffer.WriteString("queryResponse.Key")
		buffer.WriteString("\"")
		buffer.WriteString(",\"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten :=true
	}
	buffer.WriteString("]")
	return &buffer, nil56
}
*/
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err !=nil{		fmt.Printf("Error starting simple chaincode:%s", err)	}

}
