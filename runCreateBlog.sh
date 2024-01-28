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
npm run dev



