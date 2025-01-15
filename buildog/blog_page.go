package main

import (
    "fmt"
    "log"
    "os"
)

func CreateBlogPage() {
    dirName := "../my-page/app/blog/"

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

    imports := `
        import ReactMarkdown from "react-markdown"
        import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"

    `

    layout := `
    export default function Page(){
        return(
			<main className="flex-grow flex justify-center px-4">
				<div className="max-w-3xl w-full py-8">
				    <Card>
				    	<CardHeader>
				    	    <CardTitle>Blog Post</CardTitle>
				    	</CardHeader>
				    	<CardContent>
				    	    <article className="prose lg:prose-xl dark:prose-invert">
				    		    <ReactMarkdown>
                                   # Welcome to My Blog
                                </ReactMarkdown>
				    	    </article>
				    	</CardContent>
				    </Card>
				</div>
			</main>
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
