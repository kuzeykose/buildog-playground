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
		import Markdown from 'react-markdown'

	`
	mdVariable := fmt.Sprintf("const md = `%s`", data)

	code2 := `

	export default function MyApp() {
		return (
			<article class="prose lg:prose-xl">			
				<Markdown>
					{md}
				</Markdown>
			</article>
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

	tailwindConfig()

	fmt.Println("File written successfully")
}

func tailwindConfig() {
	newfile, err := os.Create("../my-page/tailwind.config.js")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer newfile.Close()

	_, err = newfile.WriteString(`
	/** @type {import('tailwindcss').Config} */
	module.exports = {
	  content: [
		"./pages/**/*.{js,ts,jsx,tsx,mdx}",
		"./components/**/*.{js,ts,jsx,tsx,mdx}",
		"./app/**/*.{js,ts,jsx,tsx,mdx}",
	  ],
	  theme: {
		extend: {
		  backgroundImage: {
			"gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
			"gradient-conic":
			  "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
		  },
		},
	  },
	  plugins: [
    	require('@tailwindcss/typography'),
  		],
	};
	
	`)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
