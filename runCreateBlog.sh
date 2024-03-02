sudo npx create-next-app@latest my-page --js --tailwind --app --use-npm --no-eslint --no-src-dir --import-alias @

chmod -R 777 my-page
cd my-page
sudo npm install
sudo npm install @radix-ui/themes
sudo npm install react-markdown

cd ..

cd buildog
go run .

cd ..
cd my-page
npm run build

mv out ../

cd ..
rm -rf my-page

aws s3 sync ./out "s3://test-os-buildog"

rm -rf out