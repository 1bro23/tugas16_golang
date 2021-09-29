package main

import "fmt"
import "database/sql"
import _ "mysql-master"

type orang struct{
  id int
  nama string
  status string
}

func main(){
  sql_tampil()
}

func connection() (*sql.DB, error){
  db, err := sql.Open("mysql","root:@tcp(localhost:3306)/belajar")
  if err!=nil{
    return nil, err
  }
  return db, nil
}

func sql_tampil(){
  db, err := connection()
  if err!=nil{
    fmt.Println(err.Error())
    return
  }
  defer db.Close()

  rows, err := db.Query("select * from table_orang")
  if err!=nil{
    fmt.Println(err.Error())
    return
  }
  defer rows.Close()

  var result []orang

  for rows.Next(){
    var each = orang{}

    var err = rows.Scan(&each.id, &each.nama, &each.status)

    if err!=nil{
      fmt.Println(err.Error())
      return
    }

    result = append(result, each)
  }
  if err=rows.Err(); err!=nil{
    fmt.Println(err.Error())
    return
  }

  for _, each:= range result{
    fmt.Println(each.nama,each.status)
  }
}
