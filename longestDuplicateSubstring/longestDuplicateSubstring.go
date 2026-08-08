package main

import "fmt"

func longestDupSubstring(s string) string {
	// Binary search for length of the substring
	// since we're searching for "minimum value satisfying the condition",
	// left will be "minimum length where repeating substring does NOT exist"

	left := 1       // even substrings of length 1 can be non-repeating, if all chars in the string are unique
	right := len(s) // Since we're searching for "minimum NOT repeating", we add 1. The sliding window for the repeating substring can have the size up to len(s) - 1.

	for left < right {
		mid := (left + right) / 2
		//fmt.Printf("left: %v, right: %v, mid: %v \n", left, right, mid)

		repeatingSubstringIndex := getRepeatingSubstringIndex_RabinKarp(s, mid) // Rabin-Karp
		//repeatingSubstringIndex := getRepeatingSubstringIndex(s, mid) // brute-force

		//fmt.Printf("Index of longest repeating substring of len %v: %v \n", mid, repeatingSubstringIndex)

		if repeatingSubstringIndex == -1 { // target condition
			right = mid
		} else {
			left = mid + 1
		}
	}

	if left == 1 { // no repeating substrings even with length 1, i.e.
		return ""
	}

	firstIndex := getRepeatingSubstringIndex_RabinKarp(s, left-1) // Rabin-Karp
	//firstIndex := getRepeatingSubstringIndex(s, left-1) // brute-force

	return s[firstIndex : firstIndex+left-1]
}

// todo: implement this method using the Rabin-Karp algorithm
func getRepeatingSubstringIndex(s string, substringLen int) int { // returns first index of repeating substring or -1
	// this is brute-force

	// substring to first index
	m := make(map[string]int)

	for i := 0; i <= len(s)-substringLen; i++ {
		substring := s[i : i+substringLen]

		if firstIndex, ok := m[substring]; ok { // substring already present -> return the index
			return firstIndex
		}

		// put first index of the new substring to the map
		m[substring] = i
	}

	return -1
}

func getRepeatingSubstringIndex_RabinKarp(s string, substringLen int) int { // returns first index of repeating substring or -1
	// Rabin-Karp algorithm - sliding window adopted to check whether there are substrings with the same hash.

	// todo: power pre-calculations can be done once for all the lengths up to ( len(s) - 1 ) and then re-used
	a := 26              // base of power - 26 since we have just lowercase English characters
	mod := 1_000_000_007 // all hashes (and big power calculations) are % mod. This should be a big prime number, to minimize the hash collisions

	// c[i] - characters of the substring

	// n - string length
	// hash = c[0] * a^(n-1) + c[1] * a^(n-2) + ... + c[l-2]*a^1 + c[l-1]*a^0
	// !!! to avoid epic overflows on a^x, we always mod the intermediate results by mod

	// first, we calculate the hash for the substring starting with s[0]
	hash := 0

	// precalc a^n powers for [0..l-1], every value % mod
	powers := make([]int, substringLen)

	aPower := 1
	for i := range substringLen {
		powers[i] = aPower
		aPower = (aPower * a) % mod
	}

	//fmt.Printf("Powers of %v mod %v: %v \n", a, mod, powers)

	for i := range substringLen {
		ch := charToInt(s[i])

		n := substringLen - i - 1 // power of A. Len = 4, i = 0, c[0]*a^3, 3 = (len - i - 1)

		aPower = powers[n]
		//fmt.Printf("%v^%v mod %v = %v \n", a, n, mod, aPower)

		hash += (ch * aPower % mod) % mod
		hash = hash % mod // !!! += operation is adding a small value, but the sum should also be modded!
	}

	//fmt.Printf("Hash for first substring [%v:%v] = \"%v\" of length %v = %v \n", 0, substringLen, s[0:substringLen], substringLen, hash)

	// map of hashes to index start
	// !!! Since we expect that collisions can happen, we put multiple indexes for the same hash
	m := make(map[int][]int)
	m[hash] = []int{0} // first hash is for substring starting at index 0

	for i := 1; i <= len(s)-substringLen; i++ {
		// we calculate the next hash from previous hash

		// we subtract c[oldFirst]*a^(n-1) from oldHash
		// then all the values should be multiplied by a (because of their powers should now be increased)
		// and then we add the last element (for last character of the new substring) that is c[l-1]*a^0
		prevFirstChar := s[i-1]
		lastChar := s[i+substringLen-1]
		//fmt.Printf("Prev first char: %c, next last char: %c \n", prevFirstChar, lastChar)

		toSubtract := (charToInt(prevFirstChar) * powers[substringLen-1]) % mod // c[oldFirst}*a^(n-1)

		// !!! + mod is to avoid negative values when (hash < toSubtract)
		hash = (hash - toSubtract + mod) % mod // oldHash - ( c[oldFirst] * a^(n-1) )
		if hash < 0 {
			fmt.Printf("!!! Hash is negative (1): %v \n", hash)
		}

		hash = (a * hash) % mod // a * ( oldHash - ( c[oldFirst] * a^(n-1) ) )
		if hash < 0 {
			fmt.Printf("!!! Hash is negative (2): %v \n", hash)
		}

		hash = (hash + charToInt(lastChar)*powers[0]) % mod // a * ( oldHash - ( c[oldFirst] * a^(n-1) ) ) + c[newLastChar] * a^0, a^0 = 1
		if hash < 0 {
			fmt.Printf("Hash is negative (3): %v \n", hash)
		}

		//fmt.Printf("Hash for substring[%v:%v] = \"%v\": %v \n", i, i+substringLen-1, s[i:i+substringLen], hash)

		if (m[hash] != nil) && (len(m[hash]) > 0) { // hash is already present
			//fmt.Printf("Hash %v is already present in the map: m[%v] = %v \n", hash, hash, m[hash])

			// for every starting index for this hash, check whether this is the same substring as the current substring.
			// we're doing these checks to avoid different substrings with the same hash
			for _, startIndex := range m[hash] {
				// todo: maybe comparison char-by-char is more effective?
				existingSubstring := s[startIndex : startIndex+substringLen]
				currentSubstring := s[i : i+substringLen]

				if existingSubstring == currentSubstring {
					//fmt.Printf("Substring [%v:%v] = \"%v\" is the same as the current substring [%v:%v] = \"%v\". Returning startIndex = %v. \n", startIndex, startIndex+substringLen, existingSubstring, i, i+substringLen, currentSubstring, startIndex)
					return startIndex
				}
			}

			// no same substring found -> add the current index to the list of strings with the same hash
			m[hash] = append(m[hash], i)

		} else { // hash is not yet present -> add index i for this  hash
			m[hash] = []int{i}
		}
	}

	return -1
}

func charToInt(r byte) int {
	return int(r - 'a') // convert to values 0-25
}

func test(s string, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String: %v \n", s)

	result := longestDupSubstring(s)

	fmt.Printf("Longest repeating substring: %v \n", result)
	fmt.Printf("Expected result:             %v \n", expectedResult)
	fmt.Printf("Len(result):          %v \n", len(result))
	fmt.Printf("Len(expected result): %v \n", len(expectedResult))

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("banana", "ana")
}

func test2() {
	test("abcd", "")
}

func test3() {
	test("aa", "a")
}

func test4() {
	// repeating substrings can have length be up to (len(s) -1).
	test(
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	)
}

func test5() {
	// test-case 63/69, failing my Rabin-Karp implementation

	test(
		"okmzpmxzwjbfssktjtebhhxfphcxefhonkncnrumgduoaeltjvwqwydpdsrbxsgmcdxrthilniqxkqzuuqzqhlccmqcmccfqddncchadnthtxjruvwsmazlzhijygmtabbzelslebyrfpyyvcwnaiqkkzlyillxmkfggyfwgzhhvyzfvnltjfxskdarvugagmnrzomkhldgqtqnghsddgrjmuhpgkfcjkkkaywkzsikptkrvbnvuyamegwempuwfpaypmuhhpuqrufsgpiojhblbihbrpwxdxzolgqmzoyeblpvvrnbnsdnonhpmbrqissifpdavvscezqzclvukfgmrmbmmwvzfpxcgecyxneipexrzqgfwzdqeeqrugeiupukpveufmnceetilfsqjprcygitjefwgcvqlsxrasvxkifeasofcdvhvrpmxvjevupqtgqfgkqjmhtkyfsjkrdczmnettzdxcqexenpxbsharuapjmdvmfygeytyqfcqigrovhzbxqxidjzxfbrlpjxibtbndgubwgihdzwoywqxegvxvdgaoarlauurxpwmxqjkidwmfuuhcqtljsvruinflvkyiiuwiiveplnxlviszwkjrvyxijqrulchzkerbdyrdhecyhscuojbecgokythwwdulgnfwvdptzdvgamoublzxdxsogqpunbtoixfnkgbdrgknvcydmphuaxqpsofmylyijpzhbqsxryqusjnqfikvoikwthrmdwrwqzrdmlugfglmlngjhpspvnfddqsvrajvielokmzpmxzwjbfssktjtebhhxfphcxefhonkncnrumgduoaeltjvwqwydpdsrbxsgmcdxrthilniqxkqzuuqzqhlccmqcmccfqddncchadnthtxjruvwsmazlzhijygmtabbzelslebyrfpyyvcwnaiqkkzlyillxmkfggyfwgzhhvyzfvnltjfxskdarvugagmnrzomkhldgqtqnghsddgrjmuhpgkfcjkkkaywkzsikptkrvbnvuyamegwempuwfpaypmuhhpuqrufsgpiojhblbihbrpwxdxzolgqmzoyeblpvvrnbnsdnonhpmbrqissifpdavvscezqzclvukfgmrmbmmwvzfpxcgecyxneipexrzqgfwzdqeeqrugeiupukpveufmnceetilfsqjprcygitjefwgcvqlsxrasvxkifeasofcdvhvrpmxvjevupqtgqfgkqjmhtkyfsjkrdczmnettzdxcqexenpxbsharuapjmdvmfygeytyqfcqigrovhzbxqxidjzxfbrlpjxibtbndgubwgihdzwoywqxegvxvdgaoarlauurxpwmxqjkidwmfuuhcqtljsvruinflvkyiiuwiiveplnxlviszwkjrvyxijqrulchzkerbdyrdhecyhscuojbecgokythwwdulgnfwvdptzdvgamoublzxdxsogqpunbtoixfnkgbdrgknvcydmphuaxqpsofmylyijpzhbqsxryqusjnqfikvoikwthrmdwrwqzrdmlugfglmlngjhpspvnfddqsvrajviel",
		"okmzpmxzwjbfssktjtebhhxfphcxefhonkncnrumgduoaeltjvwqwydpdsrbxsgmcdxrthilniqxkqzuuqzqhlccmqcmccfqddncchadnthtxjruvwsmazlzhijygmtabbzelslebyrfpyyvcwnaiqkkzlyillxmkfggyfwgzhhvyzfvnltjfxskdarvugagmnrzomkhldgqtqnghsddgrjmuhpgkfcjkkkaywkzsikptkrvbnvuyamegwempuwfpaypmuhhpuqrufsgpiojhblbihbrpwxdxzolgqmzoyeblpvvrnbnsdnonhpmbrqissifpdavvscezqzclvukfgmrmbmmwvzfpxcgecyxneipexrzqgfwzdqeeqrugeiupukpveufmnceetilfsqjprcygitjefwgcvqlsxrasvxkifeasofcdvhvrpmxvjevupqtgqfgkqjmhtkyfsjkrdczmnettzdxcqexenpxbsharuapjmdvmfygeytyqfcqigrovhzbxqxidjzxfbrlpjxibtbndgubwgihdzwoywqxegvxvdgaoarlauurxpwmxqjkidwmfuuhcqtljsvruinflvkyiiuwiiveplnxlviszwkjrvyxijqrulchzkerbdyrdhecyhscuojbecgokythwwdulgnfwvdptzdvgamoublzxdxsogqpunbtoixfnkgbdrgknvcydmphuaxqpsofmylyijpzhbqsxryqusjnqfikvoikwthrmdwrwqzrdmlugfglmlngjhpspvnfddqsvrajviel",
	)
}

func main() {
	// 1044. Longest Duplicate Substring

	// to study Rabin-Karp algorithm -> find matching substrings when sliding window on the parent string
	test1()
	test2()
	test3()
	test4()
	test5()
}
