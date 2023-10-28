package controllers

import (
	"encoding/json"
	"fmt"
	"golang-json-server/models"
	"os"

	"github.com/gin-gonic/gin"
)

var Accounts []models.Account
var List []models.GirlFriend
var Booking []models.Booking
var Loaded = false

func LoadDatabase() models.Database {
	db := models.ConnectDatabase();

	// Load Accounts
	for _, account := range db.Accounts {
		var newAccount models.Account;

		newAccount.Id = account.Id;
		newAccount.Username = account.Username;
		newAccount.Password = account.Password;

		Accounts = append(Accounts, newAccount);
	}

	// Load List
	for _, girlfriend := range db.List {
		var newGirlFriend models.GirlFriend;

		newGirlFriend.Id = girlfriend.Id;
		newGirlFriend.Name = girlfriend.Name;
		newGirlFriend.Age = girlfriend.Age;
		newGirlFriend.Rate = girlfriend.Rate;
		newGirlFriend.Description = girlfriend.Description;

		List = append(List, newGirlFriend);
	}

	// Load booking
	for _, booking := range db.Booking {
		var newBooking models.Booking;

		newBooking.Id = booking.Id;
		newBooking.GirlFriendId = booking.GirlFriendId;
		newBooking.CustomerId = booking.CustomerId;

		Booking = append(Booking, newBooking);
	}

	Loaded = true;

	return db;
}

func GetAllGirlFriends(ctx *gin.Context) {
	// Load data from database
	if Loaded != true {
		LoadDatabase();		
	}

	ctx.JSON(200, gin.H{
		"message": "Get data successfully",
		"data": List,
	})
}

func AddNewBooking(ctx *gin.Context) {
	// Load data from database
	if Loaded != true {
		LoadDatabase();		
	}
	
	newBooking := models.Booking{};
	userBooking := models.Booking{};

    // Call BindJSON to bind the received JSON to Booking Struct.
    if err := ctx.BindJSON(&newBooking); err != nil {
        return;
    }

	userBooking.Id = newBooking.Id;
	userBooking.GirlFriendId = newBooking.GirlFriendId;
	userBooking.CustomerId = newBooking.CustomerId;

	if userBooking != (models.Booking{}) {
		fmt.Println(userBooking.CustomerId);
		fmt.Println(userBooking.GirlFriendId);
		fmt.Println(userBooking.CustomerId);
		
		// Add to Booking list
		Booking = append(Booking, userBooking);
	}

	// Create data model
	data := models.Database {
		Accounts: Accounts,
		List: List,
		Booking: Booking,
	}
 
	// Convert struct to json
	content, err := json.Marshal(data);

	// Convert failed
	if err != nil {
		panic(err);
	}

	// Write to database
	err = os.WriteFile("db.json", content, 0644);

	// Write file failed
	if err != nil {
		panic(err);
	}

	ctx.JSON(200, gin.H{
		"message": "Add booking successfully",
		"booking_info": userBooking,
	})
}