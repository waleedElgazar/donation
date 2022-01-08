package DbFunctions

import (
	"Donation/Configration"
	"Donation/Models"
	"crypto/rand"
	"io"
	"log"
)

func InsertUser(user Models.User)error {
	db:=Configration.SetUpDB()
	defer db.Close()
	query:="INSERT INTO user set userName=?, userEmail=?, userPassword=?, userPhone=? "

	insert,err:=db.Prepare(query)
	if err!=nil{
		log.Println("error while executing insert query", err.Error())
		return err
	}
	_,err=insert.Exec(user.Name,user.Email,user.Password,user.Phone)
	if err!=nil{
		log.Println("error while parsing data", err.Error())
		return err
	}
	return nil
}

func RetrieveUser(id int) (Models.User,error) {
	db:=Configration.SetUpDB()
	defer db.Close()
	query:="SELECT * FROM user WHERE userId=?"
	var user Models.User
	row:=db.QueryRow(query,id)
	err:=row.Scan(&user.ID,&user.Name,&user.Email,&user.Password,&user.Phone)
	if err!=nil{
		log.Println("error while parsing data", err.Error())
		return user,err
	}
	return user,nil
}

func RetrieveUserByEmail(email string) (Models.User,error) {
	db:=Configration.SetUpDB()
	defer db.Close()
	query:="SELECT * FROM user WHERE userEmail=?"
	var user Models.User
	row:=db.QueryRow(query,email)
	err:=row.Scan(&user.ID,&user.Name,&user.Email,&user.Password,&user.Phone)
	if err!=nil{
		log.Println("error while parsing data", err.Error())
		return user,err
	}
	return user,nil
}

func RetrieveAllUsers()([]Models.User,error){
	db:=Configration.SetUpDB()
	defer db.Close()
	query:="SELECT * FROM user"
	rows,err:=db.Query(query)
	if err!=nil{
		log.Println("error while parsing data", err.Error())
		return nil,err
	}
	var users []Models.User
	for rows.Next(){
		var user Models.User
		err:=rows.Scan(&user.ID,&user.Name,&user.Email,&user.Password,&user.Phone)
		if err!=nil{
			log.Println("error while parsing data", err.Error())
			return nil,err
		}
		users=append(users,user)
	}
	return users,nil
}

func UpdateTheUser(userId int,user Models.User)error{
	db:=Configration.SetUpDB()
	defer db.Close()
	query:="UPDATE donation.user set userName=?, userPassword=?, userPhone=? WHERE userId=?"
	update,err:=db.Prepare(query)
	if err!=nil{
		log.Println("error while executing update query", err.Error())
		return err
	}
	_,err=update.Exec(user.Name,user.Password,user.Phone,userId)
	if err!=nil{
		log.Println("error while parsing data", err.Error())
		return err
	}
	return nil
}

func DeleteTheUser(userId int)error{
	db:=Configration.SetUpDB()
	defer db.Close()
	query:="DELETE FROM user WHERE userId=?"
	delete,err:=db.Prepare(query)
	if err!=nil{
		log.Println("error while executing delete query", err.Error())
		return err
	}
	_,err=delete.Exec(userId)
	if err!=nil{
		log.Println("error while parsing data", err.Error())
		return err
	}
	return nil
}

func AddVerificationCode(verify Models.VerificationCode)error{
	db:=Configration.SetUpDB()
	defer db.Close()
	query:="INSERT INTO verifications set userEmail =?,verificationCode=?"
	insert,err:=db.Prepare(query)
	if err!=nil{
		log.Println("error while executing insert query", err.Error())
		return err
	}
	_,err=insert.Exec(verify.UserEmail,verify.VerificationCode)
	if err!=nil{
		log.Println("error while parsing data", err.Error())
		return err
	}
	return nil
}

func CreateOtp()string{
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	max := 6
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}