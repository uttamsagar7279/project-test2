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

func GetEmp(w http.ResponseWriter, r *http.Request) {
    db := setupDB()

    printMessage("Getting Emp...")

    // Get all emp from emp table that don't have EmpID = "1"
    rows, err := db.Query(SELECT * FROM Emp WHERE EmpID=="1" AND Empdesignation=="software engineer")

    // check errors
    checkErr(err)

    // var response []JsonResponse
    var Emp []Emp

    // Foreach emp
    for rows.Next() {
        var id int
        var EmpID string
        var Empdesignation string

        err = rows.Scan(&id, &EmpID, &Empdesignation)

        // check errors
        checkErr(err)

        Emp = append(Emp, Emp{EmpID: EmpID, Empdesignation: Empdesignation})
    }

    var response = JsonResponse{Type: "success", Data: Emp}

    json.NewEncoder(w).Encode(response)
}
