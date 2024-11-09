package database
//This declares that the code belongs to the database package. This package will handle the connection to the database (in this case, MongoDB).


import(
	"context"
	//helps manage timeouts and cancellations for processes. When connecting to a database, we might want to limit the time the program spends waiting for the connection.

	"fmt"
	//fmt package provides formatting functions like Println or Printf for printing output to the console.

	"log"
	//is used for logging messages. If there is an issue while connecting to the database, you would use log to record the error.

	"time"
	//used for time-related functions, such as adding a timeout for the database connection.

	"go.mongodb.org/mongo-driver/mongo"
	// package provides the MongoDB Go Driver, which allows your application to interact with a MongoDB database.

	"go.mongodb.org/mongo-driver/mongo/options"
	//package helps configure different options, such as database connection settings.
)


//This function will set up a connection to MongoDB. It's a function that returns a *mongo.Client, which is a MongoDB client used to interact with the database.
//*mongo.Client: This is a pointer to a mongo.Client type. It is returned once the database connection is established, allowing you to use it in other parts of your application to read or write data.

func DBinstance() *mongo.Client{
    MongoDB :="mongodb://localhost:27017"
	//telling your program where to find the MongoDB database. In simple terms, this is like giving the address of the database.
    //"mongodb://localhost:27017": This says that the MongoDB database is running on your own computer (localhost) at the default port 27017.
    //MongoDB is a variable that stores this address, like saving someone's home address in your phone contacts.
	
	fmt.Print(MongoDB)
	//print out the MongoDB address on your screen or terminal, so you can see the address you just defined. 
	//It’s useful for debugging or checking if the address is correct.

	client, err:= mongo.NewClient(options.Client().ApplyURI(MongoDB))
	// creating a MongoDB client. Think of the client as a tool that helps you talk to the MongoDB database.
	//mongo.NewClient: This function creates a new client.
    //ApplyURI(MongoDB): This tells the client where the database is located (in this case, it's at "mongodb://localhost:27017").
    //If something goes wrong while creating this client (like if the address is wrong), an error is stored in the variable err.


	if err!=nil{
		log.Fatal(err)
	}
	//If there’s an error, this will print the error message and stop the program. It’s like saying, "If there's a problem, let me know and stop everything!"

	ctx , cancel:=context.WithTimeout(context.Background(), 10*time.Second)
	//ctx and cancel are used to control how long the program waits when trying to connect to MongoDB.
    //Try to connect to MongoDB, but if it takes longer than 10 seconds, stop and give up."This prevents the program from getting stuck if the database is unreachable.
	//context.WithTimeout: This creates a "timeout" of 10 seconds, meaning the program won’t wait longer than that to connect to the database.


	defer cancel()
	//defer is a special command that says "run this at the end of the function."
    //cancel(): This cleans up the context after the function is done. It’s like saying, "Okay, we’re done trying to connect, so let’s stop waiting."

	err=client.Connect(ctx)
    //client.Connect(ctx): After creating the client, you need to connect it to the MongoDB database. This line tries to establish that connection using the context (ctx) you created earlier (with the 10-second timeout).


	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("connected to mongodb")
	//If the connection is successful, this line prints a message saying "connected to mongodb". This lets you know that your program has successfully connected to the MongoDB database.
    //this is like the else block

	return client
	//This line returns the connected MongoDB client so that other parts of the program can use it to interact with the database.


}
//4. Storing the Client Globally
var Client *mongo.Client=DBinstance()
//This creates a global variable Client and assigns it the result of the DBinstance() function.


func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection{
	var collection *mongo.Collection= client.Database("restaurant").Collection(collectionName)
	return collection
}

//OpenColection: This function helps you get a specific collection from the MongoDB database. In MongoDB, a collection is like a table in a relational database. Each collection stores a group of similar data (e.g., foods, orders, invoices, etc.).
//Parameters:
//client *mongo.Client: This is the connected MongoDB client (that was returned from DBinstance).
//collectionName string: This is the name of the collection you want to access (e.g., "foods", "orders").



//If you wanted to interact with the "foods" collection, you could do something like this:
//foodCollection := OpenColection(Client, "foods")

