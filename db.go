package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type Settings struct {
	Model           string `json:"model"`
	Prompt          string `json:"prompt"`
	WelcomeMessages string `json:"welcomeMessages"`
	Title           string `json:"title"`
	InputText       string `json:"input_text"`
	SpeakText       string `json:"speak_text"`
	User1Path       string `json:"user1path"`
	User2Path       string `json:"user2path"`
	ChatGround      string `json:"chat_ground"`
}

var (
	settings Settings
	once     sync.Once
)

// 初始化设置
func initSettings() {
	once.Do(func() {
		// 检查文件是否存在
		if _, err := os.Stat("settings.json"); os.IsNotExist(err) {
			// 如果文件不存在，创建一个带有默认设置的文件
			settings = Settings{
				Model:           "llama3:latest",
				Prompt:          "you are an interesting ai",
				WelcomeMessages: "Hello!,I am ollama !,Nice to meet you !",
				Title:           "chat with ollama",
				InputText:       "ollama is typing...",
				SpeakText:       "ollama is speaking...",
				User1Path:       "",
				User2Path:       "",
				ChatGround:      "",
			}
			file, _ := json.MarshalIndent(settings, "", " ")
			_ = ioutil.WriteFile("settings.json", file, 0644)
		} else {
			// 如果文件存在，读取设置文件
			file, err := ioutil.ReadFile("settings.json")
			if err != nil {
				fmt.Println(err)
				return
			}

			// 解析设置文件
			err = json.Unmarshal(file, &settings)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	})
}

// 获取设置
func GetSetting() map[string]string {
	settingsMap := make(map[string]string)
	settingsMap["model"] = settings.Model
	settingsMap["prompt"] = settings.Prompt
	settingsMap["welcomeMessages"] = settings.WelcomeMessages
	settingsMap["title"] = settings.Title
	settingsMap["input_text"] = settings.InputText
	settingsMap["speak_text"] = settings.SpeakText
	settingsMap["user1path"] = settings.User1Path
	settingsMap["user2path"] = settings.User2Path
	settingsMap["chat_ground"] = settings.ChatGround
	return settingsMap
}

// 更新设置
func UpdateSettings(newSettings map[string]string) error {
	for key, value := range newSettings {
		switch key {
		case "model":
			settings.Model = value
		case "prompt":
			settings.Prompt = value
		case "welcomeMessages":
			settings.WelcomeMessages = value
		case "title":
			settings.Title = value
		case "input_text":
			settings.InputText = value
		case "speak_text":
			settings.SpeakText = value
		case "user1path":
			settings.User1Path = value
		case "user2path":
			settings.User2Path = value
		case "chat_ground":
			settings.ChatGround = value
		default:
			return fmt.Errorf("unknown setting: %s", key)
		}
	}

	// 将新的设置写入文件
	file, err := json.MarshalIndent(settings, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("settings.json", file, 0644)
	if err != nil {
		return err
	}

	return nil
}
