# tutorial 9 (module)

- 경로: TUTORIALS/09_module


---

- Rules
  - CMD
    - go mod tidy
      - 외부 모듈들을 다운로드 받는다.
      - 다운되는 경로는 go env 에서 GOPATH\pkg\mod
  - 초기화
    - func init() 에서 패키지 로딩시 초기화된다.
  - 경로
    - 같은 폴더 안의 pacakage는 모두 같은 이름의 package여야한다.
  - 공개 여부
    - 첫글자가 대문자면 공개 소문자면 비공개

- 