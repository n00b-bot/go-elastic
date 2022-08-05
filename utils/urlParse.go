package utils

import (
	"bufio"
	"os"
	"strings"
	"sync"
)

func ParseUrl(url string) []string {
	if url == "" {
		return nil
	}
	var n = 1
	path := strings.Split(url, " ")
	if strings.Contains(path[1], "?") {
		path = strings.Split(path[1], "?")
		n = 0
	}
	path = strings.Split(path[n], "/")
	return path[1:]
}

func CheckUrlExists(url, path string) bool {
	var rs bool = false
	if url == "" {
		return false
	}
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	wg := new(sync.WaitGroup)
	//fmt.Println("-------------------------------------")
	for w := 1; w <= 3; w++ {
		wg.Add(1)

		go func() {
			for scanner.Scan() {
				a := scanner.Text()
				//fmt.Println(a)
				if url == a {
					rs = true
					break
				}
			}
			wg.Done()
		}()

	}
	wg.Wait()
	return rs
}
