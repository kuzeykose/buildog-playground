package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"

	"cloud.google.com/go/storage"
)

type BucketBasics struct {
	StorageClient *storage.Client
}

func main() {
	ctx := context.Background()

	// Initialize the GCS client
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println("Couldn't initialize Google Cloud Storage client. Have you set up your Google Cloud credentials?")
		fmt.Println(err)
		return
	}
	defer storageClient.Close()

	CreateHeader()
	CreateFooter()
	CreateBlogPage()
	CreateBlogLayout()

	bucketName := "buildog"
	prefix := "ad8c0536-1c84-4ce1-9386-b5d0fea6fab0/" // The root folder prefix
	basics := BucketBasics{storageClient}

	// List and process documents
	err = basics.ProcessDocuments(ctx, bucketName, prefix)
	if err != nil {
		log.Fatalf("Error processing documents: %v", err)
	}
}

// ProcessDocuments lists objects under a specific prefix and processes them
func (basics BucketBasics) ProcessDocuments(ctx context.Context, bucketName, prefix string) error {
	it := basics.StorageClient.Bucket(bucketName).Objects(ctx, &storage.Query{
		Prefix:    prefix,
		Delimiter: "/", // To treat '/' as a folder separator
	})

	for {
		obj, err := it.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error iterating through bucket objects: %v", err)
		}

		if strings.HasSuffix(obj.Name, ".md") { // Only process .md files
			fmt.Printf("Processing file: %s\n", obj.Name)
			data, err := basics.DownloadFile(ctx, bucketName, obj.Name)
			if err != nil {
				log.Printf("Error downloading file %s: %v\n", obj.Name, err)
				continue
			}

			// Extract folder name and file name
			relativePath := strings.TrimPrefix(obj.Name, prefix)
			dir, file := path.Split(relativePath)
			fmt.Print(dir)
			fileName := strings.TrimSuffix(file, ".md")

			// Process the document
			createBlog(data, fileName)
		}
	}

	return nil
}

// DownloadFile fetches a file's content from the bucket
func (basics BucketBasics) DownloadFile(ctx context.Context, bucketName, objectName string) ([]byte, error) {
	rc, err := basics.StorageClient.Bucket(bucketName).Object(objectName).NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("couldn't read object %s from bucket %s: %v", objectName, bucketName, err)
	}
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("couldn't read object body: %v", err)
	}
	return data, nil
}

func createBlog(data []byte, fileName string) {
	dirName := "../my-page/app/blog/" + fileName + "/"

	err := os.MkdirAll(dirName, 0755) // 0755 is a common permission setting
	if err != nil {
		log.Fatalf("Failed to create directory: %s", err)
	}

	file, err := os.Create(dirName + "page.js")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

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

	fmt.Println("File written successfully")
}
