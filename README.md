# go-lang-tutorial
## Go 튜토리얼 돌려보기
### 환경
- OS : Mac Catalina
- Language : Go v1.14.2 
- Editor : Goland
### 환경설정 
- Go 설치방법 : [클릭](https://zorba91.tistory.com/30)
- 환경설정하기
1. 
```shell
  # 환경설정 파일 열기
  $ vim ~/.bash_profile
```
2. 
```shell
# GO ENV
# 작업할 path
export GOPATH="자신이 작업할 디렉토리 위치/go-lang-tutorial"
# 작업할 프로젝트의 bin 위치
export GOBIN=$GOPATH/bin

export PATH=$PATH:$GOPATH:$GOBIN
```
3. 
```shell
  # 변경한 환경설정 적용
  $ source ~/.bash_profile
```

### Godoc 설치
```shell
  $ go get -v  golang.org/x/tools/cmd/godoc
  # 설치 후 
  $ godoc -http=:8080
```
**localhost:8080**으로 접속하면 go 다큐먼트를 볼 수 있다.

---
참고 강의: 인프런 - 쉽고 빠르게 끝내는 GO언어 프로그래밍 핵심 기초 입문 과정
