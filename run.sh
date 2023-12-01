year=$1
day=$2

dir="$1/$2"

if [ ! -d $dir ]; then
  echo "No solution found for $year/$day"
  exit 1
fi

cd $dir
go run main.go
