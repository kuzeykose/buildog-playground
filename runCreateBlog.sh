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

rm -rf node_modules
rm -rf package-lock.json
rm -rf .next

mv out ../

cd ..

aws s3 sync ./my-page "s3://buildog-web/${ORGANIZATION_ID}/app/"
aws s3 sync ./out "s3://buildog-web/${ORGANIZATION_ID}/web/"

rm -rf out
rm -rf my-page