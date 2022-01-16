# Global values
MOD_NAME=`cat go.mod | grep -E '^module ' | head -1 | awk -F ' ' '{print $2}'`

echo "###### Clear ###############################"
rm -rf output

echo "###### Package #############################"
mkdir -p output/bin
cp -rf ./conf output/

echo "###### Build ###############################"
go build \
    -o output/bin/go_spider \
    -ldflags "
        -X '${MOD_NAME}/common/glbval.GitCommitID=`git rev-parse HEAD`'
        -X '${MOD_NAME}/common/glbval.BuildTime=`date +"%Y-%m-%d %H:%M:%S"`'
    " \
    main.go
