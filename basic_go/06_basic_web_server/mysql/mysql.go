package mysqlpackage

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id    string
	name  string
	age   int
	grade int
}

// * koneksi dengan database
func connect() (*sql.DB, error) {
	// membuka database dengan format
	//root@tcp(127.0.0.1:3306)/db_belajar_golang
	// user => root
	// password =>
	// host => 127.0.0.1 atau localhost
	// port => 3306
	// dbname => db_belajar_golang

	// * untuk database engine yang lain (postgres, oracle dll)
	// bisa cek di dokumentasi, karena ada settingan yang berbeda

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}

	log.Println("Connected to mysql!")
	return db, nil
}

// * pembacaan data dari database
func sqlQuery() {
	db, err := connect()
	if err != nil {
		log.Println(err.Error())
		return
	}
	// setiap membuat koneksi baru jangan lupa di close
	defer db.Close()

	/* -------------------------------------------------------------------------- */

	// pengambilan data dengan method query
	// mengembalikan kumpulan baris
	age := 27
	rows, err := db.Query("SELECT id, name, grade FROM tb_student WHERE age = ?", age)
	if err != nil {
		log.Println(err.Error())
		return
	}
	// ini juga jangan lupa di close
	defer rows.Close()

	result := []student{}

	// looping pada rows hasil query diatas
	for rows.Next() {
		// menyiapkan struct kosong untuk menampung baris data
		each := student{}

		// mengambil record pada iterasi sekarang dan disimpan pada variabel pointer struct
		err = rows.Scan(&each.id, &each.name, &each.grade)

		if err != nil {
			log.Println(err.Error())
			return
		}

		// simpan data baris pada kumpulan data struct student ([]student)
		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		log.Println(err.Error())
		return
	}

	for _, each := range result {
		log.Println(each.name)
	}
}

// * pembacaan data pada satu baris saja
// menggunakan method QueryRow() yang dichain dengan Scan()
func sqlQueryRow() {
	db, err := connect()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer db.Close()

	/* -------------------------------------------------------------------------- */

	result := student{}
	id := "E001"
	err = db.QueryRow("SELECT name, grade from tb_student WHERE ID = ?", id).Scan(&result.name, &result.grade)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Printf("name: %s, grade: %d\n", result.name, result.grade)
}

// * eksekusi query menggunakan Prepare()
// prepare digunakan untuk menulis template query, sehingga dapat di re-use nanti
// method ini bida digabung dengan Query() dan QueryRow()
func sqlPrepare() {
	db, err := connect()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer db.Close()

	/* -------------------------------------------------------------------------- */

	stmt, err := db.Prepare("SELECT name, age, grade FROM tb_student WHERE id = ?")
	if err != nil {
		log.Println(err.Error())
		return
	}

	result1 := student{}
	id1 := "B001"
	err = stmt.QueryRow(id1).Scan(&result1.name, &result1.age, &result1.grade)
	if err != nil {
		log.Println(err.Error())
		return
	}

	result2 := student{}
	id2 := "B002"
	err = stmt.QueryRow(id2).Scan(&result2.name, &result2.age, &result2.grade)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Printf("Name 1: %s\n", result1.name)
	log.Printf("Name 2: %s\n", result2.name)
}

// * insert, update, dan delete dengan Exec()
// direkomendasikan untuk menggunakan method Exec() karena pada operasi ini tidak menghasilkan nilai kembalian row
// tidak seperti dengan Query() dan QueryRow()
// * insert
func sqlInsert(data student) {
	db, err := connect()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer db.Close()

	/* -------------------------------------------------------------------------- */
	stmt, err := db.Prepare("INSERT INTO tb_student VALUES (?,?,?,?)")
	if err != nil {
		log.Println(err.Error())
		return
	}

	_, err = stmt.Exec(data.id, data.name, data.age, data.grade)
	if err != nil {
		log.Println(err.Error())
		return
	}
	// log.Println(result.LastInsertId())
	// log.Println(result.RowsAffected())
	log.Println("data berhasil ditambahkan")
}

// * update
func sqlUpdateAge(id string, age int) {
	db, err := connect()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer db.Close()

	/* -------------------------------------------------------------------------- */

	stmt, err := db.Prepare("UPDATE tb_student SET age = ? WHERE id = ?")
	if err != nil {
		log.Println(err.Error())
		return
	}

	_, err = stmt.Exec(age, id)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("data berhasil diperbarui")
}

// * delete
func sqlDelete(id string) {
	db, err := connect()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer db.Close()

	/* -------------------------------------------------------------------------- */

	stmt, err := db.Prepare("DELETE FROM tb_student WHERE id = ?")
	if err != nil {
		log.Println(err.Error())
		return
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("data berhasil dihapus")
}

func main() {
	// sqlQuery()

	// sqlQueryRow()

	// sqlPrepare()

	// data := student{
	// 	"I001",
	// 	"Iru",
	// 	21,
	// 	2,
	// }
	// sqlInsert(data)

	// sqlUpdateAge("I001", 20)

	sqlDelete("I001")
}
