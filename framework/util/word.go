package word

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

var caser = cases.Title(language.English)

func ToTitle(word string) string {
	return caser.String(word)
}

// ToTitleCamel 将下划线分割的字符串转换为驼峰字符串
// class_id => ClassId
func ToTitleCamel(word string) string {
	// 将字符串分割为单词，并使用 Title 函数将每个单词的首字母大写
	words := strings.Split(word, "_")
	for i, w := range words {
		words[i] = ToTitle(w)
	}
	// 将单词连接在一起，形成一个驼峰字符串
	camelCase := strings.Join(words, "")
	return camelCase
}

// ToNormalCamel 将下划线分割的字符串转换为驼峰字符串
// class_id => classId
func ToNormalCamel(word string) string {
	// 将字符串分割为单词，并使用 Title 函数将每个单词的首字母大写
	words := strings.Split(word, "_")
	for i, w := range words {
		if i == 0 {
			words[i] = strings.ToLower(w)
			continue
		}
		words[i] = ToTitle(w)
	}
	// 将单词连接在一起，形成一个驼峰字符串
	camelCase := strings.Join(words, "")
	return camelCase
}
