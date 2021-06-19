package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// lotto.exe filename count
func main() {

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "invalid arguments.! filename count")
		return
	}

	fileName := os.Args[1]
	count, err := strconv.Atoi(os.Args[2]) // strconv.Atoi() 스트링을 숫자로 변화
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot conver to integer !!!", count)
		return
	}

	candidates, err := readFileCandidates(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot read file", err)
		return
	}

	rand.Seed(time.Now().UnixNano())

	// 뽑을 숫자만큼 루핑을 돌려서 winners 만든다.
	winners := make([]string, count)
	for i := 0; i < count; i++ {
		n := rand.Intn(len(candidates)) // 뽑힌 후보자의 순서 == n
		winners[i] = candidates[n]      // winners에는 뽑는 숫자만큼 저장된다.
		// 이번에는 뽑힌 사람을 제거하고 candidates 관리한다.
		// 현재 뽑힌사람은 n번째, 그럼 후보자는 1부터 n-1번째까지, n+1부터 끝까지가 후보자가 된다.
		candidates = append(candidates[:n], candidates[n+1:]...) //:n 1부터 n-1까지, n+1: ... n+1 부터 끝까지
		// 결국 n번째 것을 빼는 것이다.
	}

	// 뽑힌 사람을 출력해 보자
	fmt.Println("Winners are !!!")
	for _, winner := range winners {
		fmt.Println(winner)

	}
}

func readFileCandidates(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var candidates []string
	for scanner.Scan() {
		candidates = append(candidates, scanner.Text())
	}

	return candidates, nil
}
