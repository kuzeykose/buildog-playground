package main

import (
	"fmt"
	"log"
	"os"
)

//  - next app -
// app
// 	  /layout.jsx
//    /page.jsx

func main() {
	dirName := "./test-blog/app/"

	file, err := os.Create(dirName + "page.js")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Step 2: Read the file
	data, err := os.ReadFile("s3/blog1.md")
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

	fmt.Println("File written successfully")
}
