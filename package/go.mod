module package_test

go 1.18

// indirect: go.mod 에는 추가 되었으나 source code 상에서 사용하지 않는 module
// 처음 모듈을 get 해오면 무조건 붙는다.
// go mod tidy를 통해 안쓰는 모듈을 불러오지 않도록 정리

require github.com/google/uuid v1.3.0
