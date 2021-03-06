# dilog

![travisCI](https://secure.travis-ci.org/digital-idea/dilog.png)

![screenshot](figures/screenshot01.png)

웹용 로그서버 입니다.
간혹 조직에서 만들어지는 툴에 대해서 로그 기록이 필요한 경우가 있습니다.
프로젝트 매니징툴의 로그를 기록하기위해 만들었지만, 다른툴에서도 활용할 수 있도록 제작했습니다.

### Download
오직 실행파일 하나면 됩니다.
하나의 실행파일로 클라이언트 명령어, 웹서버 및 restAPI 기능을 사용할 수 있습니다.

- [Linux 64bit](https://github.com/digital-idea/dilog/releases/download/v1.0.5/dilog_linux_x86-64.tgz)
- [Linux 64bit Digitalidea](https://github.com/digital-idea/dilog/releases/download/v1.0.5/dilog_linux_di_x86-64.tgz)
- [macOS 64bit](https://github.com/digital-idea/dilog/releases/download/v1.0.5/dilog_darwin_x86-64.tgz)
- [macOS 64bit Digitalidea](https://github.com/digital-idea/dilog/releases/download/v1.0.5/dilog_darwin_di_x86-64.tgz)
- [Windows 64bit](https://github.com/digital-idea/dilog/releases/download/v1.0.5/dilog_windows_x86-64.tgz)
- [Windows 64bit Digitalidea](https://github.com/digital-idea/dilog/releases/download/v1.0.5/dilog_windows_di_x86-64.tgz)

### mongoDB 설치, 실행
이곳에서 mongoDB에 대해서 자세히 다루지 않습니다.

monogoDB를 설치합니다.

```bash
$ sudo yum install mongodb mongodb-server // CentOS
$ brew install mongodb // macOS
```

monogoDB 실행

```
# mongod
$ brew services start mongodb // macOS
```

### 로그서버 실행
실행파일로 서버를 실행할 수 있습니다.
```
# dilog -http :80
```

mongodb 서버 IP가 10.0.20.30 이고 dilog를 :8080 포트로 실행하는 방법은 아래와 같습니다.

```bash
$ dilog -http :8080 -dbip 10.0.20.30
```


### 로그추가: Commandline
기본적으로 툴이름과 로그내용만 작성하더라도 로그를 기록할 수 있습니다.

```bash
$ dilog -tool 툴이름 -log="로그내용"
```

다른 인수들을 활용하여 로그를 처리할 수 있습니다.

> 예시 : toolA 에서 circle 프로젝트 SS_0010_org 샷에 "A버튼을 눌렀다." 라는 로그를 365일 보관

```bash
$ dilog -tool tooA -project circle -slug SS_0010_org -log "A버튼을 눌렀다." -keep 365
```

### 로그추가: RestAPI
- curl을 이용해서 로그를 POST하는 방법

```bash
$ curl -X POST -d "project=circle&slug=SS_0010_org&tool=csi3&user=woong&keep=180&log=log_text" http://127.0.0.1:8080/api/setlog
```

### 로그삭제: Commandline
아래 명령어를 통해서 시간이 지난 로그를 일괄 삭제할 수 있습니다.

```bash
$ dilog -rm
```

id를 이용해서 하나의 로그를 삭제할 수 있습니다.

```
$ dilog -rmid 0000000000000
```

### 로그접근
아래 형태의 url을 통해서 로그서버에 접근할 수 있습니다.

```
http://127.0.0.1/search?project=circle&slug=SS_0010
```

검색: daily, 2019-06-10
```
http://127.0.0.1/search?searchword=daily+2019-06-10
```

## 응용
여러 응용법이 있습니다. 아래처럼 로그를 입력합니다.
circle 프로젝트에 SS_0010_org 샷의 데일리 로그라고 생각해보겠습니다.

```bash
$ dilog -tool terminal -project circle -slug SS_0010_org -log "daily: ~/project/circle/test.mov"
```

데이터 입력후 웹에서 검색을 할 수 있습니다. 로그내용에 경로가 있다면 [dilink 프로토콜](https://github.com/digital-idea/dilink)과 연동되며 응용프로그램과 물릴 수 있습니다.
![screenshot](figures/ex_daily.png)

## Developer

### Need

빌드를 위해서 아래 라이브러리를 사용합니다.

```
$ go get github.com/GeertJohan/go.rice
$ go get github.com/GeertJohan/go.rice/rice
```

- Maintainer: HanwoongKim(hello@lazypic.org)
- Committer: 
- Contributors: 

## License
BSD 3-Clause License