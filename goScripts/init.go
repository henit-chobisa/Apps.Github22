package main

import (
	appscli "GitpodConfig/Packages/AppsCli"
	"GitpodConfig/Packages/Colors"
	"GitpodConfig/Packages/ConfigReader"
	"GitpodConfig/Packages/DockerCompose"
	"GitpodConfig/Packages/Figure"
	initiateadmin "GitpodConfig/Packages/InitiateAdmin"
	"GitpodConfig/Packages/InstallApp"
	"GitpodConfig/Packages/Logo"
	"fmt"
)

func main() {
	data := ConfigReader.ReadConfig("config.json")
	appData := ConfigReader.ReadConfig(fmt.Sprintf("%v/app.json", data["appDir"]))

	Logo.RocketChat()
	Logo.Custom(fmt.Sprintf("%v App", appData["name"]))
	fmt.Printf("\n\n\n")
	fmt.Println(Colors.Blue() + "Phase 1 : Intiating Rocket Chat Apps Test Environment\n" + Figure.Line())

	DockerCompose.Up(fmt.Sprintf("%v", data["composeFilePath"]))
	appscli.Install()

	fmt.Printf("\n")
	fmt.Println(Colors.Blue() + "Phase 2 : Configuring Rocket.Chat App, installing admin\n" + Figure.Line())
	initiateadmin.Initiate(data)

	fmt.Printf("\n\n")
	fmt.Println(Colors.Blue() + "Phase 3 : Installing App into Rocket.Chat Server\n" + Figure.Line())

	user := data["admin"].(map[string]interface{})

	InstallApp.Install(fmt.Sprintf("%v", data["appDir"]), "http://localhost:3000", fmt.Sprintf("%v", user["username"]), fmt.Sprintf("%v", user["pass"]))

	// appConfig := ConfigReader.ReadConfig(fmt.Sprintf("%v/app.json", data["appDir"]))

	// exec.Command("gp", "sync-done", "user-init").Output()
}
