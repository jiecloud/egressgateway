// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package tools

import (
	"context"
	"fmt"
	"math/rand"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/spidernet-io/egressgateway/test/e2e/err"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateString(lenNum int, isHex bool) string {
	var chars []string
	chars = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	if isHex {
		chars = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	}
	str := strings.Builder{}
	length := len(chars)
	for i := 0; i < lenNum; i++ {
		str.WriteString(chars[r.Intn(length)])
	}
	return str.String()
}

func GenerateRandomNumber(max int) string {
	return strconv.Itoa(r.Intn(max))
}

// SubtractionSlice  a, b are inclusion relationship
func SubtractionSlice(a, b []string) []string {
	sort.Strings(a)
	sort.Strings(b)
	if len(a) > len(b) {
		a, b = b, a
	}
	mapa := make(map[string]struct{}, len(a))
	var result []string

	for i := range a {
		mapa[a[i]] = struct{}{}
	}
	for _, v := range b {
		if _, ok := mapa[v]; !ok {
			result = append(result, v)
		}
	}
	return result
}

// IsSameSlice determine whether two slices are the same
func IsSameSlice(a, b []string) bool {
	sort.Strings(a)
	sort.Strings(b)
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// ExecInKindNode exec command in kind node
func ExecInKindNode(nodeName string, command string, duration time.Duration) ([]byte, error) {
	if len(nodeName) == 0 || len(command) == 0 {
		return nil, err.EMPTY_INPUT
	}
	ctx, cancel := context.WithTimeout(context.TODO(), duration)
	defer cancel()
	c := fmt.Sprintf("docker exec %s %s", nodeName, command)
	return exec.CommandContext(ctx, "sh", "-c", c).Output()
}

func ExecCommand(command string, duration time.Duration) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), duration)
	defer cancel()
	return exec.CommandContext(ctx, "sh", "-c", command).Output()
}

func GetNetStats(duration time.Duration) ([]byte, error) {
	a := "ss -tunlp "
	return ExecCommand(a, duration)
}

func GetKernelParams(duration time.Duration) ([]byte, error) {
	a := "sysctl -a "
	return ExecCommand(a, duration)
}

func GetContainerIPV4(container string, duration time.Duration) ([]byte, error) {
	a := fmt.Sprintf("docker inspect %s | grep -w IPAddress | grep -E -o '[0-9]+\\.[0-9]+\\.[0-9]+\\.[0-9]+' | tr -d '\n'", container)
	return ExecCommand(a, duration)
}

func GetContainerIPV6(container string, duration time.Duration) ([]byte, error) {
	a := fmt.Sprintf("docker inspect %s | grep -w GlobalIPv6Address  | sed 1d | awk '{print $2}' | tr -d '\",' | tr -d '\n'", container)
	return ExecCommand(a, duration)
}
