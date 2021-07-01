## TorrentMonitor Golang client package
*project sucks and code is so sh1tty x_x*

**For now it can:**
* Add given release to TorrentMonitor

## Installation

```shell
go mod download github.com/back2net/tmc
go mod tidy
```
**or use legacy:**
```shell
go get github.com/back2net/tmc
```

## Usage
It provides only one function for now, so your app code may be something like that:

```go
package main

import (
	"fmt"
	"net/url"

	"github.com/back2net/tmc"
)

func main() {

	// mock for test only, delete this later!!!!
	tracker_data := url.Values{
		"action":        {"torrent_add"},
		"name":          {""},
		"url":           {"https://nnmclub.to/forum/viewtopic.php?p=12345"},
		"path":          {"/home/media/storage"},
		"update_header": {"true"},
	}

	// mock for test only, delete this later!!!!
	series_data := url.Values{
		"action":  {"serial_add"},
		"tracker": {"lostfilm.tv"},
		"name":    {"Lost"},
		"hd":      {"1"}, //  0 -SD, 1 -1080p, 2 - 720p
		"path":    {"/home/media/storage"},
	}
	tmc.AddTitleToMonitor(tracker_data)
    	tmc.AddTitleToMonitor(series_data)
}
```
Also you need to create `.env` file with the contents:
TMON_URL=http://your.torrentmonitor.com
TMON_PASSWORD=megasecretpassword


* In work:
- [x] <del>make it work!</del>
- [ ] Nothing 


# Author doesn't undestand what he is doing 
# and only pretends that he is a programmer x_x

Thanks for [ElizarovEugene](github.com/ElizarovEugene/TorrentMonitorfor) his work!
