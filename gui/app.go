package Gui

import (
	"../common"
	"../conohaApi"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"log"
	"net/url"
	"os"
	"time"
)

var err error
var textLog = "loading....."

func Run() {
	os.Setenv("FYNE_FONT", "C:\\windows\\Fonts\\simhei.ttf")

	myApp := app.New()
	myWindow := myApp.NewWindow("Conoha-UI")

	setTabContent(myWindow)

	myWindow.ShowAndRun()
}

func setTabContent(window fyne.Window) {
	textLog = "loading....."
	displayFullScreenLog(window)
	//label_1 := widget.NewLabel("Conoha VPS Controller\n https://https://github.com/terry2010/conoha-ui")
	u, _ := url.Parse("https://https://github.com/terry2010/conoha-ui")
	tab1 := widget.NewTabItemWithIcon("index", theme.HomeIcon(),
		widget.NewVBox(widget.NewLabel("Conoha VPS Controller"),
			widget.NewHyperlink("github home page", u)))
	tabs := widget.NewTabContainer(tab1)
	textLog += "\ngetting server list"
	displayFullScreenLog(window)
	idList, _ := getVPSIDList()
	textLog += "\ngetting server finish"
	displayFullScreenLog(window)

	if len(idList) > 0 {
		for _, v := range idList {
			textLog += "\ngetting server info:" + v
			displayFullScreenLog(window)
			serverInfo, err := ConohaApi.ComputeServerInfo(v)
			textLog += "\ngetting server finish:" + Common.SafeGetError(err)
			displayFullScreenLog(window)
			//tabs.Append(widget.NewTabItem(serverInfo.Name, createServerInfoForm(window, v)))
			tabs.Append(widget.NewTabItem(serverInfo.Name, createServerInfoForm(window, v)))

		}
	}

	widget.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab"))

	tabs.SetTabLocation(widget.TabLocationLeading)

	window.SetContent(tabs)
	//window.Resize(fyne.NewSize(600, 320))
	window.Show()
}

func createServerInfoForm(window fyne.Window, _serverID string) *widget.Box {
	serverInfo, _ := ConohaApi.ComputeServerInfo(_serverID)
	entry := widget.NewEntry()
	entry.Text = serverInfo.Name

	form := widget.NewForm()

	form.Append("DELETE", widget.NewButtonWithIcon("!!!DELETE!!!", theme.DeleteIcon(), func() {
		textLog = "DELETING........."
		displayFullScreenLog(window)
		time.Sleep(time.Duration(1) * time.Second)
		textLog =  textLog + "\n5"
		displayFullScreenLog(window)
		time.Sleep(time.Duration(1) * time.Second)
		textLog =  textLog + "\n4"
		displayFullScreenLog(window)
		time.Sleep(time.Duration(1) * time.Second)
		textLog =  textLog + "\n3"
		displayFullScreenLog(window)
		time.Sleep(time.Duration(1) * time.Second)
		textLog =  textLog + "\n2"
		displayFullScreenLog(window)
		time.Sleep(time.Duration(1) * time.Second)
		textLog =  textLog + "\n1"
		displayFullScreenLog(window)
		time.Sleep(time.Duration(1) * time.Second)
		textLog =  textLog + "\n0.5"
		displayFullScreenLog(window)
		time.Sleep(time.Duration(1) * time.Second)
		textLog =  textLog + "\n0.1"
		displayFullScreenLog(window)
		time.Sleep(time.Duration(1) * time.Second)
		textLog = textLog + "\nDELETING........."
		displayFullScreenLog(window)
		err = ConohaApi.ComputeServerDelete(_serverID)
		log.Println("dddddddddddddddddd",err)
		textLog = textLog + "DELETING finish:" + Common.SafeGetError(err)
		displayFullScreenLog(window)
		setTabContent(window)
		log.Println("DELETE..............................", err)
	}))

	form.Append("ServerName", widget.NewLabel(serverInfo.Name))

	form.Append("status", widget.NewHBox(widget.NewLabel(serverInfo.Status), widget.NewButtonWithIcon("Refresh", theme.ViewRefreshIcon(), func() {
		setTabContent(window)
		log.Println("refresh.................", err)
	})))

	form.Append("Created", widget.NewLabel(serverInfo.Created.String()))
	if "ACTIVE" == serverInfo.Status {
		form.Append("action", widget.NewButtonWithIcon("Stop", theme.CancelIcon(), func() {
			textLog = "Stoping........."
			displayFullScreenLog(window)
			err = ConohaApi.ComputeServerForceShutDown(_serverID)
			textLog = "Stoping finish:" + Common.SafeGetError(err)
			displayFullScreenLog(window)
			setTabContent(window)

		}))
	} else {
		form.Append("action", widget.NewButtonWithIcon("Start", theme.MediaPlayIcon(), func() {
			textLog = "Starting........."
			displayFullScreenLog(window)
			err = ConohaApi.ComputeServerStart(_serverID)
			textLog = "Starting finish:" + Common.SafeGetError(err)
			displayFullScreenLog(window)
			setTabContent(window)
			log.Println("start..................", err)
		}))
	}

	if len(serverInfo.IPV4Address) > 0 {
		for mac, ip := range serverInfo.IPV4Address {
			t := widget.NewEntry()
			t.Text = ip
			form.Append("ipv4", t)
			t1 := widget.NewEntry()
			t1.Text = mac
			form.Append("ipv4:MAC", t1)
		}
	}

	if len(serverInfo.IPV6Address) > 0 {
		for mac, ip := range serverInfo.IPV6Address {
			t := widget.NewEntry()
			t.Text = ip
			form.Append("ipv6", t)
			t1 := widget.NewEntry()
			t1.Text = mac
			form.Append("ipv6:MAC", t1)
		}
	}

	form.Append("action", widget.NewButtonWithIcon("Reboot", theme.ContentRedoIcon(), func() {
		ConohaApi.ComputeServerReboot(_serverID)
		setTabContent(window)
		log.Println("reboot..............................", err)
	}))

	return widget.NewVBox(form)
}

func getVPSIDList() (idList []string, err error) {
	//ConohaApi.Login()
	_serverIDList, err := ConohaApi.GetJsonString("serverIDList.json")

	if nil == err {
		refreshVPSIDList()
		_serverIDList, err = ConohaApi.GetJsonString("serverIDList.json")
	}
	Common.Json.UnmarshalFromString(_serverIDList, &idList)
	return
}

func refreshVPSIDList() (err error) {
	_serverIDList, err := ConohaApi.ComputeServerList()
	log.Println("get Server ID List", _serverIDList)
	if nil == err {
		ConohaApi.SaveJsonString("serverIDList.json", Common.FastJsonMarshal(_serverIDList))
	}
	return err
}

func displayFullScreenLog(window fyne.Window) {
	var logWidget = widget.NewLabel(textLog)

	window.SetContent(widget.NewVBox(widget.NewProgressBarInfinite(), logWidget))

	window.Show()
}
