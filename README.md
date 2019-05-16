Project này có sử dụng vendor tool like Govendor. Thư viện này giống như npm ở javascript

`go get github.com/kardianos/govendor`

Sau khi clone source code về, chạy lệnh sau để govendor tải các repositories về:

`govendor sync`

Sau khi tải xong, chạy lệnh bên dưới để khởi động server:

`go run main.go`