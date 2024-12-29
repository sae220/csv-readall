# CSVのReadAllのベンチマーク比較
標準パッケージ[encoding/csv](https://pkg.go.dev/encoding/csv)の`(*Reader)ReadAll`がNamed return valuesを使って出力するスライスをゼロ値で初期化しており、Goでよく言われる「スライスは事前にメモリを確保したほうが速い」を考え、ファイルの行数を事前に読んでメモリを確保するということを試した\
CSVの規格を定めている[RFC4180](https://www.rfc-editor.org/rfc/rfc4180.html)によるとフィールドに改行を含んでも良いことになっているので、その場合は無駄にメモリ確保してしまうことになるが、あまりCSVで改行を含むデータをやりとりすることはないと感じたのでこうした

## 実際のベンチマーク
読みやすく加工しています
```
goos: linux
goarch: amd64
pkg: csv-readall
cpu: Intel(R) Core(TM) i5-10500H CPU @ 2.50GHz

size: 100
Std_ReadAll-12   1000000         1013 ns/op         4224 B/op          3 allocs/op
My_ReadAll-12     571328         2064 ns/op         8400 B/op          6 allocs/op

size: 1000
Std_ReadAll-12      6784       162606 ns/op       123816 B/op       1831 allocs/op
My_ReadAll-12       6511       175447 ns/op       119056 B/op       1827 allocs/op

size: 10000
Std_ReadAll-12       640      1856123 ns/op      1630262 B/op      19592 allocs/op
My_ReadAll-12        586      1982196 ns/op      1415329 B/op      19585 allocs/op

size: 100000
Std_ReadAll-12        49     22780813 ns/op     19603856 B/op     199957 allocs/op
My_ReadAll-12         57     20487635 ns/op     13228404 B/op     199943 allocs/op

size: 1000000
Std_ReadAll-12         5    211669824 ns/op    194235214 B/op    1999635 allocs/op
My_ReadAll-12          6    192830804 ns/op    157426528 B/op    1999614 allocs/op

size: 10000000
Std_ReadAll-12         1   2168241569 ns/op   1864869536 B/op   19999698 allocs/op
My_ReadAll-12          1   1893093999 ns/op   1440079384 B/op   19999669 allocs/op

PASS
ok      csv-readall     22.904s
```
