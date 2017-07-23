# NOTE: This should _NOT_ be used in production!
# NOTE: This file should be deleted to prevent data loss.

echo "NOTICE: All journal entires will be deleted!"
echo -n "To continue, please type continue: "

read reply
if [ "$reply" == "continue" ]; then
  echo "Deleting all journal entires and folders"
  rm -rf ./entries
  echo "Done"
fi
