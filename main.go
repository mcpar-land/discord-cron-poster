package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/fatih/color"
)

var printErr = color.New(color.FgRed).PrintlnFunc()

var gray = color.New(color.FgHiBlack)
var purp = color.New(color.FgMagenta)
var green = color.New(color.FgGreen)
var yellow = color.New(color.FgYellow)
var blue = color.New(color.FgBlue)
var cyan = color.New(color.FgCyan)

func main() {
	configPath := flag.String("config", "./config.json", "config file")
	url := flag.String("url", "", "url of webhook")
	mediaPath := flag.String("media", ".", "root of media files. Defaults to the same folder as the config file")
	tz := flag.String("tz", "", "TZ timezone for cron jobs. Example: America/New_York")

	flag.Parse()

	cfgDat, err := ioutil.ReadFile(*configPath)
	if err != nil {
		printErr("Error loading config file:", err)
		os.Exit(1)
	}
	configString := string(cfgDat)

	config, err := NewConfiguration(configString)
	if err != nil {
		printErr("Error parsing config json:", err)
		os.Exit(1)
	}

	// override JSON file config options with ones from flags
	if *url != "" {
		config.Url = *url
	}
	if *mediaPath != "" {
		config.MediaPath = *mediaPath
	}
	if *tz != "" {
		config.Timezone = *tz
	}
	// re-write the media path to be relative to the config file
	config.MediaPath = path.Join(path.Dir(*configPath), config.MediaPath)

	if config.Url == "" {
		printErr("No Webhook URL provided!")
		printErr("Either include a URL in your config.json, or use the --url flag")
		os.Exit(1)
	}

	gray := color.New(color.FgHiBlack).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	fmt.Printf("\n%s\n", gray("Starting service..."))

	service, err := NewService(config)
	if err != nil {
		printErr("Error creating service:", err)
		os.Exit(1)
	}

	service.Start()
	fmt.Printf("\n%s\n", green("Service started!"))
	fmt.Println("")
	fmt.Println(service.config.JobsString())
	fmt.Println("")

	select {}
}
