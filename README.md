# cotton_stack

2024 情報システム特別講義 B 「Ruby インタプリタに見る実際のシステムソフトウェア」で書いた、Go で書かれたプリミティブなスタックマシン

授業で提示されたサンプル命令セットに従っています。

> 凡例: s0, s1, … は、スタックトップの最初の値、2 番目の値、… という意味
> push X: X をスタックトップに積む（push(X)）
> pop: スタックトップから 1 つ値を取り除く
> jump X: PC を X にする
> jumpif X: スタックトップが 0 でなければ PC を X にする
> add: push(s1 + s0)
> sub: push(s1 - s0)
> mul: push(s1 \* s0)
> set X: s0 をローカル変数 X に代入
> get X: push(ローカル変数 X)
> print: s0 を出力
> halt: 停止
