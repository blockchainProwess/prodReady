
// Package helloworld displays helloworld and the amount of time the invoke method as been called
package main

import (
	// "errors"
	// "fmt"
	// "strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// var myLogger = logging.MustGetLogger("asset_mgm")

// HelloWorld example simple Chaincode implementation
type HelloWorld struct {
}

// Init is called during Deploy transaction after the container has been
func (t *HelloWorld) Init(shim.ChaincodeStubInterface) peer.Response {
	// myLogger.Debug("HelloWorld Init is called!")
	return shim.Success(nil)
}

// Invoke is called for every Invoke transactions. The chaincode may change
// its state variables
func (t *HelloWorld) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Error("Hello World!")
}

// Query is called for Query transactions. The chaincode may only read
// (but not modify) its state variables and return the result
func (t *HelloWorld) query(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// myLogger.Debug("HelloWorld Query is called!\n")
	return shim.Error("Hello World from query section!")
}


func main() {
		
	shim.Start(new(HelloWorld))
}
	
