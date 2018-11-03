package service

import (
	"fmt"
	"log"
	"os"

	"github.com/HeChX/agenda/entity"
)

var currentname string
var currentpassword string
var of *os.File
var Logined = false

func Init() {
	entity.Init()
	logfile, err := os.OpenFile("service/agenda.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	of = logfile
	if err != nil {
		log.Fatalln("Open file error!")
	}
}

func RegisterUser(name string, password string, email string, phone string) {
	logwrite := log.New(of, "[Info]", log.LstdFlags)
	defer of.Close()
	success := entity.Register(name, password, email, phone)
	if success {
		logwrite.Println(name, "Register successfully!")
	} else {
		logwrite.Println(name, "Register failed!")
	}
}

func Login(name string, password string) {
	logwrite := log.New(of, "[Info]", log.LstdFlags)
	defer of.Close()
	user, flag := entity.QueryUser(name)
	if flag {
		currentname = user.Name
		if user.Password != password {
			logwrite.Println(name, "Log in failed!")
			fmt.Println("Password error!")
		} else {
			currentpassword = user.Password
			logwrite.Println(name, " Log in successfully!")
			fmt.Println("Log in successfully!")
			fmt.Println("Welcome to Agenda!")
			Logined = true
		}
	} else {
		logwrite.Println(name, "Log in failed!")
		fmt.Println("User name does not exist")
	}
}

func Logout() {
	logwrite := log.New(of, "[Info]", log.LstdFlags)
	defer of.Close()
	logwrite.Println(currentname, "Log out successfully!")
	fmt.Println("Log out successfully!")
	Logined = false
}

func QueryUser(name string) {
	logwrite := log.New(of, "[Info]", log.LstdFlags)
	defer of.Close()
	user, flag := entity.QueryUser(name)
	if flag {
		logwrite.Println(currentname, " Query user ", name, " successfully!")
		fmt.Println("Name : ", user.Name)
		fmt.Println("Email : ", user.Email)
		fmt.Println("Phone : ", user.Phone)
	}
}

func QueryUserAllUser() {
	logwrite := log.New(of, "[Info]", log.LstdFlags)
	defer of.Close()
	logwrite.Println(currentname, " Query all users successfully!")
	for _, v := range entity.GetUsers() {
		fmt.Println("Name : " + v.Name + " Email: " + v.Email + " Phone: " + v.Phone)
	}
}

func DeleteUser() {
	logwrite := log.New(of, "[Operation]", log.LstdFlags)
	defer of.Close()
	logwrite.Println(currentname, " Logoff successfully!")
	entity.DeleteUser(currentname)
	Logout()
	fmt.Println("Log off successfully!")
}

func CreateMeeting(title string, start string, end string, participators []string) {
	logwrite := log.New(of, "[Operation]", log.LstdFlags)
	defer of.Close()
	start_time, _ := entity.StringtoDate(start)
	end_time, _ := entity.StringtoDate(end)
	if entity.CreateMeeting(title, start_time, end_time, currentname, participators) {
		logwrite.Println(currentname + " creat meeting " + title + " successfully!")
	} else {
		logwrite.Println(currentname + " creat meeting " + title + " failed!")
	}
}

func AddParticipator(title string, name string) {
	logwrite := log.New(of, "[Operation]", log.LstdFlags)
	defer of.Close()
	meeting, flag, _ := entity.QueryMeetingByTitle(title)
	if !(meeting.Sponser == currentname) {
		logwrite.Println(currentname + " add participator " + name + " to " + title + " failed!")
		fmt.Println("Add participator " + name + " to " + title + " failed!")
	} else {
		if flag {
			if entity.AddParticipator(title, name) {
				logwrite.Println(currentname + " add participator " + name + " to " + title + " successfully!")
				fmt.Println("Add participator " + name + " to " + title + " successfully!")
			} else {
				logwrite.Println(currentname + " add participator " + name + " to " + title + " failed!")
				fmt.Println("Add participator " + name + " to " + title + " failed!")
			}
		}
	}
}

func DeleteParticipator(title string, name string) {
	logwrite := log.New(of, "[Operation]", log.LstdFlags)
	defer of.Close()
	meeting, flag, _ := entity.QueryMeetingByTitle(title)
	if !(meeting.Sponser == currentname) {
		logwrite.Println(currentname + " delete participator " + name + " from " + title + " failed!")
		fmt.Println("Delete participator " + name + " from " + title + " failed!")
	} else {
		if flag {
			if entity.DeleteParticipator(title, name) {
				logwrite.Println(currentname + " delete participator " + name + " from " + title + " successfully!")
				fmt.Println("Delete participator " + name + " from " + title + " successfully!")
			} else {
				logwrite.Println(currentname + " delete participator " + name + " from " + title + " failed!")
				fmt.Println("Delete participator " + name + " from " + title + " failed!")
			}
		}
	}
}

func QueryMeeting(start string, end string) {
	logwrite := log.New(of, "[Operation]", log.LstdFlags)
	defer of.Close()
	start_time, _ := entity.StringtoDate(start)
	end_time, _ := entity.StringtoDate(end)
	meetings, flag := entity.QueryMeetingByTime(start_time, end_time, currentname)
	if flag {
		logwrite.Println(currentname + " query meeting successfully!")
		for _, v := range meetings {
			starttime := entity.DatetoString(v.Time_start)
			endtime := entity.DatetoString(v.Time_end)
			fmt.Println("Title : " + v.Title)
			fmt.Println("Start time : " + starttime)
			fmt.Println("End time : " + endtime)
			fmt.Println("Sponsor : " + v.Sponser)
			fmt.Println("Paticipator: ", v.Participator)
		}
	} else {
		logwrite.Println(currentname + " query meeting failed!")
		fmt.Println("Query meeting failed!")
	}
}

func CancelMeeting(title string) {
	logwrite := log.New(of, "[Operation]", log.LstdFlags)
	defer of.Close()
	flag := entity.CancelMeeting(title, currentname)
	if flag {
		logwrite.Println(currentname + " cancel meeting " + title + " failed!")
		fmt.Println("Cancel meeting " + title + " failed!")
	} else {
		logwrite.Println(currentname + " cancel meeting " + title + " successfully!")
		fmt.Println("Cancel meeting " + title + " successfully!")
	}
}

func ExitMeeting(title string) {
	logwrite := log.New(of, "[Operation]", log.LstdFlags)
	defer of.Close()
	flag := entity.DeleteParticipator(title, currentname)
	if flag {
		logwrite.Println(currentname + " exit meeting " + title + " failed!")
		fmt.Println("Exit meeting " + title + " failed!")
	} else {
		logwrite.Println(currentname + " exit meeting " + title + " successfully!")
		fmt.Println("Exit meeting " + title + " successfully!")
	}
}

func ClearMeeting() {
	logwrite := log.New(of, "[Operation]", log.LstdFlags)
	defer of.Close()
	entity.ClearMeeting(currentname)
	logwrite.Println(currentname + "clear meeting!")
	fmt.Println("Clear all meetings!")
}
