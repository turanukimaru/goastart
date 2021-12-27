// Package goastart
/*
↑パッケージコメントとコメント本体は別にしたほうが良いみたい。多分タイトルとして認識されるのがそこだけなのだろう。
サブディレクトリに go.mod があるとサブモジュールになる。
上位の go.mod は配下のディレクトリをソースとして使うときに go.mod があるディレクトリは除外するようだ。
なので、特に意味が無ければサブモジュールを作る必要はない。
 ...じゃあサブモジュールの意味はなんだ？

なお、ディレクトリが gopath 配下にあると誤解を生みそう（他のディレクトリにパスが通るとか通らないとか）
なのでむしろ別のところのほうが良いみたい。
*/
package goastart
