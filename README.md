__钱包api及后台管理__

本系统采用beego框架，在运行前需要安装golang环境；

1. 安装go环境（linux）

    wget https://studygolang.com/dl/golang/go1.11.linux-amd64.tar.gz

    tar -C /opt -xzf go1.10.4.linux-amd64.tar.gz

    mkdir -p /opt/go

    mkdir -p /opt/gopath

    vim /etc/profile

    添加以下内容：

    #根目录

    export GOROOT=/opt/go

    #bin目录

    export GOBIN=$GOROOT/bin

    #工作目录

    export GOPATH=/opt/gopath

    export PATH=$PATH:$GOPATH:$GOBIN:$GOPATH

    编辑保存并退出vi后，记得把这些环境载入：

    source /etc/profile

    运行以下命令查看当前go的版本，如果能够显示go版本，那么说明我们的go安装成功.

    go version

2. 获取源代码

    cd $GOPATH/src

    git clone http://192.168.1.230/timeShuttle/walletApi.git

    cd walletApi

3. 修改配置文件（conf/app.conf）

    appname = walletApi

    httpport = 9080

    runmode = dev

    dbhost=127.0.0.1

    dbport=3306

    dbuser=root

    dbpassword=xxxxx

    db=xxx

    tablename=t_user

    copyrequestbody = true

    EnableDocs = true

    chainurl = http://localhost:4000/v1/wallet

    jpushurl = https://api.jpush.cn/v3/push

    jpushappkey = xxx

    jpushsecret = xxxx

4. 安装第三方库

    go get github.com/astaxie/beego

    go get github.com/beego/bee

    go get github.com/astaxie/beego/orm

    go get github.com/go-sql-driver/mysql

    go get github.com/satori/go.uuid

    go get github.com/beego/i18n

5. 运行

    首次运行程序时，需要对以下数据进行清理

    DELETE from t_block_ext_info;

    DELETE from t_chain_code;

    DELETE from t_charge_record;

    DELETE from t_idx_gas;

    DELETE from t_idx_record;

    DELETE from t_sign_data;

    DELETE from t_token_info;

    DELETE from t_transaction;

    执行运行命令如下：
    
    bee run -gendoc=true -downdoc=true

    -gendoc=true 表示每次自动化的 build 文档，-downdoc=true 就会自动的下载 swagger 文档查看器

