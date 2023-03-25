# tutorial 1 (Hello)

- 경로: TUTORIALS/01_hello
- run
  > go run hello.go
- Build 순서 (.exe 파일 생성)
  - mod 파일 생성 및 초기화 (manage dendencies through code's own module)
    > go mod init TUTORIALS/01_hello
  - build (as Windows executable file)
    > go build
  - build (as Linux executable file)
    > $Env:GOOS = "linux"
    > gp bio;d
