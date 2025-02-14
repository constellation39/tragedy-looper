#!/bin/bash

# Function to prompt for directory with default value
get_directory() {
    local default_dir="."
    read -p "输入要扫描的目录 (默认为当前目录): " dir
    echo ${dir:-$default_dir}
}

# Function to prompt for output file with default value
get_output_file() {
    local default_file="output.md"
    read -p "输入输出文件名 (默认为 output.md): " output
    echo ${output:-$default_file}
}

# Function to prompt for file filter pattern
get_filter_pattern() {
    read -p "输入要过滤的文件正则表达式 (例如: \.go$|\.git$, 直接回车表示不过滤): " pattern
    echo "$pattern"
}

# Function to count lines in a file
count_lines() {
    wc -l < "$1"
}

# Function to create tree structure with file contents
create_tree() {
    local prefix="$2"
    local dir="$1"
    
    # Loop through all items in directory
    for item in "$dir"/*; do
        # Skip if item doesn't exist
        [ ! -e "$item" ] && continue
        
        # Get basename of item
        local name=$(basename "$item")
        
        # Skip if matches filter pattern (if provided)
        if [ ! -z "$filter_pattern" ] && [[ "$name" =~ $filter_pattern ]]; then
            continue
        fi
        
        # If directory
        if [ -d "$item" ]; then
            echo "${prefix}├── ${name}/" >> "$output_file"
            create_tree "$item" "${prefix}│   "
        # If file
        elif [ -f "$item" ]; then
            echo "${prefix}├── ${name}" >> "$output_file"
            echo "" >> "$output_file"
            if [ -z "$filter_pattern" ] || ! [[ "$name" =~ $filter_pattern ]]; then
                echo "\`\`\`" >> "$output_file"
                # Check if file is binary
                if file "$item" | grep -q "text"; then
                    cat "$item" >> "$output_file"
                    echo "\`\`\`" >> "$output_file"
                    echo "" >> "$output_file"
                    echo "Total lines: $(count_lines "$item")" >> "$output_file"
                else
                    echo "[Binary file]" >> "$output_file"
                    echo "\`\`\`" >> "$output_file"
                fi
                echo "" >> "$output_file"
            fi
        fi
    done
}

# Get user input
echo "=== 目录树生成器 ==="
source_dir=$(get_directory)
output_file=$(get_output_file)
filter_pattern=$(get_filter_pattern)

# Validate directory
if [ ! -d "$source_dir" ]; then
    echo "错误: 目录 '$source_dir' 不存在"
    exit 1
fi

# Initialize output file
echo "# Directory Structure and Contents" > "$output_file"
echo "Generated at: $(date)" >> "$output_file"
echo "" >> "$output_file"

# Start creating tree from root directory
create_tree "$source_dir" ""

echo "✅ 目录树结构已保存到 $output_file"
