package main

import (
	"fmt"
	"log"
	"os"
)

func CreateHeader() {
	dirName := "../my-page/components/"

	// create folder
	err := os.MkdirAll(dirName, 0755) // 0755 is a common permission setting
	if err != nil {
		log.Fatalf("Failed to create directory: %s", err)
	}

	// create file
	file, err := os.Create(dirName + "Header.jsx")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	imports := `
		import Link from 'next/link'
		import { Button } from "@/components/ui/button"

	`

	codeBlock := `
		export default function Header() {
			return (
				<header className="border-b">
					<div className="container mx-auto max-w-3xl py-4">
						<nav className="flex justify-between items-center">
						<Link href="/" className="text-2xl font-bold">
							My Blog
						</Link>
						<div className="flex space-x-2">
							<Button variant="ghost" asChild>
							<Link href="/">Home</Link>
							</Button>
							<Button variant="ghost" asChild>
							<Link href="/about">About</Link>
							</Button>
							<Button variant="ghost" asChild>
							<Link href="/contact">Contact</Link>
							</Button>
						</div>
						</nav>
					</div>
				</header>
			)
		}
	`

	_, err = file.WriteString(imports + codeBlock)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Header written successfully")
}

func CreateFooter() {
	dirName := "../my-page/components/"

	// create folder
	err := os.MkdirAll(dirName, 0755) // 0755 is a common permission setting
	if err != nil {
		log.Fatalf("Failed to create directory: %s", err)
	}

	// create file
	file, err := os.Create(dirName + "Footer.jsx")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	footer := `
		export default function Footer() {
			return (
				<footer className="border-t mt-8">
        			<div className="container mx-auto max-w-3xl py-4 text-center text-sm text-muted-foreground">
          				<p>&copy; {new Date().getFullYear()} My Blog. All rights reserved.</p>
        			</div>
      			</footer>
			)
		}
	`

	_, err = file.WriteString(footer)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Footer written successfully")
}
