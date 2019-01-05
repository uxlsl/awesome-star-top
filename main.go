package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
)
type Project struct{
	url string
	star int
}

type ByStar[]Project

func (s ByStar) Len() int {
	return len(s)
}
func (s ByStar) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByStar) Less(i, j int) bool {
	return s[i].star < s[j].star
}

func GetProjectInfo(url string) string{
	resp, err := http.Get(url)
	if err != nil{
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	text := string(body)
	re := regexp.MustCompile("\\d+ users starred this repository")
	str := re.FindString(text)
	num_re := regexp.MustCompile("\\d+")
	loc := num_re.FindStringIndex(str)
	if len(loc) <= 0{
		return ""
	}
	return str[loc[0]:loc[1]]
}

func main(){
	if len(os.Args) <= 1{
		fmt.Println("please give url arg!")
		os.Exit(0)
	}
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	text := string(body)
	re := regexp.MustCompile("https://github.com/\\w+/[\\w-]+")
	urls := re.FindAllString(text,-1)
	urlsMap := make(map[string]bool)
	for _,url := range urls {
		urlsMap[url] = true
	}
	count := 10
	projects := make([]Project, 0)
	for url,_ := range urlsMap{
		star,err := strconv.Atoi(GetProjectInfo(url))
		if err!=nil{
			fmt.Println(err)
			continue
		}
		p := Project{url, star}
		projects = append(projects, p)
		count--
		if count == 0{
			break
		}
	}
	sort.Sort(sort.Reverse(ByStar(projects)))
	fmt.Println("results")
	for _,v := range projects{
		fmt.Println(v)
	}
}
