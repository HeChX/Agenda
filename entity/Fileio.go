package entity

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func BytestoUser(bytes []byte) User {
	var user User
	err := json.Unmarshal(bytes, &user)
	if err != nil {
		fmt.Println("error:", err)
	}
	return user
}

func UsertoBytes(user User) []byte {
	bytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println("error:", err)
	}
	return bytes
}

func BytestoMeeting(bytes []byte) Meeting {
	var meeting Meeting
	err := json.Unmarshal(bytes, &meeting)
	if err != nil {
		fmt.Println("error:", err)
	}
	return meeting
}

func MeetingtoBytes(meeting Meeting) []byte {
	bytes, err := json.Marshal(meeting)
	if err != nil {
		fmt.Println("error:", err)
	}
	return bytes
}

func ReadUsersFromFile(path string) []User {
	var users []User
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		users = append(users, BytestoUser([]byte(line)))
	}
	return users
}

func WriteUsersToFile(path string, users []User) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	os.Truncate(path, 0)
	if err != nil {
		fmt.Println("open file failed.", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	for i := 0; i < len(users); i++ {
		file.WriteString(string(UsertoBytes(users[i])[:]))
		file.WriteString("\n")
	}
}

func ReadMeetingsFromFile(path string) []Meeting {
	var meetings []Meeting
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		meetings = append(meetings, BytestoMeeting([]byte(line)))
	}
	return meetings
}

func WriteMeetingsToFile(path string, meetings []Meeting) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	os.Truncate(path, 0)
	if err != nil {
		fmt.Println("open file failed.", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	for i := 0; i < len(meetings); i++ {
		file.WriteString(string(MeetingtoBytes(meetings[i])[:]))
		file.WriteString("\n")
	}
}
