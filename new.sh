year=$1
day=$2

dir="$1/$2"

if [ -d $dir ]; then
  echo "Solution already exists for $year/$day"
  exit 0
fi

mkdir -p $dir

cat << EOF > $dir/main.go
package main

import (
	"aoc_go/utils"
	"fmt"
	"time"
)

func part1(input string) string {
	return ""
}

func part2(input string) string {
	return ""
}

func main() {
	input := utils.ReadInput()

	part1_start := time.Now()
	part1_result := part1(input)
	part1_duration := time.Since(part1_start)

	part2_start := time.Now()
	part2_result := part2(input)
	part2_duration := time.Since(part2_start)

	fmt.Println()
	fmt.Printf("Part 1: %s (%v)", part1_result, part1_duration)
	fmt.Println()
	fmt.Printf("Part 2: %s (%v)", part2_result, part2_duration)
}
EOF

cat << EOF > $dir/main_test.go
package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
        input := strings.TrimSpace(\`\`)
        expected := strings.TrimSpace(\`\`)

	result := part1(input)

	if result != expected {
		t.Errorf("\n\nPart 1:\n  Expected: %q\n  Received: %q\n\n", expected, result)
	}
}

func TestPart2(t *testing.T) {
        input := strings.TrimSpace(\`\`)
        expected := strings.TrimSpace(\`\`)

	result := part2(input)

	if result != expected {
		t.Errorf("\n\nPart 2:\n  Expected: %q\n  Received: %q\n\n", expected, result)
	}
}
EOF

touch $dir/input.txt

echo "Created boilerplate code for $dir"
