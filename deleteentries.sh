echo "NOTICE: All journal entires will be deleted!"
echo -n "To continue, please type continue: "

read reply
if [ "$reply" == "continue" ]; then
  echo "Deleting all journal entires and folders"
  rm -rf ./entries
  echo "Done"
fi
