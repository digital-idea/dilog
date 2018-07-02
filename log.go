package main

type Log struct {
	Cip     string // Client IP
	Port    string // Client Port
	Id      string // log ID
	Keep    string // 보관일수
	Log     string // 로그내용
	Project string // 프로젝트
	Slug    string // CSI같은 툴에서 사용하는 Slug
	Time    string // 로그가 기입된 시간
	Tool    string // 로그가 보내진 툴
	User    string // 사번
}