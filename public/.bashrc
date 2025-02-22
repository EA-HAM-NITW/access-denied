alias check="~/cli check"

reveal() {
  if [ "$#" -eq 0 ]; then
    ~/reveal -P readme.md
  else
    ~/reveal -P "$@"
  fi
}

info() {
  echo -e "Hello $USER! Congratulations for hacking into the Phantom's machine! But this is not it, the real challenge begins now.\n\nInstructions\n\n1. Go through each folder inside 404 from  your root directory\n2. Type \`reveal\` to know more about the puzzle\n3. Find out how to get the answer with other resources in the same directory, type \`ls\` to see the other clues\n4. Write your answers in \`script.sh\` in the respective folder.\n5. Run \`check\` once you are done typing your answer.\n\nAll the best\n\nREMEMBER!\nIf you are stuck in any part of the puzzle you can ask my minions to help,\nBut beware!!! They will make sure to waste your time before giving you any sort of help\n\nYour team ID is $TEAM_ID\n\nTo re-read the instructions, run \`info\`"
}

info()

