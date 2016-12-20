package day20

import (
	"fmt"
	"math"
	"strings"

	"strconv"

	"github.com/gsmcwhirter/advent2016/lib"
)

type ExclusionNode struct {
	NodeBlocked  bool
	BreakPoint   uint32
	BlockedLeft  bool
	BlockedRight bool
	Left         *ExclusionNode
	Right        *ExclusionNode
}

func NewFirewall() *ExclusionNode {
	return NewNode(math.MaxUint32 / 2)
}

func NewNode(breakPoint uint32) *ExclusionNode {
	node := &ExclusionNode{
		BreakPoint:   breakPoint,
		NodeBlocked:  false,
		BlockedLeft:  false,
		BlockedRight: false,
		Left:         nil,
		Right:        nil,
	}

	return node
}

func (n *ExclusionNode) ApplyExclusion(start uint32, stop uint32, min uint32, max uint32) bool {
	// fmt.Printf("Applying %v-%v to %v (min=%v, max=%v)\n", start, stop, n.BreakPoint, min, max)

	if start <= n.BreakPoint && stop >= n.BreakPoint {
		n.NodeBlocked = true
	}

	rightStart := lib.UInt32Max(n.BreakPoint, start)
	leftStop := lib.UInt32Min(n.BreakPoint, stop)

	switch {
	case start <= min && stop >= n.BreakPoint:
		n.BlockedLeft = true
	case start < n.BreakPoint && !n.BlockedLeft && n.Left == nil:
		// n.Left = NewNode(start/2 + n.BreakPoint/2)
		n.Left = NewNode(start)
		// fmt.Println("Go left with create")
		n.BlockedLeft = n.Left.ApplyExclusion(start, leftStop, min, n.BreakPoint-1)
	case start < n.BreakPoint && !n.BlockedLeft && n.Left != nil:
		// fmt.Println("Go left no create")
		n.BlockedLeft = n.Left.ApplyExclusion(start, leftStop, min, n.BreakPoint-1)
	}

	switch {
	case start <= n.BreakPoint && stop >= max:
		n.BlockedRight = true
	case stop > n.BreakPoint && !n.BlockedRight && n.Right == nil:
		// n.Right = NewNode(stop/2 + n.BreakPoint/2 + 1)
		n.Right = NewNode(stop)
		// fmt.Println("Go right with create")
		n.BlockedRight = n.Right.ApplyExclusion(rightStart, stop, n.BreakPoint+1, max)

	case stop > n.BreakPoint && !n.BlockedRight && n.Right != nil:
		// fmt.Println("Go right no create")
		n.BlockedRight = n.Right.ApplyExclusion(rightStart, stop, n.BreakPoint+1, max)
	}

	return n.BlockedLeft && n.BlockedRight && n.NodeBlocked
}

func (n *ExclusionNode) LeastUnblocked() uint32 {
	fmt.Printf("Considering %v... ", n.BreakPoint)
	if n.BlockedLeft {
		fmt.Print("Left blocked, considering right... ")
		if n.Right != nil {
			if n.NodeBlocked {
				fmt.Printf("Returning max of my value + 1 (%v) and right's min.\n", n.BreakPoint+1)
				return lib.UInt32Max(n.BreakPoint+1, n.Right.LeastUnblocked())
			}

			fmt.Printf("Returning max of my value (%v) and right's min.\n", n.BreakPoint)
			return lib.UInt32Max(n.BreakPoint, n.Right.LeastUnblocked())
		}

		if n.NodeBlocked {
			fmt.Printf("Returning my value + 1 (%v).\n", n.BreakPoint+1)
			return n.BreakPoint + 1
		}

		fmt.Printf("Returning my value (%v).\n", n.BreakPoint)
		return n.BreakPoint
	}

	fmt.Print("Left not blocked... ")
	if n.Left != nil {
		fmt.Println("Returning left's value.")
		return n.Left.LeastUnblocked()
	}

	fmt.Println("Returning 0.")
	return 0
}

func (n *ExclusionNode) NumberAllowed(min, max uint32) uint32 {
	num := uint32(0)

	// if !n.NodeBlocked {
	// 	num++
	// }

	switch {
	case !n.BlockedLeft && n.Left != nil:
		num += n.Left.NumberAllowed(min, n.BreakPoint-1)
	case !n.BlockedLeft && n.Left == nil:
		num += n.BreakPoint - min
	}

	switch {
	case !n.BlockedRight && n.Right != nil:
		num += n.Right.NumberAllowed(n.BreakPoint+1, max)
	case !n.BlockedRight && n.Right == nil:
		num += max - n.BreakPoint
	}

	return num
}

func (n *ExclusionNode) Print(indent string) {
	var leftVal interface{}
	var rightVal interface{}

	if n.Left != nil {
		leftVal = n.Left.BreakPoint
	} else {
		leftVal = nil
	}

	if n.Right != nil {
		rightVal = n.Right.BreakPoint
	} else {
		rightVal = nil
	}

	fmt.Printf("%vBP: %v, B: %v, L: %v (blocked=%v), R: %v (blocked=%v)\n",
		indent,
		n.BreakPoint,
		n.NodeBlocked,
		leftVal,
		n.BlockedLeft,
		rightVal,
		n.BlockedRight)

	if n.Left != nil && !n.BlockedLeft {
		n.Left.Print(indent + "  ")
	}

	if n.Right != nil && !n.BlockedRight {
		n.Right.Print(indent + "  ")
	}
}

func ParseDataString(dat string) [][2]uint32 {
	lines := strings.Split(dat, "\n")
	exclusions := make([][2]uint32, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		stop, _ := strconv.Atoi(parts[1])

		exclusions[i] = [2]uint32{uint32(start), uint32(stop)}
	}

	return exclusions
}

func LoadData(filename string) [][2]uint32 {
	dat := lib.ReadFileData(filename)

	return ParseDataString(strings.Trim(string(dat), "\n"))
}

func RunPartA(filename string) {
	exclusions := LoadData(filename)
	firewall := NewFirewall()

	for _, ex := range exclusions {
		fmt.Printf("Excluding %v-%v\n", ex[0], ex[1])
		firewall.ApplyExclusion(ex[0], ex[1], 0, math.MaxUint32)
		// firewall.Print("")
		// fmt.Println(firewall.LeastUnblocked())
	}

	// firewall.Print("")
	fmt.Println(firewall.LeastUnblocked())
}

func RunPartB(filename string) {
	exclusions := LoadData(filename)
	firewall := NewFirewall()

	for _, ex := range exclusions {
		fmt.Printf("Excluding %v-%v\n", ex[0], ex[1])
		firewall.ApplyExclusion(ex[0], ex[1], 0, math.MaxUint32)
		// firewall.Print("")
		// fmt.Println(firewall.LeastUnblocked())
	}

	// firewall.Print("")
	fmt.Println(firewall.NumberAllowed(0, math.MaxUint32))
}
