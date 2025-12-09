#! /usr/bin/env bash

echo "--- generating README ---"

cat <<EOF > README.md
# Advent of code 2025

done in golang, some of them on my phone lol
(since inputs are meant to be secret, they are not included here)

| task | time (seconds) | time (nanoseconds) |
| ---- | -------------- | ------------------ |
EOF


for i in {1..12}; do
    for j in {1..2}; do
        filename="./${i}/part-${j}/main.go"
        input="./${i}/input"
        if [ -f ${filename} ] && [ -f ${input} ]; then

            #get time running
            start=$(date +%s%N)

            go run ${filename} ${input}

            end=$(date +%s%N)
            ns=$((end - start))
            seconds=$(printf '%d.%09d' $((ns/1000000000)) $((ns%1000000000)))

            echo "|${i} part ${j} | ${seconds} | ${ns} |" >> README.md
        else
            echo "${filename} does not exists"
            echo "|${i} part ${j} | NA | NA |" >> README.md
        fi
    done
done
