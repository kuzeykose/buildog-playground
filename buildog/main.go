package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

type BucketBasics struct {
	S3Client *s3.Client
}

func main() {
	ctx := context.Background()

	// Initialize the AWS S3 client
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Println("Couldn't initialize AWS configuration. Have you set up your AWS credentials?")
		fmt.Println(err)
		return
	}
	s3Client := s3.NewFromConfig(cfg)

	CreateHeader()
	CreateFooter()
	CreateBlogPage()
	CreateBlogLayout()

	bucketName := "buildog-web"
	prefix := os.Getenv("ORGANIZATION_ID") + "/documents/"
	basics := BucketBasics{s3Client}

	// List and process documents
	err = basics.ProcessDocuments(ctx, bucketName, prefix)
	if err != nil {
		log.Fatalf("Error processing documents: %v", err)
	}
}

// ProcessDocuments lists objects under a specific prefix and processes them
func (basics BucketBasics) ProcessDocuments(ctx context.Context, bucketName, prefix string) error {
	paginator := s3.NewListObjectsV2Paginator(basics.S3Client, &s3.ListObjectsV2Input{
		Bucket:    &bucketName,
		Prefix:    &prefix,
		Delimiter: aws.String("/"),
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return fmt.Errorf("error listing bucket objects: %v", err)
		}

		for _, object := range page.Contents {
			if strings.HasSuffix(*object.Key, ".md") {
				fmt.Printf("Processing file: %s\n", *object.Key)
				data, err := basics.DownloadFile(ctx, bucketName, *object.Key)
				if err != nil {
					log.Printf("Error downloading file %s: %v\n", *object.Key, err)
					continue
				}

				// Extract folder name and file name
				relativePath := strings.TrimPrefix(*object.Key, prefix)
				_, file := path.Split(relativePath)
				fileName := strings.TrimSuffix(file, ".md")

				// Process the document
				createBlog(data, fileName)
			}
		}
	}

	return nil
}

// DownloadFile fetches a file's content from the bucket
func (basics BucketBasics) DownloadFile(ctx context.Context, bucketName, objectKey string) ([]byte, error) {
	result, err := basics.S3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
	})
	if err != nil {
		return nil, fmt.Errorf("couldn't get object %s from bucket %s: %v", objectKey, bucketName, err)
	}
	defer result.Body.Close()

	data, err := io.ReadAll(result.Body)
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
