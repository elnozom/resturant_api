#!/bin/sh
FILES="./im/*"
for f in $FILES
do
  filename=$(basename -- "$f")
  echo $filename
done



remove_word() (
  set -f
  IFS=' '

  s=$1
  w=$2

  set -- $1
  for arg do
    shift
    [ "$arg" = "$w" ] && continue
    set -- "$@" "$arg"
  done

  printf '%s\n' "$*"
)