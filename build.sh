#/bin/bash
source /etc/profile
echo $HOME
go env -w GO111MODULE=auto

cd -
go build -o dy dy.go
if [ $? -eq 0 ]; then
	echo "build dy ok!"
else
	echo "build dy false!"
	exit -1;
fi

cd -
cd index
go build -o index index.go
if [ $? -eq 0 ]; then
	echo "build index ok!"
else
	echo "build index false!"
	exit -1;
fi

cd -
cd news
go build -o news news.go
if [ $? -eq 0 ]; then
	echo "build news ok!"
else
	echo "build news false!"
	exit -1;
fi

cd -
cd topic
go build -o topic topic.go
if [ $? -eq 0 ]; then
	echo "build topic ok!"
else
	echo "build topic false!"
	exit -1;
fi

cd -
cd update
go build -o update update.go
if [ $? -eq 0 ]; then
	echo "build update ok!"
else
	echo "build update false!"
	exit -1;
fi

