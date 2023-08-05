package main

// func connect(c echo.Context) error {
// 	db, _ := model.DB.DB()
// 	defer db.Close()
// 	err := db.Ping()
// 	if err != nil {
// 		return c.String(http.StatusInternalServerError, "DB接続失敗しました")
// 	}

// 	return c.String(http.StatusOK, "DB接続しました")
// }

func main() {

	r := newRouter()
	r.Logger.Fatal(r.Start(":8080"))
}
