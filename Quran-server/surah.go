package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"text/template"

	"github.com/gin-gonic/gin"
)

/*
verses
	Translations

text_uthmani
verse_key
id
*/

func getSurah(c *gin.Context) {
	idx := c.Param("id")

	surahArabic, err := os.Open("static/json/arabic/" + idx + ".json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer surahArabic.Close()

	surahInfo, err := os.Open("static/json/chapters/" + idx + ".json")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer surahInfo.Close()

	surahEnlish, err := os.Open("static/json/english/" + idx + ".json")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer surahEnlish.Close()

	surahBangla, err := os.Open("static/json/bangla/" + idx + ".json")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer surahBangla.Close()

	var all CompleteSurah
	// ara
	if err = json.NewDecoder(surahArabic).Decode(&all.Aarabic); err != nil {
		fmt.Println(err)
		return
	}

	// info
	if err = json.NewDecoder(surahInfo).Decode(&all.SurahInfo); err != nil {
		fmt.Println(err)
		return
	}

	// trans
	all.Tranlaions = append(all.Tranlaions, TranslatedVerses{})
	if err = json.NewDecoder(surahEnlish).Decode(&all.Tranlaions[0]); err != nil {
		fmt.Println(err)
		return
	}

	all.Tranlaions = append(all.Tranlaions, TranslatedVerses{})
	if err = json.NewDecoder(surahBangla).Decode(&all.Tranlaions[1]); err != nil {
		fmt.Println(err)
		return
	}

	surahTemplate, err := template.ParseFiles("static/html/surah.html", "static/css/sura-s.css", "static/html/common.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	surahBuff := new(bytes.Buffer)
	surahTemplate.Execute(surahBuff, all)
	c.Writer.Write(surahBuff.Bytes())
}
