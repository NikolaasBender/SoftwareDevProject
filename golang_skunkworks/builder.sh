#!/bin/bash
#sudo chmod +x builder.sh
filesToBuild='main.go structs.go handlers.go assetFix.go'
finalFile='main'

go build $filesToBuild
./$finalFile