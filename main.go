package main

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/gosimple/slug"
)

type FileKind string

const (
	FileKind_Directory FileKind = "dir"
	FileKind_File      FileKind = "file"
	FileKind_Section   FileKind = "section"
	FileKind_Text      FileKind = "text"
	FileKind_Unknown   FileKind = "unknown"
)

const (
	BasePath string = "./content/en"
	Tab      string = "    " // four spaces
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	menu, err := os.ReadFile("./menu.md")
	check(err)
	menuItems := parse(strings.TrimSpace(string(menu)))

	// sort.Slice(menuItems, func(i, j int) bool {
	// 	return menuItems[i].Level > menuItems[j].Level
	// })

	// create the folder structure
	// var weight int = 1
	uniqueFolders := make(map[string][]MenuItem)
	sections := make(map[string][]string)

	for _, menuItem := range menuItems {
		menuItem.DetectKind()
		menuItem.FixFilename()
		menuItem.FixParentPath()
		menuItem.FixText()

		var folderPath string
		if len(menuItem.ParentPath) == 0 {
			folderPath = fmt.Sprintf("/%s", slug.MakeLang(menuItem.Name, "en"))
		} else {
			for _, item := range menuItem.ParentPath {
				folderPath = strings.TrimSpace(fmt.Sprintf("%s/%s", folderPath, item))
			}
		}

		folderPath = BasePath + folderPath
		menuItem.FolderPath = folderPath
		if _, ok := uniqueFolders[folderPath]; !ok {
			uniqueFolders[folderPath] = []MenuItem{}
		}

		uniqueFolders[folderPath] = append(uniqueFolders[folderPath], *menuItem)
	}

	// we need the sections before we write the files
	for path, menuItems := range uniqueFolders {
		for _, menuItem := range menuItems {
			if menuItem.Kind == FileKind_Section {
				sections[path] = append(sections[path], menuItem.Name)
			}
		}
	}

	for _, menuItem := range menuItems {
		// write the directories
		if menuItem.Kind == FileKind_Directory {
			// fmt.Println(folderPath, BasePath+"/"+menuItem.Filename)
			folderPath := menuItem.FolderPath

			if folderPath != BasePath+"/"+menuItem.Filename {
				folderPath = folderPath + "/" + menuItem.Filename
			}

			if _, err := os.Stat(folderPath); os.IsNotExist(err) {
				// fmt.Println(folderPath)
				os.MkdirAll(folderPath, fs.FileMode(0777))

				// write the _index.md file
				filePath := folderPath + "/_index.md"
				temp := template.Must(template.ParseFiles("./templates/_index.tpl"))
				f, err := os.Create(filePath)
				if err != nil {
					log.Println("create file: ", err)
					return
				}
				temp.Execute(f, &FolderIndex{
					Title:       menuItem.Name,
					LinkTitle:   menuItem.Name,
					Description: menuItem.Name + " Overview",
					Sections:    sections[folderPath],
					Text:        menuItem.Text,
				})
				f.Close()
			}
		}
	}

	for folderPath, menuItems := range uniqueFolders {
		for i, menuItem := range menuItems {
			weight := (i + 1) * 10
			if menuItem.Kind == FileKind_File {
				filePath := folderPath + "/" + menuItem.Filename + ".md"

				// fmt.Println(sections[folderPath+"/"+menuItem.Filename])

				if _, err := os.Stat(filePath); os.IsNotExist(err) {
					// fmt.Println(menuItem.Filename)

					temp := template.Must(template.ParseFiles("./templates/file.tpl"))
					f, err := os.Create(filePath)
					if err != nil {
						log.Println("create file: ", err)
						return
					}

					temp.Execute(f, &File{
						Title:       menuItem.Name,
						LinkTitle:   menuItem.Name,
						Weight:      weight,
						Description: "This page describes " + menuItem.Name,
						Sections:    sections[folderPath+"/"+menuItem.Filename],
					})

					// fmt.Println(folderPath+"/"+menuItem.Filename, sections[folderPath+"/"+menuItem.Filename])
					f.Close()
				}
			}
		}
	}
}

type FolderIndex struct {
	Title       string
	LinkTitle   string
	Description string
	Sections    []string
	Text        string
}

type File struct {
	Title       string
	LinkTitle   string
	Weight      int
	Description string
	Sections    []string
	Text        string
}

type MenuItem struct {
	Kind       FileKind
	Name       string
	Filename   string
	ParentPath []string
	Level      int
	Sections   []string
	FolderPath string
	Text       string
}

func (m *MenuItem) DetectKind() {
	if m.Kind != "" {
		return
	}
	kind := m.Name[0:2]
	switch kind {
	case "D-":
		m.Kind = FileKind_Directory
	case "F-":
		m.Kind = FileKind_File
	case "S-":
		m.Kind = FileKind_Section
	case "T-":
		m.Kind = FileKind_Text
	default:
		m.Kind = FileKind_Unknown
	}

	// fmt.Println(m.Name, kind, m.Kind)
}

func (m *MenuItem) FixFilename() {
	regex := regexp.MustCompile(`^[DFST-]{2}`)
	m.Name = regex.ReplaceAllString(m.Name, "")

	regex = regexp.MustCompile(`^[dfst-]{2}`)
	m.Filename = regex.ReplaceAllString(m.Filename, "")
}

func (m *MenuItem) FixParentPath() {
	var fixed []string
	for _, item := range m.ParentPath {
		regex := regexp.MustCompile(`^[DFST-]{2}`)
		fixed = append(fixed, slug.MakeLang(regex.ReplaceAllString(item, ""), "en"))
	}
	m.ParentPath = fixed
}

func (m *MenuItem) FixText() {
	regex := regexp.MustCompile(`^[DFST-]{2}`)
	m.Text = regex.ReplaceAllString(m.Text, "")
}

func parse(menu string) []*MenuItem {
	var menuItems []*MenuItem
	// convert all tags to spaces
	menu = strings.ReplaceAll(menu, "\t", Tab)
	// split the lines
	lines := strings.Split(menu, "\n")

	var levelPath []int
	var stringPath []string
	for _, line := range lines {
		// fmt.Print(line)
		level := strings.Count(line, Tab)
		// levels := strings.Split(line, Tab)
		lineText := strings.TrimSpace(line)
		if level == 0 {
			levelPath = []int{level}
			stringPath = []string{lineText}
		} else {
			// do we add to the path or remove from the path?
			if len(levelPath) >= level {
				levelPath = levelPath[0:level]
				stringPath = stringPath[0:level]
			}

			levelPath = append(levelPath, level)
			stringPath = append(stringPath, lineText)
		}
		// fmt.Println(" ::", levelPath)

		fileName := slug.MakeLang(stringPath[len(stringPath)-1], "en")

		// this makes the current value of stringPath immutable
		var pathArr []string
		pathArr = append(pathArr, stringPath...)

		menuItems = append(menuItems, &MenuItem{
			ParentPath: pathArr[:len(pathArr)-1], // remove last element
			Name:       stringPath[len(stringPath)-1],
			Level:      level,
			Filename:   fileName,
		})
	}

	return menuItems
}
