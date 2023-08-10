#!/bin/bash
# 撰寫人員: Neil_Hsieh
# 撰寫日期：2023/05/08
# 說明： 啟動 im-chat-logic local 的服務
#
# 備註：
#   

# 取 OS 系統
SYSTEM=$(uname)
# 執行專案的目錄
WORK_PATH=""
# 執行各容器，須掛載的資料夾位置
VOLUME_PATH=""

# 執行 RunService.sh 的目錄(透過readlink 獲取執行腳本的絕對路徑，再透過dirname取出目錄)
if [ "$SYSTEM" = "Linux" ]
then
    WORK_PATH=$(dirname $(readlink -f $0))
    VOLUME_PATH=$(dirname $(readlink -f $0))/../
fi

# for mac
if [ "$SYSTEM" = "Darwin" ]
then
    # 檢查指令是否存在，不存在直接安裝
    isExist=$(which greadlink)
    if [ -z $isExist ]
    then
        echo "指令不存在，開始安裝 greadlink"
        brew install coreutils
    fi

    WORK_PATH=$(dirname $(greadlink -f $0))
    VOLUME_PATH=$(dirname $(greadlink -f $0))/../storage
fi

#############################
#############################

# 判斷是否需要建立 docker 新網路
docker network ls | grep "web_service" >/dev/null 2>&1
    if  [ $? -ne 0 ]; then
        docker network create web_service
    fi

cp $WORK_PATH/../.env.default $WORK_PATH/../.env

# 判斷是否需要寫入 .env.default
less $WORK_PATH/../.env|grep "VOLUME_PATH="
if [ $? -ne 0 ]; then
    echo "\nVOLUME_PATH=$VOLUME_PATH">>$WORK_PATH/../.env
fi

# 啟動容器服務
docker-compose up -d
