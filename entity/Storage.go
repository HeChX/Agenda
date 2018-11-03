package entity

import (
	"fmt"
	"regexp"
)

var users []User
var meetings []Meeting

func Init() {
	tmp_u := ReadUsersFromFile("User.txt")
	tmp_m := ReadMeetingsFromFile("Meeting.txt")
	for i := 0; i < len(tmp_u); i++ {
		users = append(users, tmp_u[i])
	}
	for i := 0; i < len(tmp_m); i++ {
		meetings = append(meetings, tmp_m[i])
	}
}

func GetUsers() []User {
	return users
}

func Register(name string, password string, email string, phone string) bool {
	if len(name) == 0 || len(password) < 6 || !ValidEmail(email) || !ValidPhone(phone) {
		fmt.Println("The registration information format is incorrect!")
		return false
	}
	_, exist := QueryUser(name)
	if exist {
		fmt.Println("The username already exists!")
	}
	user := User{
		Name:     name,
		Password: password,
		Email:    email,
		Phone:    phone,
	}
	users = append(users, user)
	WriteUsersToFile("User.txt", users)
	fmt.Println("Register successfully!")
	return true
}

func ValidEmail(email string) bool {
	valid, _ := regexp.MatchString("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$", email)
	return valid
}

func ValidPhone(phone string) bool {
	valid, _ := regexp.MatchString("^1[0-9]{10}$", phone)
	return valid
}

func QueryUser(name string) (User, bool) {
	var use User
	for _, v := range users {
		if v.Name == name {
			return v, true
		}
	}
	return use, false
}

func CreateMeeting(title string, start Date, end Date, sponser string, participator []string) bool {
	if len(participator) == 0 {
		fmt.Println("The number of people attending the meeting can not be 0!")
		fmt.Println("Create meeting failed!")
		return false
	}

	if !IsValid(start) || !IsValid(end) {
		fmt.Println("Time error!")
		fmt.Println(start, end)
		fmt.Println("Create meeting failed!")
		return false
	}

	if Dateless(end, start) {
		fmt.Println("Start time should not greater than end time!")
		fmt.Println("Create meeting failed!")
		return false
	}

	for _, v := range meetings {
		if v.Title == title {
			fmt.Println("The meeting already exists!")
			fmt.Println("Create meeting failed!")
			return false
		}
	}
	_, exist := QueryUser(sponser)
	if !exist {
		fmt.Println("The user of " + sponser + " does not exist!")
		fmt.Println("Create meeting failed!")
		return false
	}
	for _, v := range participator {
		_, exist := QueryUser(v)
		if !exist {
			fmt.Println("The user of " + v + " does not exist!")
			fmt.Println("Create meeting failed!")
			return false
		}
	}
	user := append(participator, sponser)
	if HaveSameUser(user) {
		fmt.Println("Please do not repeat the same user!")
		fmt.Println("Create meeting failed!")
		return false
	}

	for _, v := range user {
		meeting := QueryMeetingByName(v)
		for _, w := range meeting {
			if DateOverlap(w.Time_start, w.Time_end, start, end) {
				fmt.Println("The meeting time conflicts with the meeting " + w.Title + " time!")
				fmt.Println("Create meeting failed!")
				return false
			}
		}
	}
	metg := Meeting{
		Sponser:      sponser,
		Title:        title,
		Participator: participator,
		Time_start:   start,
		Time_end:     end,
	}
	meetings = append(meetings, metg)
	WriteMeetingsToFile("Meeting.txt", meetings)
	fmt.Println("Create meeting successfully!")
	return true

}

func CancelMeeting(title string, sponsor string) bool {
	_, exist, index := QueryMeetingByTitle(title)
	if exist {
		if meetings[index].Sponser == sponsor {
			meetings = append(meetings[:index], meetings[index+1:]...)
			WriteMeetingsToFile("Meeting.txt", meetings)
			return true
		}
	}
	return false
}

func QueryMeetingByTitle(title string) (Meeting, bool, int) {
	var met Meeting
	for i := 0; i < len(meetings); i++ {
		if meetings[i].Title == title {
			return meetings[i], true, i
		}
	}
	return met, false, 0
}

func QueryMeetingByTime(start Date, end Date, name string) ([]Meeting, bool) {
	var met []Meeting
	if !IsValid(start) || !IsValid(end) {
		fmt.Println("Time error!")
		fmt.Println("Create meeting failed!")
		return met, false
	}

	if Dateless(end, start) {
		fmt.Println("Start time should not greater than end time!")
		fmt.Println("Create meeting failed!")
		return met, false
	}
	mts := QueryMeetingByName(name)
	for _, v := range mts {
		if DateOverlap(start, end, v.Time_start, v.Time_end) {
			met = append(met, v)
		}
	}
	return mts, true
}

func QueryMeetingBySponser(sponser string) []Meeting {
	var mts []Meeting
	for _, v := range meetings {
		if v.Sponser == sponser {
			mts = append(mts, v)
		}
	}
	return mts
}

func QueryMeetingByParticipator(name string) []Meeting {
	var mts []Meeting
	for _, v := range meetings {
		for _, w := range v.Participator {
			if w == name {
				mts = append(mts, v)
			}
		}
	}
	return mts
}

func QueryMeetingByName(name string) []Meeting {
	spmeeting := QueryMeetingBySponser(name)
	pameeting := QueryMeetingByParticipator(name)
	spmeeting = append(spmeeting, pameeting...)
	return spmeeting
}

func AddParticipator(title string, name string) bool {
	mt, exist, index := QueryMeetingByTitle(title)
	_, est := QueryUser(name)
	if !est {
		fmt.Println("The user does not exist!")
		return false
	}
	if exist {
		for _, v := range mt.Participator {
			if v == name {
				fmt.Println("The user is already a member of the meeting!")
				return false
			}
		}
		mts := QueryMeetingByName(name)
		for _, v := range mts {
			if DateOverlap(v.Time_start, v.Time_end, mt.Time_start, mt.Time_end) {
				fmt.Println("The user participation in meeting " + v.Title + " conflicts with the meeting time!")
				return false
			}
		}
		meetings[index].Participator = append(meetings[index].Participator, name)
		WriteMeetingsToFile("Meeting.txt", meetings)
		fmt.Println("Added successfully!")
		return true
	}
	fmt.Println("The meeting does not exist!")
	return false
}

func DeleteParticipator(title string, name string) bool {
	mt, exist, index := QueryMeetingByTitle(title)
	_, est := QueryUser(name)
	if !est {
		fmt.Println("The user does not exist!")
		return false
	}
	if exist {
		for i := 0; i < len(mt.Participator); i++ {
			if mt.Participator[i] == name {
				meetings[index].Participator = append(meetings[index].Participator[:i], meetings[index].Participator[i+1:]...)
				if len(meetings[index].Participator) == 0 {
					CancelMeeting(title, meetings[index].Sponser)
				}
				WriteMeetingsToFile("Meeting.txt", meetings)
				fmt.Println("Delete success!")
				return true
			}
		}
		fmt.Println("The user is not at the meeting!")
		return false
	}
	if mt.Sponser == name {
		CancelMeeting(title, name)
		WriteMeetingsToFile("Meeting.txt", meetings)
		fmt.Println("Delete success!")
		return true
	}
	fmt.Println("The meeting does not exist!")
	return false
}

func DeleteUser(name string) bool {
	success := false
	for i := 0; i < len(users); i++ {
		if users[i].Name == name {
			users = append(users[:i], users[i+1:]...)
			success = true
		}
	}
	mts := QueryMeetingByName(name)
	for _, v := range mts {
		DeleteParticipator(v.Title, name)
	}
	if success {
		fmt.Println("User logoff success!")
		return true
	}
	fmt.Println("User logoff failure!")
	return false
}

func ClearMeeting(sponsor string) {
	for i := 0; i < len(meetings); i++ {
		if meetings[i].Sponser == sponsor {
			meetings = append(meetings[:i], meetings[i+1:]...)
		}
	}
	WriteMeetingsToFile("Meeting.txt", meetings)
}

func ExitMeeting(name string, title string) {
	DeleteParticipator(title, name)
}

func HaveSameUser(user []string) bool {
	for i := 0; i < len(user); i++ {
		for j := 0; j < len(user); j++ {
			if i != j && user[i] == user[j] {
				return true
			}
		}
	}
	return false
}
