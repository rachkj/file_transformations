#!/bin/bash

# find the difference between the two files to understand how to reformat it.
diff ../file1 /Users/kj/file2 | grep "<" | sed 's/^<//g' > "diff_file"

# sort file.
sort ../file1 > temp.txt

# remove lines with certain words.
sed -i '' '/SOA/d' temp.txt
sed -i '' '/NS/d' temp.txt
sed -i '' '/10800/d' temp.txt
sed -i '' '/cli1db-scan/d' temp.txt
sed -i '' '/solaris/d' temp.txt
sed -i '' '/t1/d' temp.txt

# remove unnecessary lines.
sed -i "" "1,9d" temp.txt

# reverse the lines
awk '{
  for (i = NF; i >= 1; i--) {
    printf "%s", $i
    if (i > 1) {
      printf " "
    }
  }
  printf "\n"
}' temp.txt > file3

# Uncomment
rm diff_file temp.txt file3
