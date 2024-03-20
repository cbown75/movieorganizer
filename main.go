package main

import (
  "fmt"
  "log"
  "os"
  "runtime"
  "path/filepath"
  "strings"
)

func main()  {
  var path = ""
  

  switch opsy := runtime.GOOS; opsy {
    case "darwin":
      path = "/Volumes/movies/Movies/"
    case "windows":
      path = "R:\\Movies\\"
    default:
      path = ""
  }

  if pathExists(path) {
    pattern := path + "*.mkv"
    mkvs, err := filepath.Glob(pattern)

    
    if err != nil {
      log.Fatal(err)
      return
    } else if len(mkvs) == 0 {
      fmt.Println("No New Files Found")
    } else {
      for f := range mkvs {
        newFolder, newFile := splitFolderFile(path, mkvs[f])
        newFolderPath := filepath.Join(path, newFolder)
        destFilePath := filepath.Join(path, newFolder, newFile)

        if !pathExists(newFolderPath) {

          err := os.Mkdir(newFolderPath, os.ModePerm)
          if err != nil && !os.IsExist(err) {
            log.Fatal(err)
            return
          } else {
            fmt.Println("New Folder Created: ", newFolderPath)
          }

          err = os.Rename(mkvs[f], destFilePath)
          if err != nil {
            log.Fatal(err)
            return
          } else {
            fmt.Println("File Moved: ", destFilePath)
          }

        } else {
          fmt.Println("Path Does Exist")
          //add code to compare files and try to figure out if the new one is bigger than the old one
        }
      }
    }
  
  } else {
    fmt.Println(os.Environ())
    fmt.Println("The Drive is Not Mapped Here")
  }
}

func splitFolderFile(path string, file string) (string, string) {
  filename := strings.ReplaceAll(file, path, "")
  foldername := strings.ReplaceAll(filename, ".mkv", "")

  
  return foldername, filename
}


func pathExists(path string) bool {
  _, err := os.Stat(path)

  if err != nil {
    return false
  } else {
    return true
  }
}
