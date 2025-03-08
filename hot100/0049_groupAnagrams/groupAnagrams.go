package groupanagrams

import "strconv"

func groupAnagrams(strs []string) [][]string {
	key2words := make(map[string][]string)
	for _, str := range strs {
		index := make([]int, 26)
		for _, c := range str {
			index[c-'a'] += 1
		}
		key := getKey(index)
		if _, ok := key2words[key]; !ok {
			key2words[key] = make([]string, 0)
		}
		key2words[key] = append(key2words[key], str)
	}

	rtn := make([][]string, 0)
	for _, v := range key2words {
		rtn = append(rtn, v)
	}

	return rtn
}

func getKey(index []int) string {
	key := ""
	for i := 0; i < len(index); i++ {
		if index[i] != 0 {
			key += string(byte('a'+i)) + strconv.Itoa(index[i])
		}
	}
	return key
}
