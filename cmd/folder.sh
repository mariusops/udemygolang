#!/bin/bash

# create folders and files in the current directory, starting witht the name 143 for the first folder. Ending at 300.
for i in {143..300}
do
  mkdir $i
  touch $i/$i.go
done