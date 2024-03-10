package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type BucketBasics struct {
	S3Client *s3.Client
}

func main() {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}

	bucketName := "blogs-repo-test"
	s3Client := s3.NewFromConfig(sdkConfig)
	basics := BucketBasics{s3Client}
	a, _ := basics.ListObjects(bucketName)

	for _, v := range a {
		file, err := basics.DownloadFile(bucketName, *v.Key, *v.Key)
		if err != nil {
			log.Printf("Bucket: %v. Here's why: %v\n", bucketName, err)
		}

		createBlog(file, *v.Key)
	}
}

func createBlog(data []byte, fileName string) {
	// selectedFileName, _ := strings.CutSuffix(selectedFile.Name(), ".md")

	dirName := "../my-page/app/" + fileName + "/"

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

	// data, err := os.ReadFile("../s3/" + selectedFile.Name())
	// if err != nil {
	// 	log.Fatal(err)
	// }

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

func (basics BucketBasics) ListObjects(bucketName string) ([]types.Object, error) {
	result, err := basics.S3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	var contents []types.Object
	if err != nil {
		log.Printf("Couldn't list objects in bucket %v. Here's why: %v\n", bucketName, err)
	} else {
		contents = result.Contents
	}
	return contents, err
}

func (basics BucketBasics) DownloadFile(bucketName string, objectKey string, fileName string) ([]byte, error) {
	result, err := basics.S3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Printf("Couldn't get object %v:%v. Here's why: %v\n", bucketName, objectKey, err)
		// return err
	}
	defer result.Body.Close()
	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("Couldn't create file %v. Here's why: %v\n", fileName, err)
		// return err
	}
	defer file.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		log.Printf("Couldn't read object body from %v. Here's why: %v\n", objectKey, err)
	}
	return body, err
}
