# 실무에 바로 쓰는 Go언어 핸즈온 가이드

## 커맨드 라인 애플리케이션

Main 함수를 테스트하는 방법도 있는데 그땐 `os.Exec()` 를 사용하여 지정해주도록하자.  
처음 내용으로는 좀 어렵지만 쓸만한 키워드를 많이 배울수있다.  
기본적으로 전체 테스트를 돌릴때는 `go test -v` 를  쓰고  
코드커버리지를 보고싶을때는 `go test -coverprofile cover.out` 를 쓴다.  
마지막으로 `html` 형태로 보고싶을땐 `go tool cover -html=cover.out` 를 사용하면된다.  
