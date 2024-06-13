# BiliBili RSS 

## How to use 

### Install

```sh
go get github.com/Ysoding/bilibilirss
```

### Usage
```go
package main

import (
	"fmt"

	"github.com/Ysoding/bilibilirss"
)

func main() {
	c := bilibilirss.New("cookie", "uid")
	fmt.Println(c.GetUpLikeVideo("208259"))
}

```

## Features
- UP 主点赞视频
