#/bin/bash
source /etc/profile
cur_dir=$(pwd)
echo $cur_dir
go env -w GO111MODULE=auto
export GOPROXY=https://goproxy.io


cd $cur_dir
go build -o dy dy.go
if [ $? -eq 0 ]; then
	echo "build dy ok!"
else
	echo "build dy false!"
	exit -1;
fi

cd index
go build -o index index.go
if [ $? -eq 0 ]; then
	echo "build index ok!"
else
	echo "build index false!"
	exit -1;
fi

cd  $cur_dir
cd news
go build -o news news.go
if [ $? -eq 0 ]; then
	echo "build news ok!"
else
	echo "build news false!"
	exit -1;
fi

cd $cur_dir
cd topic
go build -o topic topic.go
if [ $? -eq 0 ]; then
	echo "build topic ok!"
else
	echo "build topic false!"
	exit -1;
fi

cd $cur_dir
cd update
go build -o update update.go
if [ $? -eq 0 ]; then
	echo "build update ok!"
else
	echo "build update false!"
	exit -1;
fi

cd $cur_dir
cd fix
go build -o fix fix.go
if [ $? -eq 0 ]; then
	echo "build fix ok!"
else
	echo "build fix false!"
	exit -1;
fi

cd $cur_dir
cd tool
go build -o tool tool.go
if [ $? -eq 0 ]; then
	echo "build tool ok!"
else
	echo "build tool false!"
	exit -1;
fi


cd $cur_dir
cd import_image
go build -o importImage importImage.go
if [ $? -eq 0 ]; then
	echo "build importImage ok!"
else
	echo "build importImage false!"
	exit -1;
fi


cd $cur_dir
cd tkcj
go build -o tk_cj tk_cj.go
if [ $? -eq 0 ]; then
	echo "build tk_cj ok!"
else
	echo "build tk_cj false!"
	exit -1;
fi

cd $cur_dir
cd upload_tk
go build -o upload_tk upload_tk.go
if [ $? -eq 0 ]; then
	echo "build upload_tk ok!"
else
	echo "build upload_tk false!"
	exit -1;
fi

cd $cur_dir
cd fix_tk_desc
go build -o fix_tk_desc fix_tk_desc.go
if [ $? -eq 0 ]; then
	echo "build fix_tk_desc ok!"
else
	echo "build fix_tk_desc false!"
	exit -1;
fi

