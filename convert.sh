#!/bin/bash

# .webpファイルを、.jpgに変換するスクリプト
# .webpファイルが入ったフォルダを引数に指定

if ! test $(which dwebp 2>/dev/null); then
    echo "not found dwebp command"
    exit 1
fi

if [ $# -ne 1 ] || [ "$1" == "" ]; then
    echo "no directory input"
    exit 1
fi

echo "convert webp files in ($1) folder"

cd $1

for e in $(ls *.webp)
do
    after=(${e//webp/png})
    # echo "$e -> $after"
    dwebp $e -o $after
done

rm *.webp && echo "delete *.webp files"