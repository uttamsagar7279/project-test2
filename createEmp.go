const (
	DB_USER     = "postgres"
	DB_PASSWORD = "12345678"
	DB_NAME     = "Emp"
)

// DB set up
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return DB
}

type Movie struct {
	EmpID   string `json:"EmpID"`
	Empdesignation string `json:" Empdesignation"`
}

type JsonResponse struct {
	Type    string `json:"type"`
	Data    []Emp  `json:"data"`
	Message string `json:"message"`
}

func CreateEmp(w http.ResponseWriter, r *http.Request) {
	EmpID := r.FormValue("EmpID")
	Empdesignation := r.FormValue(" Empdesignation")

	var response = JsonResponse{}

	if EmpID == "" || Empdesignation == "" {
		response = JsonResponse{Type: "error", Message: "You are missing EmpID or Empdesignation parameter."}
	} else {
		db := setupDB()

		printMessage("Inserting Emp into DB")

		fmt.Println("Inserting new Emp with ID: " + EmpID + " and name: " + Empdesignation)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO Emp(EmpID, Empdesignation) VALUES("1", "software engineer ") returning id;", EmpID, Empdesignation).Scan(&lastInsertID)

		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The emp has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}
