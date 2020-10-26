

PREFIX=/usr/local
ABSOLUTE_PATH="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
GONAME="$(basename "$(PWD)")"


function usage() {
    echo "Cheemit Installer For UNIX"
    echo
    echo "Syntax: $0 [-p]"
    echo
    echo "Options:"
    echo "-p | --prefix            Change the prefix where the software is going to be installed to (Default: /usr/local)"
    echo "-h | --help               Display this help menu"
    echo
}

if [ "$1" == "" ]; then
    usage
    exit 1
fi

while [ "$1" != "" ]; do
    case $1 in
        -p | --prefix )                shift
                                                PREFIX="$1"
                                                ;;
        -h | --help )                   usage
                                                exit
                                                ;;
        * )                                    usage
                                                exit 1
                                                ;;
    esac
    shift
done

if [ ! -d "$PREFIX/share/cheemit/image" ]; then
    echo "The folder to create the cheems doesn't exist, will attempt to create it"
    mkdir "$PREFIX/share/cheemit/image"
fi

for image in $ABSOLUTE_PATH/resources/images/*;do
    install -m 777 "$image" "$PREFIX/share/cheemit/image"
done

go build -ldflags="-X github.com/diegomagdaleno/cheemit/lib.prefix=${PREFIX}" -o "bin/$GONAME"

install -m 777 "./bin/$GONAME" "$PREFIX/bin/cheemit"

