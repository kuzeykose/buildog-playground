npx create-next-app@14 my-page --js --tailwind --app --use-npm --no-eslint --no-src-dir --no-import-alias

# sudo chmod -R 777 my-page
cd my-page
npm install
npm install react-markdown
npm install -D @tailwindcss/typography

# install shadcn
npx shadcn@latest init -d 
npx shadcn add -a


cd ..

cd buildog
go run .

cd ..
cd my-page
npm run build

mv out ../

# cd ..
# rm -rf my-page

# aws s3 sync ./out "s3://test-os-buildog"

# rm -rf out