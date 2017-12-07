package medium

import (
	"bytes"
    "strings"
    "math/rand"
)

const (
    strMap = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type Solution struct {
    pre string
    size int
    short2long map[string]string
    long2short map[string]string
}

func InitSolution(pre string, size int) *Solution {
    sl := new(Solution)
    rand.Seed(time.Now().UnixNano())
    sl.short2long = make(map[string]string)
    sl.long2short = make(map[string]string)
    sl.pre = pre
    sl.size = sizes
    return sl
}

func (sl *Solution) Encode(longUrl string) string {
    if re, ok := sl.long2short[longUrl]; ok {
        return sl.pre + sl.long2short[longUrl]
    }

    key := sl.randomString(sl.size)
    sl.short2long[key] = longUrl
    sl.long2short[longUrl] = key

    return sl.pre + key
}

func (sl *Solution) randomString(size int) string {
    re := bytes.NewBufferString("")
    for {
        for i := 0; i < size; i++ {
            idx := rand.Intn(len(strMap))
            re.WriteString(strMap[idx])
        }
        if _, ok := sl.short2long[re.String()]; !ok {
            break
        }
    }
    return re.String()
}

func (sl *Solution) Decode(shortUrl string) string {
    key := strings.LastIndex(shortUrl, "/")
    
    if re, ok := sl.short2long[shortUrl[key+1:]]; ok {
        return re
    }
    return shortUrl
}