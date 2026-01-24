#!/bin/bash

# SETTINGS
tests_dir="./acceptance-tests/"
input_dir="${tests_dir}input/"
output_dir="${tests_dir}output/"
user_dir="${tests_dir}user/"
path_to_bin="./uniq"
# SETTINGS

function parse_params () {
  local flags=""
 
  cdu_flag=$(echo "$1" | grep -oP "[cdu]")
  if [ "$cdu_flag" != "" ] && [ ${#cdu_flag} -eq 1 ]; then
    flags="${flags}-$cdu_flag "
  fi
  
  if [[ $(echo "$1" | grep -oP "i") != "" ]]; then
    flags="${flags}-i "
  fi
  
  if [[ $(echo "$1" | grep -oP "f") != "" ]]; then
    flags="${flags}-f $(echo "$1" | grep -oP "(?<=f)\d+(?=s|\.)") "
  fi
  
  if [[ $(echo "$1" | grep -oP "s") != "" ]]; then
    flags="${flags}-s $(echo "$1" | grep -oP "(?<=s)\d+(?=\.)") "
  fi
  
  echo "$flags"
}

function print_msg () {
  if [ "$3" == "0" ]; then 
      echo -e "\e[32m$1$2: SUCCESS\e[0m"
  else
      echo -e "\e[31m$1$2: FAIL\e[0m"
  fi 
}

function test_with_std () {
  "$path_to_bin" $params < "$path_to_input" > "$path_to_user"
  difference=$(diff "$path_to_output" "$path_to_user")
  success="$?"
  
  print_msg "$test_prefix" "--" "$success"
  if [ "$success" != "0" ]; then
      echo -e "$difference"
      fail=true
  fi
}

function test_with_input() {
  "$path_to_bin" $params  "$path_to_input" > "$path_to_user"
  diff "$path_to_output" "$path_to_user" > /dev/null
  success="$?"
  
  print_msg "$test_prefix" "I-" "$success"
  if [ "$success" != "0" ]; then
      fail=true
  fi
}

function test_with_input_and_output() {
  "$path_to_bin" $params "$path_to_input" "$path_to_user"
  diff "$path_to_output" "$path_to_user" > /dev/null
  success="$?"
  
  print_msg "$test_prefix" "IO" "$success"
  if [ "$success" != "0" ]; then
      fail=true
  fi
}

fail=false
input_files=$(ls $input_dir)
for input_file in $input_files
do
  test_number="${input_file:0:2}"

  output_user_file="${test_number}.txt"
  
  path_to_input="$input_dir$input_file"
  path_to_output="$output_dir$output_user_file"
  path_to_user="$user_dir$output_user_file"

  test_prefix="Test[${test_number}]"

  params=$(parse_params "$input_file")
  
  test_with_std
  
  test_with_input

  test_with_input_and_output
done

echo "==================="
if $fail ; then
  echo -e "\e[31mSUMMARY   : FAIL\e[0m"
  exit 1
fi
echo -e "\e[32mSUMMARY   : SUCCESS\e[0m"
