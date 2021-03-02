package main

import (
	"fmt"
	"regexp"
)


type Phone struct{
	name string
	email string
	number string
}

func main() {
	var inputMenu,inputDalamMenu,inputNamaSementara,InputEmailSementara string
	var inputNomorSementara string
	var ListBook = []Phone{}
	var phone Phone
	for {
		menu()
		fmt.Scanln(&inputMenu)
		if inputMenu == "!"{
			fmt.Println("Anda Sudah Keluar Dari Apilikasi")
			break
		}else if inputMenu == "1" {
			fmt.Println()
			fmt.Println("Anda masuk Ke Menu Tambah Data  Nomor")
			for {
				fmt.Println()
				fmt.Print("Masukan Nama Anda : ")
				fmt.Scanln(&phone.name)
				if  len(phone.name) <=5 && checkHuruf(phone.name){
					fmt.Println("[Minimal 5 Huruf Dan Tidak Boleh Angka !]")
					break
				}else {
					fmt.Println("Nama Anda adalah:", phone.name)
				}
				fmt.Println()
				fmt.Print("Masukan Email Anda : ")
				fmt.Scanln(&phone.email)
				if checkEmail(phone.email){
					fmt.Println("Email Anda adalah:", phone.email)
				}else{
					fmt.Println("[Email Yang Anda Masukan Tidak Valid!]")
					break
				}
				fmt.Println()
				fmt.Print("Masukan Nomer Anda :  ")
				fmt.Scanln(&phone.number)
				if len(phone.number) <= 10 && checkNomor(phone.number) == false{
					fmt.Println("[Nomor Tidak Boleh Kurang Dari 10 Digit dan Tidak Boleh Huruf!!]")
					fmt.Println("Nomor Anda Adalah" , phone.number)
					fmt.Println()
					break
				}
				fmt.Println()
				ListBook= append(ListBook, phone)
				fmt.Println("Nama :",phone.name,"Email :", phone.email, "Nomor : ", phone.number)
				fmt.Println()
				fmt.Println("Nomor Sudah Tersimpan, Ketik [!] Untuk Kembali Ke Menu")
				fmt.Println()
				fmt.Scanln(&inputDalamMenu)
				if inputDalamMenu == "!"{
					fmt.Println("Anda eluar dari menu")
					fmt.Println()

					break
				}else {
					fmt.Println("")
					continue
				}
			}
		} else if inputMenu == "2" {
			fmt.Println("Anda Masuk Ke Menu ListBook")
			fmt.Println(ListBook)
			fmt.Println("Apakah Anda Ingin Keluar Dari Menu ? Ketik [!] Untuk Kembali Ke Menu Ketik [Y] Untuk Keluar Apilikasi")
			fmt.Scanln(&inputDalamMenu)
				if inputDalamMenu == "!"{
					fmt.Println("Anda Keluar dari Menu")
					menu()
				}else if inputDalamMenu == "Y" {
					fmt.Println("Anda Keluar dari aplikasi")
					break
				}
		}else if inputMenu == "3"{
			fmt.Println("Anda masuk ke menu update email")
			fmt.Print("Masukan Nama List Contact Yang Akan anda Update :   ")
			fmt.Scanln(&inputNamaSementara)
			fmt.Println()
			fmt.Print("Masukan Email Baru :")
			fmt.Scanln(&InputEmailSementara)
			if !checkEmail(InputEmailSementara) {
				fmt.Println("[Email Yang Anda Masukan Tidak Valid!]")
				break
			}
			for i:= 0; i < len(ListBook);i++{
				if ListBook[i].name == inputNamaSementara{
					ListBook[i].email = InputEmailSementara
					fmt.Println("Email sudah terupdate")
				}else{
					fmt.Println("Error")
				}
				break			
			}
		}else if inputMenu == "4"{
			fmt.Println("Anda masuk ke menu update nomor")
			fmt.Print("Masukan Nama List Contact Yang Akan Di Update : ")
			fmt.Scanln(&inputNamaSementara)
			fmt.Println()
			fmt.Print("Masukan nomor baru :  ")
			fmt.Scanln(&inputNomorSementara)
			if !checkNomor(inputNomorSementara){
				fmt.Println("[Masukan Minimal 10 Digit dan Tidak Boleh Huruf!]")
			}

			for i:= 0; i < len(ListBook);i++{
				if ListBook[i].name == inputNamaSementara{
					ListBook[i].number = inputNomorSementara
					fmt.Println("Nomor sudah terupdate")
				}else{
					fmt.Println("Error")
				}
				break			
			}
		}
	}
}

func menu(){
	fmt.Println("Selamat Datang Di Simple App PhoneBook")
	fmt.Println("1. Input Data")
	fmt.Println("2. List Book")
	fmt.Println("3. Update Email")
	fmt.Println("4. Update Nomor")
	fmt.Print("Masukan Menu pilihan Anda :")
}

func checkHuruf(name string) bool {
	isStringAlphabetic,_ := regexp.MatchString(`^[a-z A-Z]+$`, name)
	return isStringAlphabetic
}

func checkNomor(nomor string) bool {
	isStringNumeric,_ := regexp.MatchString(`^[\d]+$`, nomor)
	return isStringNumeric
}

func checkEmail(email string) bool{
	isStringValid,_ := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",email)
	return isStringValid
}