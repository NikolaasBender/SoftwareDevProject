#!/bin/bash
#sudo chmod +x builder.sh
filesToBuild='main.go structs.go handlers.go'
finalFile='main'

go build $filesToBuild
./$finalFile