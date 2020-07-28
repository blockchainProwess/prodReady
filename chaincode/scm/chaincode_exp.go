package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// SmartContract struct
type SmartContract struct {
}

//ProducerTransaction struct
type ProducerTransaction struct {
	TxID                     string `json:tx_id`
	AssetID                  string `json:"asset_id"`
	TotalWeightShippedPerLot string `json:"total_weight_shipped_per_lot"`
	BoxesShippedPerLot       string `json:"boxes_shipped_per_lot"`
	DateofProduction         string `json:"date_of_production"`
	DateofHarvest            string `json:"date_of_harvest"`
	NameoftheFarm            string `json:"name_of_the_farm"`
	ProductionMethod         string `json:"production_method"`
	FarmingCountry           string `json:"farming_country"`
	SpeciesName              string `json:"species_name"`
	DispatchDate             string `json:"pr_dispatch_date"`
}

//ProcessorTransaction struct
type ProcessorTransaction struct {
	TxID                string `json:tx_id`
	AssetID             string `json:"asset_id"`
	ReceptionDate       string `json:"po_reception_date"`
	DispatchDate        string `json:"po_dispatch_date"`
	ProcessingMethod    string `json:"processing_method"`
	CountryofProcessing string `json:"country_of_processing"`
	ProductForm         string `json:"product_form"`
	ProductCondition    string `json:"product_condition"`
	SizeGrade           string `json:size_grade`
	SellbyDate          string `json:sell_by_date`
	ProcessorName       string `json:processor_name`
}

//DistributorTransaction struct
type DistributorTransaction struct {
	TxID                  string `json:tx_id`
	AssetID               string `json:"asset_id"`
	ReceptionDate         string `json:"di_reception_date"`
	DispatchDate          string `json:"di_dispatch_date"`
	ColdStorageConditions string `json:cold_storage_conditions`
	DistributorName       string `json:distributor_name`
}

//GrocerTransaction struct
type GrocerTransaction struct {
	TxID          string `json:"tx_id"`
	AssetID       string `json:"asset_id"`
	ReceptionDate string `json:"re_reception_date"`
	GrocerName    string `json:"grocer_name"`
}

// Init function
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke function
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger
	if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "recordProducerTransaction" {
		return s.recordProducerTransaction(APIstub, args)
	} else if function == "recordProcessorTransaction" {
		return s.recordProcessorTransaction(APIstub, args)
	} else if function == "recordDistributorTransaction" {
		return s.recordDistributorTransaction(APIstub, args)
	} else if function == "recordGrocerTransaction" {
		return s.recordGrocerTransaction(APIstub, args)
	} else if function == "queryTransaction" {
		return s.queryTransaction(APIstub, args)
	} else if function == "queryTransactionHistory" {
		return s.queryTransactionHistory(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

/*
 * The initLedger method *
Will add test data (10 asset catches)to our network
*/
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {

	producertx := []ProducerTransaction{
		ProducerTransaction{TxID: "ASID001Tx01", AssetID: "ASID001", TotalWeightShippedPerLot: "100", BoxesShippedPerLot: "10", DateofProduction: "01-01-2020", DateofHarvest: "01-02-2020", NameoftheFarm: "XFarmers", ProductionMethod: "Farmed", FarmingCountry: "India", SpeciesName: "Shrimp", DispatchDate: "02-02-2020"},
		ProducerTransaction{TxID: "ASID002Tx01", AssetID: "ASID002", TotalWeightShippedPerLot: "100", BoxesShippedPerLot: "10", DateofProduction: "01-01-2020", DateofHarvest: "01-02-2020", NameoftheFarm: "XFarmers", ProductionMethod: "Farmed", FarmingCountry: "India", SpeciesName: "Fresh Crab", DispatchDate: "02-02-2020"},
		ProducerTransaction{TxID: "ASID003Tx01", AssetID: "ASID003", TotalWeightShippedPerLot: "100", BoxesShippedPerLot: "10", DateofProduction: "01-01-2020", DateofHarvest: "01-02-2020", NameoftheFarm: "XFarmers", ProductionMethod: "Farmed", FarmingCountry: "India", SpeciesName: "Cat Fish", DispatchDate: "02-02-2020"},
		ProducerTransaction{TxID: "ASID004Tx01", AssetID: "ASID004", TotalWeightShippedPerLot: "100", BoxesShippedPerLot: "10", DateofProduction: "01-01-2020", DateofHarvest: "01-02-2020", NameoftheFarm: "XFarmers", ProductionMethod: "Farmed", FarmingCountry: "India", SpeciesName: "Clam", DispatchDate: "02-02-2020"},
		ProducerTransaction{TxID: "ASID005Tx01", AssetID: "ASID005", TotalWeightShippedPerLot: "100", BoxesShippedPerLot: "10", DateofProduction: "01-01-2020", DateofHarvest: "01-02-2020", NameoftheFarm: "XFarmers", ProductionMethod: "Farmed", FarmingCountry: "India", SpeciesName: "Shrimp", DispatchDate: "02-02-2020"},
	}
	processortx := []ProcessorTransaction{
		ProcessorTransaction{TxID: "ASID001Tx02", AssetID: "ASID001", ReceptionDate: "03-02-2020", DispatchDate: "04-02-2020", ProcessingMethod: "XYZ", CountryofProcessing: "China", ProductForm: "Canned", ProductCondition: "Frozen", SizeGrade: "10", SellbyDate: "15-03-2020", ProcessorName: "XProcessors"},
		ProcessorTransaction{TxID: "ASID002Tx02", AssetID: "ASID002", ReceptionDate: "03-02-2020", DispatchDate: "04-02-2020", ProcessingMethod: "XYZ", CountryofProcessing: "China", ProductForm: "Canned", ProductCondition: "Frozen", SizeGrade: "10", SellbyDate: "15-03-2020", ProcessorName: "XProcessors"},
		ProcessorTransaction{TxID: "ASID003Tx02", AssetID: "ASID003", ReceptionDate: "03-02-2020", DispatchDate: "04-02-2020", ProcessingMethod: "XYZ", CountryofProcessing: "China", ProductForm: "Canned", ProductCondition: "Frozen", SizeGrade: "10", SellbyDate: "15-03-2020", ProcessorName: "XProcessors"},
		ProcessorTransaction{TxID: "ASID004Tx02", AssetID: "ASID004", ReceptionDate: "03-02-2020", DispatchDate: "04-02-2020", ProcessingMethod: "XYZ", CountryofProcessing: "China", ProductForm: "Canned", ProductCondition: "Frozen", SizeGrade: "10", SellbyDate: "15-03-2020", ProcessorName: "XProcessors"},
		ProcessorTransaction{TxID: "ASID005Tx02", AssetID: "ASID005", ReceptionDate: "03-02-2020", DispatchDate: "04-02-2020", ProcessingMethod: "XYZ", CountryofProcessing: "China", ProductForm: "Canned", ProductCondition: "Frozen", SizeGrade: "10", SellbyDate: "15-03-2020", ProcessorName: "XProcessors"},
	}
	distributortx := []DistributorTransaction{
		DistributorTransaction{TxID: "ASID001Tx03", AssetID: "ASID001", ReceptionDate: "06-02-2020", DispatchDate: "08-02-2020", ColdStorageConditions: "-1 Deg", DistributorName: "XDistributors"},
		DistributorTransaction{TxID: "ASID002Tx03", AssetID: "ASID002", ReceptionDate: "06-02-2020", DispatchDate: "08-02-2020", ColdStorageConditions: "-1 Deg", DistributorName: "XDistributors"},
		DistributorTransaction{TxID: "ASID003Tx03", AssetID: "ASID003", ReceptionDate: "06-02-2020", DispatchDate: "08-02-2020", ColdStorageConditions: "-1 Deg", DistributorName: "XDistributors"},
		DistributorTransaction{TxID: "ASID004Tx03", AssetID: "ASID004", ReceptionDate: "06-02-2020", DispatchDate: "08-02-2020", ColdStorageConditions: "-1 Deg", DistributorName: "XDistributors"},
		DistributorTransaction{TxID: "ASID005Tx03", AssetID: "ASID005", ReceptionDate: "06-02-2020", DispatchDate: "08-02-2020", ColdStorageConditions: "-1 Deg", DistributorName: "XDistributors"},
	}
	grocertx := []GrocerTransaction{
		GrocerTransaction{TxID: "ASID001Tx04", AssetID: "ASID001", ReceptionDate: "10-02-2020", GrocerName: "XGrocers"},
		GrocerTransaction{TxID: "ASID002Tx04", AssetID: "ASID002", ReceptionDate: "10-02-2020", GrocerName: "XGrocers"},
		GrocerTransaction{TxID: "ASID003Tx04", AssetID: "ASID003", ReceptionDate: "10-02-2020", GrocerName: "XGrocers"},
		GrocerTransaction{TxID: "ASID004Tx04", AssetID: "ASID004", ReceptionDate: "10-02-2020", GrocerName: "XGrocers"},
		GrocerTransaction{TxID: "ASID005Tx04", AssetID: "ASID005", ReceptionDate: "10-02-2020", GrocerName: "XGrocers"},
	}

	i := 0
	for i < len(producertx) {
		fmt.Println("i is ", i)
		producertxAsBytes, _ := json.Marshal(producertx[i])
		APIstub.PutState(producertx[i].TxID, producertxAsBytes)
		fmt.Println("Added", producertx[i])
		i = i + 1
	}

	j := 0
	for j < len(processortx) {
		fmt.Println("j is ", j)
		processortxAsBytes, _ := json.Marshal(processortx[j])
		APIstub.PutState(processortx[j].TxID, processortxAsBytes)
		fmt.Println("Added", processortx[j])
		j = j + 1
	}

	k := 0
	for k < len(distributortx) {
		fmt.Println("k is ", k)
		distributortxAsBytes, _ := json.Marshal(distributortx[k])
		APIstub.PutState(distributortx[k].TxID, distributortxAsBytes)
		fmt.Println("Added", distributortx[k])
		k = k + 1
	}

	l := 0
	for l < len(grocertx) {
		fmt.Println("l is ", l)
		grocertxAsBytes, _ := json.Marshal(grocertx[l])
		APIstub.PutState(grocertx[l].TxID, grocertxAsBytes)
		fmt.Println("Added", grocertx[l])
		l = l + 1
	}

	return shim.Success(nil)
}

/*
 * The recordProducerTransaction method *
 */
func (s *SmartContract) recordProducerTransaction(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 11 {
		return shim.Error("Incorrect number of arguments. Expecting 11")
	}

	var producertx = ProducerTransaction{TxID: args[0], AssetID: args[1], TotalWeightShippedPerLot: args[2], BoxesShippedPerLot: args[3], DateofProduction: args[4], DateofHarvest: args[5], NameoftheFarm: args[6], ProductionMethod: args[7], FarmingCountry: args[8], SpeciesName: args[9], DispatchDate: args[10]}

	producertxAsBytes, _ := json.Marshal(producertx)
	err := APIstub.PutState(args[0], producertxAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record producertx catch: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * The recordProcessorTransaction method *
 */
func (s *SmartContract) recordProcessorTransaction(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 11 {
		return shim.Error("Incorrect number of arguments. Expecting 11")
	}

	var processortx = ProcessorTransaction{TxID: args[0], AssetID: args[1], ReceptionDate: args[2], DispatchDate: args[3], ProcessingMethod: args[4], CountryofProcessing: args[5], ProductForm: args[6], ProductCondition: args[7], SizeGrade: args[8], SellbyDate: args[9], ProcessorName: args[10]}

	processortxAsBytes, _ := json.Marshal(processortx)
	err := APIstub.PutState(args[0], processortxAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record processortx catch: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * The recordDistributorTransaction method *
 */
func (s *SmartContract) recordDistributorTransaction(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	var distributortx = DistributorTransaction{TxID: args[0], AssetID: args[1], ReceptionDate: args[2], DispatchDate: args[3], ColdStorageConditions: args[4], DistributorName: args[5]}

	distributortxAsBytes, _ := json.Marshal(distributortx)
	err := APIstub.PutState(args[0], distributortxAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record distributortx catch: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * The recordGrocerTransaction method *
 */
func (s *SmartContract) recordGrocerTransaction(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	var grocertx = GrocerTransaction{TxID: args[0], AssetID: args[1], ReceptionDate: args[2], GrocerName: args[3]}

	grocertxAsBytes, _ := json.Marshal(grocertx)
	err := APIstub.PutState(args[0], grocertxAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record grocertx catch: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * The queryAllTuna method *
 */
func (s *SmartContract) queryTransactionHistory(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	startKey := args[0]
	endKey := args[1]

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add comma before array members,suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true

	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllAssetTransactions:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

/*
 * The query method *
 */
func (s *SmartContract) queryTransaction(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	assetAsBytes, _ := APIstub.GetState(args[0])
	if assetAsBytes == nil {
		return shim.Error("Could not locate asset")
	}
	return shim.Success(assetAsBytes)
}

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
