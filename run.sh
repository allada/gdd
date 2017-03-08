#!/bin/sh
node ./build-scripts/convert_protocol.js &&
go build -o bin/server server.go &&
#sh -c "nohup sh -c 'sleep .5; open \"https://chrome-devtools-frontend.appspot.com/serve_file/@939b32ee5ba05c396eef3fd992822fcca9a2e262/inspector.html?experiments=true&ws=localhost:9922/go-devtools-debug\"' </dev/null >/dev/null 2>&1 &" &&
echo "https://chrome-devtools-frontend.appspot.com/serve_file/@939b32ee5ba05c396eef3fd992822fcca9a2e262/inspector.html?experiments=true&ws=localhost:9922/go-devtools-debug\n"
./bin/server