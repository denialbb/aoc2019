#!/bin/bash

read -p "Enter your session cookie: " cookie
read -p "Enter the day number: " day

mkdir "day$day"

url="https://adventofcode.com/2019/day/$day"

curl $url -o "day$day/input" --cookie "session=$cookie"
