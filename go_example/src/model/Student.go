package model

import (
	"database/sql"
	"fmt"
)

/**
 *  属性名大写才能被其他包引用到
 */
type Student struct {
	Id int
	// 加上*能接收数据库查询出的null值 不然会报错
	Name  *string
	Sex   *string
	Phone *string
}

func (temp Student) SayHi() {
	name := temp.Name
	fmt.Printf("Student id=%d, name= %q, sex= %q, phone= %q sayhi \n", temp.Id, *name, temp.Sex, temp.Phone)
}

// 执行命令，可以用于插入、删除、修改操作
func ExecuteCommand(db *sql.DB, cmd string) error {
	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Begin Failed")
		return err
	}
	// 准备SQL语句
	stmt, err := tx.Prepare(cmd)
	if err != nil {
		fmt.Println("Prepare Failed")
		return err
	}
	// 将参数传递到SQL语句中执行
	_, err = stmt.Exec()
	if nil != err {
		fmt.Println("Exec Failed: ", cmd)
		return err
	}
	//提交事务
	err = tx.Commit()
	if nil != err {
		fmt.Println("Commit Failed")
		return err
	}
	return nil
}

// 插入数据示例
func InsertStudent(db *sql.DB, stu Student) (id int64, err error) {
	tx, err := db.Begin()
	if nil != err {
		fmt.Println("Begin Failed")
		return -1, err
	}
	// 准备SQL语句
	stmt, err := tx.Prepare("insert into student(`name`,`sex`,`phone`) values(?,?,?)")
	if nil != err {
		fmt.Println("Prepare Failed!")
		return -1, err
	}
	// 将参数传递到SQL语句中执行
	res, err := stmt.Exec(stu.Name, stu.Sex, stu.Phone)
	if nil != err {
		fmt.Println("Exec Failed!")
		return -1, err
	}
	// 将事务提交
	err = tx.Commit()
	if nil != err {
		fmt.Println("Commit Failed")
		return -1, err
	}
	id, err = res.LastInsertId()
	return id, nil
}

// 删除数据示例
func DeleteStudentById(db *sql.DB, id int) (bool, error) {
	tx, err := db.Begin()
	if nil != err {
		fmt.Println("Begin Failed!", err)
		return false, err
	}
	// 准备SQL
	stmt, err := tx.Prepare("delete from student where id = ?")
	if nil != err {
		fmt.Println("prepare sql cmd err:", err)
		return false, err
	}
	// 设置参数以及执行SQL
	res, err := stmt.Exec(id)
	if nil != err {
		fmt.Println("exec failed!")
		return false, err
	}
	// 提交事务
	iId, err := res.LastInsertId()
	fmt.Println("LastInsertId:", iId)
	err = tx.Commit()
	return true, err
}

func UpdateStudent(db *sql.DB, stu Student) (bool, error) {
	//开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Begin failed")
		return false, err
	}
	//准备sql语句
	stmt, err := tx.Prepare("UPDATE student SET name = ?, sex = ?, phone = ? WHERE id = ?")
	if err != nil {
		fmt.Println("Prepare failed")
		return false, err
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(stu.Name, stu.Sex, stu.Phone, stu.Id)
	if err != nil {
		fmt.Println("Exec failed")
		return false, err
	}
	//提交事务
	err = tx.Commit()
	if err != nil {
		fmt.Println("Commit failed")
		return false, err
	}
	fmt.Println(res.LastInsertId())
	return true, nil
}

// 查询单行数据：编写SQL语句，利用QueryRow,并且利用Scan将查询结果赋值到对应的对象中
func SelectById(db *sql.DB, id int) (Student, error) {
	var stu Student
	err := db.QueryRow("select * from student where id = ?", id).Scan(&stu.Id, &stu.Name, &stu.Sex, &stu.Phone)
	if nil != err {
		fmt.Println("QueryRow Error")
		return stu, err
	}
	return stu, nil
}

// 编写sql语句，执行Query函数
// 利用Next()读取每一行返回的结果，并且利用Scan赋值到相应的对象中
// 参考代码如下
func SelectAll(db *sql.DB) ([]Student, error) {
	rows, err := db.Query("select * from student")
	if nil != err {
		fmt.Println("Query Error:", err)
		return nil, err
	}
	var stus []Student
	for rows.Next() {
		var stu Student
		err = rows.Scan(&stu.Id, &stu.Name, &stu.Sex, &stu.Phone)
		if err != nil {
			return nil, err
		}

		stus = append(stus, stu)
	}
	return stus, nil
}
