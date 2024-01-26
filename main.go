package main

import (
	"fmt"
	"sync"
	"time"
	// "strconv"
)

type UserData struct{
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var conferenceName = "Linkin Park Tribute"
const conferenceTickets=100
var remainingTickets uint =50
// var bookings=make([]map[string]string,0)
var bookings=make([]UserData,0)

var wg = sync.WaitGroup{}

func main(){


	greetUsers()

	for{
		firstName,lastName,userEmail,userTickets:=getUserInput()
		isValidName,isValidEmail,isValidTicketNumber:= validateUserInput(firstName,lastName,userEmail,userTickets,remainingTickets)

		if isValidName&& isValidEmail && isValidTicketNumber{
		fmt.Printf("Hi! %v %v, you have bought %v tickets. You will get a confirmation email on %v shortly.\n",firstName,lastName,userTickets,userEmail)

		remainingTickets=remainingTickets-userTickets

		var userData=UserData{
			firstName:firstName,
			lastName: lastName,
			email: userEmail,
			numberOfTickets: userTickets,
		}

		// var userData=make(map[string]string)
		// userData["firstName"]=firstName
		// userData["lastName"]=lastName
		// userData["userEmail"]=userEmail
		// userData["numberOfTickets"]=strconv.FormatUint(uint64(userTickets),10)

		bookings=append(bookings,userData)
		fmt.Printf("The list of bookings are %v",bookings)

		fmt.Printf("%v tickets are remaining for %v\n",remainingTickets,conferenceName)
		wg.Add(1)
		go sendTickets(userTickets,firstName,lastName,userEmail)

		firstNames:=getFirstNames()
		fmt.Printf("The first names of bookings are %v\n",firstNames)

		if remainingTickets ==0{
			fmt.Println("We are completely booked out, Please visit us next year! ")
			break
		}
		}else{
			if !isValidName{
				fmt.Println("Sorry, the First name or the Last name that you entered is incorrect")
			}
			if !isValidEmail{
				fmt.Println("Sorry, the email that you have enetered is incorrect")
			}
			if !isValidTicketNumber{
				fmt.Println("Sorry, the the number of tickets that you have entered is invalid")
			}

		}
		
}
wg.Wait()
}

func greetUsers(){
	fmt.Printf("Hi there, Grab your tickets for the %v that is scheduled for next week!\n",conferenceName)
	fmt.Printf("Please note that all the tickets will be sold here ONLY! We have a total of %v tickets and the total remaining tickets are %v tickets \n",conferenceTickets,remainingTickets)
}

func getFirstNames()[]string{
	firstNames:=[]string{}
	for _,booking := range bookings{
		// firstNames=append(firstNames, booking["firstName"])
		firstNames=append(firstNames,booking.firstName)
	}
	return firstNames
}

func getUserInput()(string,string,string,uint){
	var firstName string
	var lastName string
	var userEmail string
	var userTickets uint

	fmt.Println("Please enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address")
	fmt.Scan(&userEmail)

	fmt.Println("Please enter the no. of tickets you want to book")
	fmt.Scan(&userTickets)

	return firstName,lastName,userEmail,userTickets
}

func sendTickets(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(time.Second*10)
	var ticket=fmt.Sprintf("%v tickets for %v %v",userTickets,firstName,lastName)
	fmt.Println("####################################")
	fmt.Printf("Sending Tickets:\n%v \nto email address %v\n",ticket,email)
	fmt.Println("####################################")
	wg.Done()
}