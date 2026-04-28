#!/bin/bash

# =================================================================
# reprobuild 实验自动化脚本
# 用法: ./run_experiment.sh [input_path] [ld_preload_path]
# =================================================================


DEFAULT_INPUT_BASE=""
DEFAULT_LD_PRELOAD=""

INPUT_PATH=${1:-$DEFAULT_INPUT_BASE}
LD_PRELOAD_PATH=${2:-$DEFAULT_LD_PRELOAD}

# 实验项目列表
PROJECTS=("lz4" "sqlite" "TinyCC" "e2fsprogs" "nginx" "zlib" "busybox" "lua" "openssl" "redis")

echo "开始可重复构建实验..."
echo "使用 LD_PRELOAD: $LD_PRELOAD_PATH"

python3 tag_server.py &

# 创建结果保存目录
mkdir -p ./experiment_results

go mod tidy

for PROJECT in "${PROJECTS[@]}"; do
    echo "------------------------------------------------"
    echo "正在测试项目: $PROJECT"
    
    YAML_FILE="$INPUT_PATH/$PROJECT/build_record.yaml"
    
    if [ ! -f "$YAML_FILE" ]; then
        echo "跳过: 未找到配置文件 $YAML_FILE"
        continue
    fi

    go run cmd/c_build/*.go \
        --input="$YAML_FILE" \
        --ld_preload="$LD_PRELOAD_PATH" \
        --debug \
        --create \
        --enable_timer \
        2>&1 | tee "./experiment_results/${PROJECT}_build.log"

    echo "项目 $PROJECT 测试完成，日志已保存至 ./experiment_results/${PROJECT}_build.log"
done

echo "所有实验任务已执行完毕。"