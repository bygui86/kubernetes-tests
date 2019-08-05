#!/bin/sh

SRC_PATTERN="$1"
REPL_PATTERN="$2"
BAK_SUFFIX="$3"
DIR="$4"

if [[ -z "$4" ]]; then
	echo
	echo "Missing input parameters:"
	echo "    \$1  Source pattern"
	echo "    \$2  Replacement pattern"
	echo "    \$3  Backup suffix"
	echo "    \$4  Directory"
	echo
	echo "For example:"
	echo "    \$1  (.*)image: (.*)"
	echo "    \$2  \1image: eas-docker-virtual.artifactory.swisscom.com\/\2"
	echo "    \$3  _bak"
	echo "    \$4  ."
	echo '    ./kubernetes-yaml-image-replacer.sh "(.*)image: (.*)" "\1image: eas-docker-virtual.artifactory.swisscom.com\/\2" "_bak" "."'
	echo
	echo "Please retry..."
	echo
	exit -1
fi

# samples
# SRC_PATTERN="(.*)image: (.*)"
# REPL_PATTERN="\1image: eas-docker-virtual.artifactory.swisscom.com\/\2"
# BAK_SUFFIX="_bak"

echo
echo "Parse $DIR to add replace Docker images definitions"

for filepath in $(find $DIR -type f -regex ".*.yaml"); do
	# if [[ $filepath != *"DS_Store"* ]]; then
		# filename=$(basename $filepath)
		# echo "    ... parsing $filename"
		echo "    ... parsing $filepath"
		sed -i "$BAK_SUFFIX" -E "s/$SRC_PATTERN/$REPL_PATTERN/g" "$filepath"
	# fi
done

echo
echo "Parse $DIR completed"
echo
