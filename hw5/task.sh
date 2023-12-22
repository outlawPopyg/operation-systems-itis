RED='\e[31m'
GREEN='\e[32m'
RESET='\e[0m'

declare -a numbers
declare -i count=1
declare -i guessed_count=0

while :
do
  random_number=$((RANDOM % 10))
  echo Step: $count
  read -p "Please enter number from 0 to 9 (q-quit): " input

  case ${input} in
  [0-9])
    if [ ${random_number} == "${input}" ]
      then
        echo "Hit! My number: ${random_number}"
        number_string="${GREEN}${input}${RESET}"

        guessed_count+=1
      else
        echo "Miss! My number: ${random_number}"
        number_string="${RED}${input}${RESET}"
    fi
    numbers+=($number_string)
    let hit_percent=guessed_count*100/count
    let miss_percent=100-hit_percent
    echo "Hit: ${hit_percent}%" "Miss: ${miss_percent}%"
    arr_len=${#numbers[@]}

    echo -n "Numbers "
    if [ $arr_len -lt 10 ]; then
      echo -e "${numbers[@]}"
    else
      echo -e "${numbers[@]: -10}"
    fi
    count+=1
    ;;
  q)
    echo "Bye"
    break
    ;;
  *)
    echo "Not valid input"
  esac
done
