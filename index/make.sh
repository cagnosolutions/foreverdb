#!/usr/bin/env bash

function HELP {
	echo -e "used to create new b+tree package with custom key and value types"
	echo -e "usage $0 [args]"
	echo -e "args:"
	echo -e "  -f [name]\t--folder name"
	echo -e "  -k <type>\t--key type (ex: -k uint64)"
	echo -e "  -v <type>\t--val type (ex: -v []byte)\n"
	exit 1
}

if [ $# -lt 1 ]; then
   HELP
fi

FOLDER=
KEY=
VALUE=

while getopts ":f:k:v:h" OPT; do
	case $OPT in
		f)
			FOLDER=$OPTARG
			;;
		k)
			KEY=$OPTARG
			;;
		v)
			VALUE=$OPTARG
			;;
		h)
			HELP
			;;
		\?)
			echo "use ${BOLD}$SCRIPT -h${NORM} to see the help documentation."
     			exit 1 
			;;
		:)
			echo "option -$OPTARG requires an argument, use${BOLD}$SCRIPT -h${NORM} to see the help documentation."
			exit 1
			;;
	esac
done
shift $((OPTIND-1))

if [[ -z $FOLDER ]] || [[ -z $KEY ]] || [[ -z $VALUE ]]; then
     HELP
fi	

echo "::> creating \"$FOLDER\" with types $KEY and $VALUE"
mkdir $FOLDER; cp *.go $FOLDER; cd $FOLDER
sed -i "s/package index/package $FOLDER/g" *.go
sed -i "s|interface{}[^{]*/\*K\*/|$KEY|g" *.go
sed -i "s|interface{}[^{]*/\*V\*/|$VALUE|g" *.go
echo "::> checking for any compile errors..."
go build
cd -
echo "::> complete"
