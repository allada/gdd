#!/bin/sh
node ./build-scripts/convert_protocol.js > /dev/null 2>&1;
go build -o bin/server server.go &&
#sh -c "nohup sh -c 'sleep .5; open \"https://chrome-devtools-frontend.appspot.com/serve_file/@939b32ee5ba05c396eef3fd992822fcca9a2e262/inspector.html?experiments=true&ws=localhost:9922/go-devtools-debug\"' </dev/null >/dev/null 2>&1 &" &&
echo "https://chrome-devtools-frontend.appspot.com/serve_file/@3f9be266524354849e17f6a123ee20dd6e7d53e0/inspector.html?experiments=true&ws=localhost:9922/go-devtools-debug"
./bin/server $@