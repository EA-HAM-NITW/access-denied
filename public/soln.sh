#!/bin/bash

# Check for required argument
if [ "$#" -ne 1 ]; then
  echo "Usage: $0 <html-file>"
  exit 1
fi

html_file="$1"

# Search for a <p> tag that contains "trichy" (case-insensitive) and extract its id attribute.
# This command works on the assumption that the entire <p> tag is on one line.
id_value=$(grep -i '<p' "$html_file" | grep -i 'trichy' | sed -n 's/.*<p[^>]*id="\([^"]*\)".*/\1/p')

if [ -n "$id_value" ]; then
  echo "Found id: $id_value"
else
  echo "No <p> tag containing 'trichy' with an id attribute was found."
fi
