# Insighter Image Augment
이미지 증강을 위한 명령어 인터페이스를 사용 할 수 있습니다.

# 사용법
주의: `golang` 설치 필요
```bash
# go version 확인
go version

# bin 파일 build
go build -o imgaug main.go

# cli 실행
./imgaug help
```

# 데이터 처리 절차
> collect -> resize -> crop -> segment -> coco

`주의: ~/datalake 폴더 생성 필수`

0. 데이터 수집
  - `annotation-tool` 관리자 페이지 접속
  - `프로젝트 목록`에서 프로젝트 선택
  - `배포용 Json 파일 다운로드` 클릭하여 다운로드
  - `~/datalake/` 폴더에 압축해제
  - 압축 해제한 폴더명 변경(폴더명은 비엔나 코드를 따름. e.g. 010102)

1. resize
parentDirPath: `~/datalake/{압축해제 폴더 이름}`

```bash
./imgaug resize --parentDirs {parentDirPath} --size 416 create

# 생성된 augment 폴더 확인
# resize 된 이미지 확인
```

2. crop
parentDirPath: `~/datalake/{압축해제 폴더 이름}`
```bash
./imgaug crop --parentDirs {parentDirPath} --size 416 create

# augment 폴더 확인
# crop 된 이미지 확인
```

3. segment
parentDirPath: `~/datalake/{압축해제 폴더 이름}`
outputDir: `~/datalake/segemnt`
```bash
./imgaug segment --parentDirs {parentDirPath} --outputDir {outputDir} --sample create

# 생성된 segemnt 폴더 확인
```

4. coco
- parentDirPath: `~/datalake/segemnt/{데이터셋 분류}/{압축해제 폴더 이름}`  
  - 데이터셋 분류 - train, validation, test, sample
- category: 세분류에 해당하는 카테고리 이름(010101 - 별)
```bash
./imgaug coco --parentDirs {parentDirPath} --categories '{category}' create

# {압축해제 폴더 이름}_coco 폴더 확인
# annotation.json 파일 확인
```

## 사용 예시
```bash
# 0. 데이터 수집
1. ~/datalake 폴더 생성 (루트)
2. 별@01-01-01-01#1 파일 다운로드
3. datalake 폴더에 압축 해제
4. 폴더명 비엔나코드로 변경
  비엔나코드 추출 방법
  [별@]01[-]01[-][01][-]01 => [] 안에 있는 문자 생략 => 010101

# 1. resize
5. ./imgaug resize --parentDirs ~/datalake/010101 --size 416 create

# 2. crop
6. ./imgaug crop --parentDirs ~/datalake/010101 --size 416 create

# 3. segment
7. ./imgaug segment --parentDirs ~/datalake/010101 --outputDir ~/datalake/segment --sample create

# 4. coco
8. ./imgaug coco --parentDirs ~/datalake/segment/train/010101 --categories '별' create
```