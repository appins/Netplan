#!/bin/bash

for d in entries/*/; do
  FILE="${d::-1}.json"
  echo -n '{"entries" : [' > $FILE
  for x in $d*; do
    if [[ $x == *theme.setting ]]; then
      continue
    else
      echo -n \"$(cat $x)\"\, >> $FILE
    fi
  done
  FILECONTENTS="$(cat $FILE)"
  echo "${FILECONTENTS::-1}]," > $FILE
  echo \"theme\" \: \"$(cat "$d/theme.setting")\"\} >> $FILE
done
