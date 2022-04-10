package main

import (
	"fmt"
	"strconv"
)

// Color 终端颜色
type Color int

// 前背景色
const (
	FrontendBlack Color = iota + 30
	FrontendRed
	FrontendGreen
	FrontendYellow
	FrontendBlue
	FrontendWineRed
	FrontendUltramarine
	FrontendWhite
)

// 后背景色
const (
	BackendBlack Color = iota + 40
	BackendRed
	BackendGreen
	BackendYellow
	BackendBlue
	BackendWineRed
	BackendUltramarine
	BackendWhite
)

// String impl fmt.Stringer
func (c Color) String() string {
	return strconv.Itoa(int(c))
}

// 预设的几种字体颜色
var (
	GreenWord  = WordColor{FrontendGreen, BackendBlack}
	RedWord    = WordColor{FrontendRed, BackendBlack}
	YellowWord = WordColor{FrontendYellow, BackendBlack}
)

// ColorStringMap 字体颜色与字符串的映射
var ColorStringMap map[string]*WordColor = map[string]*WordColor{
	"green":  &GreenWord,
	"red":    &RedWord,
	"yellow": &YellowWord,
}

// WordColor 字体颜色
type WordColor struct {
	Frontend, Backend Color
}

// String impl fmt.Stringer
func (wc *WordColor) String() string {
	return wc.Frontend.String() + ";" + wc.Backend.String()
}

// GetColor 获取预设的字体颜色
func GetColor(color string) (*WordColor, error) {
	if c, ok := ColorStringMap[color]; ok {
		return c, nil
	}
	return nil, fmt.Errorf("color[%s] not found.", color)
}
