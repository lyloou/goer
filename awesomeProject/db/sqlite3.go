package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "foo.db")
	checkErr2(err)

	fmt.Println(db)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr2(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr2(err)

	id, err := res.LastInsertId()
	checkErr2(err)

	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr2(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr2(err)

	affect, err := res.RowsAffected()
	checkErr2(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr2(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr2(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr2(err)

	res, err = stmt.Exec(id)
	checkErr2(err)

	affect, err = res.RowsAffected()
	checkErr2(err)

	fmt.Println(affect)

	db.Close()

}

func checkErr2(err error) {
	if err != nil {
		panic(err)
	}
}
