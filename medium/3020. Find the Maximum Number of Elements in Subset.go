package medium

func maximumLength260627(nums []int) int {
	countDict := map[int]int{}
    ansDict := map[int]int{}
    ans := 0
    for _, num := range nums {
        countDict[num] += 1
    }
    if countDict[1] > 0 {
        if countDict[1] % 2 == 0 {
            ans = countDict[1] - 1
        } else {
            ans = countDict[1]
        }
        delete(countDict, 1)
    }
    for base, bCount := range countDict {
        if bCount == 1 {
            ansDict[base] = 1
        } else {
            curr, currCount := base, bCount
            currAns := 2
            for currCount >= 2 {
                next := curr * curr
                if nextCount, exist := ansDict[next]; exist {
                    ansDict[base] = nextCount + currAns
                    break
                }
                if nextCount, exist := countDict[next]; exist {
                    if nextCount == 1 {
                        ansDict[base] = currAns + 1
                        break
                    } else {
                        currAns += 2
                        curr, currCount = next, nextCount
                    }
                } else {
                    ansDict[base] = currAns - 1
                    break
                }
                // fmt.Println(currAns, curr, currCount)
            }
        }
        // fmt.Println(base, bCount, ansDict)
        ans = max(ans, ansDict[base])
    }
    return ans
}