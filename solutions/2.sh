#!/bin/bash



html_file=index.html

# Search for a <p> tag that contains "trichy" (case-insensitive) and extract its id attribute.
# This command works on the assumption that the entire <p> tag is on one line.
id_value=$(grep -i '<p' "$html_file" | grep -i 'Istanbul' | sed -n 's/.*<p[^>]*id="\([^"]*\)".*/\1/p')

if [ -n "$id_value" ]; then
  echo "$id_value"
else
  echo "No <p> tag containing 'Istanbul' with an id attribute was found."
fi