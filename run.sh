#!/bin/sh
node ./build-scripts/convert_protocol.js;
echo "Finished building protocol files." &&
go install &&
echo "Finished building GDD." &&
#sh -c "nohup sh -c 'sleep .5; open \"https://chrome-devtools-frontend.appspot.com/serve_file/@939b32ee5ba05c396eef3fd992822fcca9a2e262/inspector.html?experiments=true&ws=localhost:9922/go-devtools-debug\"' </dev/null >/dev/null 2>&1 &" &&
$GOPATH/bin/gdd $@