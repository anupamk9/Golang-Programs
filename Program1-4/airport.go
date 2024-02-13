/*Program: Airport Management System
  
  Course: MCA372D Go Programming Lab 

  Description: This program manages an airport with flights, planes and passengers. 
  It includes functionalities to add a new flight, view all flights, search for a flight, 
  update departure time, delete a flight and display all flights after performing CRUD operations.

  Author: Anupam Kumar 2347104
 */

package main

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

// Constants
const (
	menuHeader         = "\n===== Airport Management System ====="
	menuOptions        = "1. Add New Flight\n2. View All Flights\n3. Search for a Flight\n" +
		"4. Update Departure Time\n5. Delete a Flight\n6. Display All Flights After Operations\n7. Exit"
	invalidChoiceMsg   = "Invalid choice. Please enter a number between 1 and 7."
	noFlightsAvailable = "No flights available."
)

// Struct to represent information about a flight
type Flight struct {
	FlightNumber string
	Destination  string
	Departure    string
}

// Slice to store the list of flights
var flights []Flight

// Function to display the main menu
func displayMenu() {
	fmt.Println(menuHeader)
	fmt.Println(menuOptions)
	fmt.Print("Enter your choice (1-7): ")
}

// Function to add a new flight to the system
func addFlight() {
	fmt.Println("\n===== Add New Flight =====")
	var newFlight Flight

	// Get flight details from the user
	fmt.Print("Enter Flight Number (alphanumeric): ")
	fmt.Scanln(&newFlight.FlightNumber)
	if !isValidFlightNumber(newFlight.FlightNumber) {
		fmt.Println("Invalid Flight Number format. Please use alphanumeric characters.")
		return
	}

	// Validate if the flight number already exists
	if isFlightNumberExists(newFlight.FlightNumber) {
		fmt.Println("Flight with the same number already exists. Please choose a different number.")
		return
	}

	fmt.Print("Enter Destination: ")
	fmt.Scanln(&newFlight.Destination)

	fmt.Print("Enter Departure Time (HH:mm): ")
	fmt.Scanln(&newFlight.Departure)
	if !isValidTimeFormat(newFlight.Departure) {
		fmt.Println("Invalid time format. Please use the format HH:mm.")
		return
	}

	// Add the new flight to the flights slice
	flights = append(flights, newFlight)

	fmt.Println("Flight added successfully!")
}

// Function to view the list of flights
func viewFlights() {
	fmt.Println("\n===== View All Flights =====")
	if len(flights) == 0 {
		fmt.Println(noFlightsAvailable)
		return
	}

	// Display flights in a formatted table
	fmt.Printf("%-15s %-20s %-15s\n", "Flight Number", "Destination", "Departure Time")
	fmt.Println("-----------------------------------------------------")
	for _, flight := range flights {
		fmt.Printf("%-15s %-20s %-15s\n", flight.FlightNumber, flight.Destination, flight.Departure)
	}
}

// Function to search for a flight based on the flight number
func searchFlight() {
	fmt.Println("\n===== Search for a Flight =====")
	if len(flights) == 0 {
		fmt.Println(noFlightsAvailable)
		return
	}

	fmt.Print("Enter Flight Number to search: ")
	var searchFlightNumber string
	fmt.Scanln(&searchFlightNumber)

	// Iterate through flights to find the matching flight
	for _, flight := range flights {
		if flight.FlightNumber == searchFlightNumber {
			fmt.Printf("\nFlight Found!\n%-15s %-20s %-15s\n", "Flight Number", "Destination", "Departure Time")
			fmt.Println("-----------------------------------------------------")
			fmt.Printf("%-15s %-20s %-15s\n", flight.FlightNumber, flight.Destination, flight.Departure)
			return
		}
	}

	// Flight not found
	fmt.Println("Flight not found for the given Flight Number.")
}

// Function to update the departure time of a flight
func updateFlightDepartureTime() {
	fmt.Println("\n===== Update Departure Time =====")
	if len(flights) == 0 {
		fmt.Println(noFlightsAvailable)
		return
	}

	fmt.Print("Enter Flight Number to update departure time: ")
	var updateFlightNumber string
	fmt.Scanln(&updateFlightNumber)

	// Iterate through flights to find the matching flight
	for i, flight := range flights {
		if flight.FlightNumber == updateFlightNumber {
			fmt.Printf("Current Departure Time for Flight %s: %s\n", flight.FlightNumber, flight.Departure)
			fmt.Print("Enter new Departure Time (HH:mm): ")
			var newDepartureTime string
			fmt.Scanln(&newDepartureTime)
			if !isValidTimeFormat(newDepartureTime) {
				fmt.Println("Invalid time format. Please use the format HH:mm.")
				return
			}

			// Update the departure time of the flight
			flights[i].Departure = newDepartureTime
			fmt.Println("Departure Time updated successfully.")
			return
		}
	}

	// Flight not found
	fmt.Println("Flight not found for the given Flight Number.")
}

// Function to delete a flight from the system
func deleteFlight() {
	fmt.Println("\n===== Delete a Flight =====")
	if len(flights) == 0 {
		fmt.Println(noFlightsAvailable)
		return
	}

	fmt.Print("Enter Flight Number to delete: ")
	var deleteFlightNumber string
	fmt.Scanln(&deleteFlightNumber)

	// Iterate through flights to find the matching flight
	for i, flight := range flights {
		if flight.FlightNumber == deleteFlightNumber {
			// Remove the flight from the flights slice
			flights = append(flights[:i], flights[i+1:]...)
			fmt.Println("Flight deleted successfully.")
			return
		}
	}

	// Flight not found
	fmt.Println("Flight not found for the given Flight Number.")
}

// Function to check if a flight number already exists in the system
func isFlightNumberExists(flightNumber string) bool {
	for _, flight := range flights {
		if flight.FlightNumber == flightNumber {
			return true
		}
	}
	return false
}

// Function to display all flights after performing CRUD operations
func displayAllFlights() {
	fmt.Println("\n===== All Flights After Operations =====")
	if len(flights) == 0 {
		fmt.Println(noFlightsAvailable)
		return
	}

	// Display flights in a formatted table
	fmt.Printf("%-15s %-20s %-15s\n", "Flight Number", "Destination", "Departure Time")
	fmt.Println("-----------------------------------------------------")
	for _, flight := range flights {
		fmt.Printf("%-15s %-20s %-15s\n", flight.FlightNumber, flight.Destination, flight.Departure)
	}
}

// Function to validate if a flight number is in alphanumeric format
func isValidFlightNumber(flightNumber string) bool {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9]+$", flightNumber)
	return matched
}

// Function to validate if the time format is in HH:mm
func isValidTimeFormat(timeStr string) bool {
	_, err := time.Parse("15:04", timeStr)
	return err == nil
}

func main() {
	fmt.Println("Welcome to the Airport Management System!")

	for {
		// Display the main menu
		displayMenu()

		// Get user choice
		var choice int
		fmt.Scanln(&choice)

		// Check if there are no flights before any operation
		if len(flights) == 0 && choice != 1 && choice != 7 {
			fmt.Println(noFlightsAvailable + " Please add a new flight first.")
			continue
		}

		// Process user choice
		switch choice {
		case 1:
			addFlight()
		case 2:
			viewFlights()
		case 3:
			searchFlight()
		case 4:
			updateFlightDepartureTime()
		case 5:
			deleteFlight()
		case 6:
			displayAllFlights()
		case 7:
			fmt.Println("Exiting the Airport Management System. Goodbye!")
			os.Exit(0)
		default:
			fmt.Println(invalidChoiceMsg)
		}
	}
}
