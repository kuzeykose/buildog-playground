package main

import (
	"fmt"
	"log"
	"os"
)

func CreateBlogLayout() {
	dirName := "../my-page/app/blog/"

	// create folder
	err := os.MkdirAll(dirName, 0755) // 0755 is a common permission setting
	if err != nil {
		log.Fatalf("Failed to create directory: %s", err)
	}

	// create file
	file, err := os.Create(dirName + "layout.js")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	imports := `
		import Header from '@/components/Header'
		import Footer from '@/components/Footer'

	`

	layout := `
    export default function Layout({children}){
	    return(
            <div className="flex flex-col min-h-screen">
      		    <Header />
                    {children}
			    <Footer />
		    </div>
        )
    }
	`

	_, err = file.WriteString(imports + layout)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File written successfully")
}
