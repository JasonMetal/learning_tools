package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"path/filepath"
)

// main
//
//	@Description:
//
// 初始化监控器：创建 fsnotify.Watcher 实例并封装到 Watch 结构体中。
// 递归添加监控路径：遍历指定目录，将所有子目录添加到监控列表。
// 启动监控服务：在一个 goroutine 中监听文件变化事件。
// 事件处理：根据事件类型（创建、写入、删除、重命名）执行相应操作。
// 错误处理：捕获并打印监控过程中产生的错误，然后退出。
func main() {
	watch, _ := fsnotify.NewWatcher()
	w := Watch{
		watch: watch,
	}
	w.watchDir("./filewatch/watch")
	select {}
}

type Watch struct {
	watch *fsnotify.Watcher
}

// watchDir
//
//	@Description: 监控文件夹下的文件的创建、修改、删除事件
//	@receiver w
//	@param dir
func (w *Watch) watchDir(dir string) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			path, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			err = w.watch.Add(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	log.Println("监控服务已经启动")
	go func() {
		for {
			select {
			case ev := <-w.watch.Events:
				{
					if ev.Op&fsnotify.Create == fsnotify.Create {
						fmt.Println("创建文件 : ", ev.Name)
						fi, err := os.Stat(ev.Name)
						if err == nil && fi.IsDir() {
							w.watch.Add(ev.Name)
							fmt.Println("添加监控 : ", ev.Name)
						}
					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						fmt.Println("写入文件 : ", ev.Name)
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						fmt.Println("删除文件 : ", ev.Name)
						fi, err := os.Stat(ev.Name)
						if err == nil && fi.IsDir() {
							w.watch.Remove(ev.Name)
							fmt.Println("删除监控 : ", ev.Name)
						}
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						fmt.Println("重命名文件 : ", ev.Name)
						w.watch.Remove(ev.Name)
					}
				}
			case err := <-w.watch.Errors:
				{
					fmt.Println("error : ", err)
					return
				}
			}
		}
	}()
}
