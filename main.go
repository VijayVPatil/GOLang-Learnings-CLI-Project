package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// these variable are declared outside all the functions or top of all functions
// this is because it can be accessed by any function
// these are called PACKAGE
//
//	NOTE if anywhere in the function declaration of calling you dont see any parameters require
//
// see once at below package variables they might be declared here so they are not used in the function
// calling
var conferenceName string = "Go Conference"

const totalTickets int = 50

var remainingTicktes uint = 50

// This is how we create a slice
// var bookings []string

// This is how we create a empty slice of type map
// create empty list of maps
// var bookings = make([]map[string]string, 0)

// This is how we write with user data struct type
var bookings = make([]UserData, 0)

type UserData struct { //struct has mix data type
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	// var conferenceName string = "Go Conference" // if the reference var. is declared as var it can be updated anywhere in the program
	// const totalTickets int = 50                 // if the reference var. is declared as const it can not be updated anywhere in the program
	// var remainingTicktes uint = 50

	// var bookings [50]string // This is how we create an array
	// var bookings []string //This is how we create slice
	// var bookings = []string{} // This is alternate way how we create slice

	//var name string="Vijay"
	// name := " Vijay" //instead of using what we wrote above we can write same what we wrote this line
	// fmt.Printf("The name is %T type\n ", name)  //by using %T placeholder we can know the type of name reference variabll

	// fmt.Printf("Conference Ticket is %T ,remainingTickets is %T,conferenceName is %T\n", totalTickets, remainingTicktes, conferenceName)
	// greetUsers(conferenceName, totalTickets, int(remainingTicktes)) //In place of this as the parameters are declared a package var we can remove it as shown below
	greetUsers()

	// for remainingTicktes > 0 && len(bookings) < 50 {

	//following is how we declare multiple ref var when we get multiple return
	//from the function
	firstName, lastName, emailId, userTickets := getUserInput()

	//following is how we how we store multiple return values.
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, emailId, userTickets, remainingTicktes)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, emailId)
		// fmt.Printf("The whole slice: %v\n", bookings)   //printing array
		// fmt.Printf("The first value:%v\n", bookings[0]) // printing value of first element of array

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, emailId)
		//below will return type of bookings
		// fmt.Printf("The type of bookings slice is %T \n", bookings)

		//below will return length of array
		// fmt.Printf("slice length : %v \n", len(bookings))

		//This Function returns first Names of users
		//so we return firstNames right from the function but we need to store it hencce we
		//declare a ref var like firstNames.
		// firstNames := getprintFirstNames(bookings)   //In place of this as the parameters are declared a package var we can remove it as shown below
		firstNames := getprintFirstNames()
		fmt.Printf("First names of bookings are  %v\n", firstNames)

		if remainingTicktes == 0 {
			//end the program
			fmt.Println("Our conference is booked come back next year")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("First name or Last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets entered is too short")
		}
		fmt.Printf("Your input data is invalid try again \n")
	}

	wg.Wait() //waitgroup function that blocks the wait group counter is 0

	//Following is assiging a data type explicitly
	// var userName string
	// var userTickets int
	// userName = "Tom"
	// userTickets = 2
	// fmt.Printf("User %v booked %v tickets. \n", userName, userTickets)
	//-----------------------------------------------
	// }
}

func greetUsers() {
	fmt.Println("Welcome to our ", conferenceName, ".")

	fmt.Printf("We have a total of %v and remaining tickets are %v.\n", totalTickets, remainingTicktes)
	fmt.Println("Book your t ickets now")
}

func getprintFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings { // see here we have used _(underscore) called blank identifier , if we dont want to use any ref. var. in program we can use _
		// var names = strings.Fields(booking) //by performing strings.Field on booking ref. var. we get an o/p as slice stored in the ref. var. called names
		// var firstName = names[0]

		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames // this is how we return a value in the function

}

// Following line tells how functions accept parameters and how we can state the type of return values like in bracket(bool,bool,bool)
func validateUserInput(firstName string, lastName string, emailId string, userTickets uint, remainingTicktes uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(emailId, "@")
	var isValidTicketNumber bool = userTickets > 0 && userTickets <= remainingTicktes

	//following is how we return multiple values
	return isValidEmail, isValidName, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var emailId string
	var userTickets uint

	//taking user input
	fmt.Println("Enter the your first name")
	fmt.Scan(&firstName) //here the scan function will take user input scan it assign it the
	// the value to the user variable because it has pointer toits address
	fmt.Println("Enter the your last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter the your emailId")
	fmt.Scan(&emailId)
	fmt.Println("Enter your required tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, emailId, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, emailId string) {
	remainingTicktes = remainingTicktes - userTickets

	//create a empty map
	// var userData = make(map[string]string)
	//adding key value pair for a map
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["emailId"] = emailId
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           emailId,
		numberOfTickets: userTickets,
	}

	//bookings[0] = firstName + " " + lastName //This is how we add value in array
	// bookings = append(bookings, firstName+" "+lastName) //appends the value at last position of array and returns the updated slice

	bookings = append(bookings, userData)
	fmt.Printf("List Of bookings is %v \n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, emailId)
	fmt.Printf("%v tickets remaining for %v\n", remainingTicktes, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("-------------------------")
	fmt.Printf("Sending ticket %v to email address %v \n", ticket, email)
	fmt.Println("-----------------")

	wg.Done() // removes the thread that we added from the waiting list
}
