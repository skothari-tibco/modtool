package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"strings"
	"bytes"
	
	"path/filepath"
	
)
var mainFolder string

func main(){

	//Get all the files and directories 
	files, err := ioutil.ReadDir(os.Args[1])

	if err != nil{
		log.Fatal(err)
	}
	//Get Current folder name
	pwd, _ := os.Getwd()
	
	cwd := strings.Split(pwd, "/")
	
	mainFolder = cwd[len(cwd)-1]

	//Get Directories present in the current list of files obtained from line 18
	dir, mod := contains(files)
	//Check if the dir has go.mod if not declare one
	check(mod)
	
	list := dir
	//Iterate over the list of dir
	for ;  ; {
		if (len(list) == 0){
			break
		}
		
		newList := CheckForSubAndAdd(list[0])
		//Add any sub dir that was found to the list
		list = append(list, newList...)
		//Remove the first entry as it was iterated.
		list = list[1:]
		
	}
	

}

func CheckForSubAndAdd(name string)([] string ){

	err := os.Chdir(name)
	if err != nil{
		fmt.Println(err)
	}

	pwd,_ := os.Getwd()
	files ,_ :=ioutil.ReadDir(pwd)
	
	dir, mod := contains(files)
		
	check(mod)
	
	return dir
}

func check(val bool){
	if !val{
		f, err := os.OpenFile("go.mod", os.O_RDWR|os.O_CREATE ,0755)
		var buffer bytes.Buffer
		if err != nil{
			fmt.Println(err)
		}
		buffer.WriteString("module ")
		buffer.WriteString(os.Args[2])
		pwd, _ := os.Getwd()
		//cwd := strings.Split(pwd, "/")
		index := strings.Index(pwd,mainFolder)
		buffer.WriteString(pwd[index:])
		f.Write([]byte(buffer.String()))
	}else{

	}

}


func isGoFile(name string) bool{

	return strings.Contains(name,".go")
}
func contains(arr []os.FileInfo) (name []string,mod bool){
	mod = false
	for _, file := range arr{
		if (file.Name()=="go.mod"){
			mod = true
		}
		if file.IsDir() && isNotHidden(file.Name()){
			path , err := filepath.Abs(file.Name())
			if err != nil {
				fmt.Println("Error in getting absolute path")
			}
			name = append(name, path)
		}
	}
	 
	return name, mod
}

func isNotHidden(name string)bool{
	return !strings.Contains(name,".")
}

func 