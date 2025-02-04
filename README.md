# Friday.go

간단한 웹 페이지 저장을 위한 WEB Application. 

## Features

- sign-up & sign-in
- 웹 사이트의 호스트를 기준으로 URL 정보를 저장하는 기능
  - 간단한 별칭과 설명을 추가할 수 있음.(즐겨찾기와 유사) 
  - 최대 1장의 이미지 파일을 업로드하여 보다 직관적으로 등록한 사이트를 식별할 수 있음 
- 조회 및 검색
  - 키워드 검색 기능
    - 사이트 별  
  - 분류별 조회 기능
    - 사이트 별 조회
    - 태그 별 조회
## Language and Framework

### Backend
- Golang 1.23
- gofiber v2 (web framework)
- gorm 1.25 (ORM)
  - sqlite3 (DB)

### Frontend

- Node.js v22
- React for Typescript 18 
- MUI v6
- Vite v6