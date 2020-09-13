#!/bin/zsh
HOME_PATH='../'
BUILD_FILE=$1
EXEC_FILE=`echo $BUILD_FILE | sed s/.go//`

# move to build file path
PATH=`find $HOME_PATH -name $BUILD_FILE | sed s/$BUILD_FILE//`
cd $PATH

# go build
go build

# go exuc
./$EXEC_FILE
