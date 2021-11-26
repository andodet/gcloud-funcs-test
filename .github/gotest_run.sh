#!/bin/sh

for i in $(ls); do
  # If it's a folder
  if [[ -d "$i" ]]; then
      cd $i
      is_go=0
      for f in $(ls); do
          # Check if contains python files
          if [[ $f == *.go ]]; then
          is_go=$((is_go+1))
          fi
      done
      # Run tests if it does
      if ((is_go == 0)); then
          cd .. 
          continue
      else
          go test ./...
          cd ..
      fi
  fi
done
