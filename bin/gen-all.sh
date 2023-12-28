dir=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
$dir/serialize-generator $GOFILE
$dir/mapper-generator $GOFILE