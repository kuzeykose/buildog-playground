package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

func main() {
	Start()
}

func Start() {
	// dirName := "./test-blog/app/"

	dirRead := "../s3"

	files, err := os.ReadDir(dirRead)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Print(file)
		createBlog(file)
	}
}

func createBlog(selectedFile fs.DirEntry) {
	selectedFileName, _ := strings.CutSuffix(selectedFile.Name(), ".md")

	dirName := "../my-page/app/" + selectedFileName + "/"

	// create folder
	err := os.MkdirAll(dirName, 0755) // 0755 is a common permission setting
	if err != nil {
		log.Fatalf("Failed to create directory: %s", err)
	}

	// create file
	file, err := os.Create(dirName + "page.js")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	data, err := os.ReadFile("../s3/" + selectedFile.Name())
	if err != nil {
		log.Fatal(err)
	}

	imports := `
		import { Flex, Text, Button } from '@radix-ui/themes';
		import Markdown from 'react-markdown'

	`
	mdVariable := fmt.Sprintf("const md = `%s`", data)

	code2 := `

	export default function MyApp() {
		return (
			<Flex direction="column" gap="2">
				<Markdown>
					{md}
				</Markdown>
			</Flex>
		);
	}
	`

	_, err = file.WriteString(imports + mdVariable + code2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	newfile, err := os.Create("../my-page/next.config.mjs")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer newfile.Close()

	_, err = newfile.WriteString(`
/** @type {import('next').NextConfig} */
const nextConfig = {
	output: "export",
};

export default nextConfig;	
	`)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File written successfully")
}
