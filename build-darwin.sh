AppName="Bug Reproduction"
Version=$(node -e 'console.log(require("./package.json").version)')
rm -rf ./$AppName.app
fyne package -os darwin -icon Icon.png --exe "$AppName" --appID "myapp.$AppName" --name "$AppName" --appVersion $Version
cp config.json "$AppName.app/Contents/MacOS/"
