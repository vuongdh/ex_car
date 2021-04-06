package main
import (
 "encoding/json"
"github.com/hyperledger/fabric/core/chaincode/shim"
 pb "github.com/hyperledger/fabric/protos/peer"
)
type Car struct {
 modelName    string
 color        string
 serialNo     string
 manufacturer string
 owner        Owner
}
type Owner struct {
 name            string
 nationaIdentity string
 gender          string
 address         string
}
func (c *Car) changeOwner(newOwner Owner) {
 c.owner = newOwner
}
type CarChaincode struct {
}
func (c *CarChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
//Declare owners
 tom := Owner{name: "Tom H", nationaIdentity: "ABCD33457", gender: "M", address: "1, Tumbbad"}
 bob := Owner{name: "Bob M", nationaIdentity: "QWER33457", gender: "M", address: "2, Tumbbad"}
//Decale Car
 car := Car{modelName: "Land Rover", color: "white", serialNo: "334712531234", manufacturer: "Tata Motors", owner: tom}
// convert tom Owner to []byte
 tomAsJSONBytes, _ := json.Marshal(tom)
 //Add Tom to ledger
 err := stub.PutState(tom.nationaIdentity, tomAsJSONBytes)
 if err != nil {
  return shim.Error("Failed to create asset " + tom.name)
 }
//Add Bob to ledger
 bobAsJSONBytes, _ := json.Marshal(bob)
 err = stub.PutState(bob.nationaIdentity, bobAsJSONBytes)
 if err != nil {
  return shim.Error("Failed to create asset " + bob.name)
 }
//Add car to ledger
 carAsJSONBytes, _ := json.Marshal(car)
 err = stub.PutState(car.serialNo, carAsJSONBytes)
 if err != nil {
  return shim.Error("Failed to create asset " + car.serialNo)
 }
 return shim.Success([]byte("Assets created successfully."))
}

func (c *CarChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
// Read args from the transaction proposal.
// fc=> method to invoke
 fc, args := stub.GetFunctionAndParameters()
 if fc == "TransferOwnership" {
  return c.TransferOwnership(stub, args)
 } else if fc == "queryAllCars" {
	return s.queryAllCars(stub)
 }
 return shim.Error("Called function is not defined in the chaincode ")

 // Retrieve the requested Smart Contract function and arguments
 function, args := APIstub.GetFunctionAndParameters()
 // Route to the appropriate handler function to interact with the ledger appropriately
 /*
 if function == "queryCar" {
	 return s.queryCar(APIstub, args)
 } else if function == "initLedger" {
	 return s.initLedger(APIstub)
 } else if function == "createCar" {
	 return s.createCar(APIstub, args)
 } else if function == "queryAllCars" {
	 return s.queryAllCars(APIstub)
 } else if function == "changeCarOwner" {
	 return s.changeCarOwner(APIstub, args)
 }

 return shim.Error("Invalid Smart Contract function name.")
 */
}
func (s *SmartContract) queryAllCars(stub shim.ChaincodeStubInterface) pd.Response {

	resultsIterator, err := stub.GetStateByRange("", "")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var cars []*Car 
	for resultsIterator.HasNext(){
		queryResponse, err := resultsIterator.Next()
		if err != nill {
			return nil, err
		}

		var car Car 
		err = json.Unmarshal(queryResponse.Value, &car)
		if (err != nil) {
			return nil, err 
		}

		cars = append(cars, &car)
	} 
	return cars, nil
}

func (c *CarChaincode) TransferOwnership(stub shim.ChaincodeStubInterface, args []string) pb.Response {
 // args[0]=> car serial no
 // args[1]==> new owner national identity
 // Read car asset
 carAsBytes, _ := stub.GetState(args[0])
 if carAsBytes == nil {
  return shim.Error("car asset not found")
 }
 car := Car{}
 _ = json.Unmarshal(carAsBytes, &car)
//Read newOwnerDetails
 ownerAsBytes, _ := stub.GetState(args[1])
 if ownerAsBytes == nil {
  return shim.Error("owner asset not found")
 }
 newOwner := Owner{}
 _ = json.Unmarshal(ownerAsBytes, &newOwner)
car.changeOwner(newOwner)
 carAsJSONBytes, _ := json.Marshal(car)
 err := stub.PutState(car.serialNo, carAsJSONBytes)
 if err != nil {
  return shim.Error("Failed to create asset " + car.serialNo)
 }
return shim.Success([]byte("Asset modified."))
}

func main() {
logger.SetLevel(shim.LogInfo)
// Start the chaincode process
err := shim.Start(new(CarChaincode))
if err != nil {
logger.Error("Error starting PhantomChaincode - ", err.Error()
}
}