package command

import (
	"sync"
)

// pidMap 用于存储所有的子进程pid，以便在主程序退出时，kill掉所有的相关子进程
var pidMap sync.Map
