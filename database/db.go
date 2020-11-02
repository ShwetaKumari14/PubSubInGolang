package database

import (
	"GoAssignment/utils"
	"database/sql"
	"encoding/json"
)

func PerformDBAction(message map[string][]map[string]interface{}) (string, error){

	database, err := utils.GetDBConnection()
	if err != nil {
		utils.GenerateLogs(err, "Error in making connection to sqlite")
		return "failure", err
	}
	defer database.Close()

	statement1, err := database.Prepare("CREATE TABLE IF NOT EXISTS hotel (hotel_id TEXT PRIMARY KEY, name TEXT, country TEXT, address TEXT, latitude TEXT, longitude TEXT, telephone TEXT, amenities TEXT, description TEXT, room_count TEXT, currency TEXT)")
	if err != nil {
		utils.GenerateLogs(err, "Error In Creating Hotel Table")
		return "failure", err
	}
	statement1.Exec()

	statement2, err := database.Prepare("CREATE TABLE IF NOT EXISTS room (hotel_id  TEXT, room_id TEXT, description TEXT, name TEXT, capacity TEXT)")
	if err != nil {
		utils.GenerateLogs(err, "Error In Creating Room Table")
		return "failure", err
	}
	statement2.Exec()

	statement3, err := database.Prepare("CREATE TABLE IF NOT EXISTS rate_plan (hotel_id TEXT, rate_plan_id TEXT, cancellation_policy TEXT, name TEXT, other_conditions TEXT, meal_plan TEXT)")
	if err != nil {
		utils.GenerateLogs(err, "Error In Creating Rate_plan Table")
		return "failure", err
	}
	statement3.Exec()

	status1, err := InsertIntoHotelTable(statement1, database, message)
	status2, err := InsertIntoRoomTable(statement2, database, message)
	status3, err := InsertIntoRatePlanlTable(statement3, database, message)

	if status1 == "success" && status2 == "success" && status3 == "success" {
		return "success", nil
	}
	
	return status1 + status2 + status3, err

	return "sdsd", nil
}

func InsertIntoHotelTable(statement *sql.Stmt, database *sql.DB, message map[string][]map[string]interface{})(string, error){

	var err error
	var hotel_id, name, country, address, telephone, description, currency, amenities string
	var latitude, longitude, room_count float64

	for _, value := range message {
		for key,val := range value[0]["hotel"].(map[string]interface{}) {
			if key == "hotel_id"{
				hotel_id = val.(string)
			}
			if key == "name"{
				name = val.(string)
			}
			if key == "country"{
				country = val.(string)
			}
			if key == "address"{
				address = val.(string)
			}
			if key == "latitude"{
				latitude = val.(float64)
			}
			if key == "longitude"{
				longitude = val.(float64)
			}
			if key == "description"{
				description = val.(string)
			}
			if key == "room_count"{
				room_count = val.(float64)
			}
			if key == "currency"{
				currency = val.(string)
			}
			if key == "amenities"{
				resBd, _ := json.Marshal(val)
				amenities =  string(resBd)
			}
		}
	}

	statement, err = database.Prepare("INSERT INTO hotel (hotel_id , name , country , address , latitude , longitude , telephone , amenities , description , room_count , currency ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		return "Error In Prepare Statement Of Insert Query :", err
	}
	defer statement.Close()

	_, err = statement.Exec(hotel_id, name, country, address, latitude, longitude, telephone, amenities, description, room_count, currency)
	if err != nil {
		return "Error In Inserting Record Into Table :", err
	}

	return "success", nil
}

func InsertIntoRoomTable(statement *sql.Stmt, database *sql.DB, message map[string][]map[string]interface{})(string, error){

	var err error
	var hotel_id, name, room_id, description, capacity string

	for _, value := range message {
		for key,val := range value[0]["room"].(map[string]interface{}) {
			if key == "hotel_id"{
				hotel_id = val.(string)
			}
			if key == "name"{
				name = val.(string)
			}
			if key == "room_id"{
				room_id = val.(string)
			}
			if key == "description"{
				description = val.(string)
			}
			
			if key == "capacity"{
				resBd, _ := json.Marshal(val)
				capacity =  string(resBd)
			}
		}
	}

	statement, err = database.Prepare("INSERT INTO room (hotel_id , room_id , description , name , capacity) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		return "Error In Prepare Statement Of Insert Query :", err
	}
	defer statement.Close()

	_, err = statement.Exec(hotel_id, room_id, description, name, capacity)
	if err != nil {
		return "Error In Inserting Record Into Table :", err
	}

	return "success", nil
}


func InsertIntoRatePlanlTable(statement *sql.Stmt, database *sql.DB, message map[string][]map[string]interface{})(string, error){

	var err error
	var hotel_id, rate_plan_id, cancellation_policy, name , other_conditions, meal_plan string

	for _, value := range message {
		for key,val := range value[0]["rate_plan"].(map[string]interface{}) {
			if key == "hotel_id"{
				hotel_id = val.(string)
			}
			if key == "name"{
				name = val.(string)
			}
			if key == "rate_plan_id"{
				rate_plan_id = val.(string)
			}
			if key == "meal_plan"{
				meal_plan = val.(string)
			}
			if key == "cancellation_policy"{
				resBd, _ := json.Marshal(val)
				cancellation_policy =  string(resBd)
			}
			if key == "other_conditions"{
				resBd, _ := json.Marshal(val)
				other_conditions =  string(resBd)
			}
		}
	}

	statement, err = database.Prepare("INSERT INTO rate_plan (hotel_id, rate_plan_id, cancellation_policy, name , other_conditions, meal_plan) VALUES (?, ?, ?, ?, ?, ?)")

	if err != nil {
		return "Error In Prepare Statement Of Insert Query :", err
	}
	defer statement.Close()

	_, err = statement.Exec(hotel_id, rate_plan_id, cancellation_policy, name , other_conditions, meal_plan)
	if err != nil {
		return "Error In Inserting Record Into Table :", err
	}

	return "success", nil
}