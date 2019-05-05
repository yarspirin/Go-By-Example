TOPIC="$1"
EXT=".go"
SRC="$TOPIC$EXT"

mkdir $TOPIC
cd $TOPIC
touch $SRC
echo -e "package main\n\nfunc main() {\n\n}\n" > $SRC
echo "$TOPIC folder has been generated."
