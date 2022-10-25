package filters

type Result struct {
	Start   int64
	End     int64
	Taken   int64
	TxCount int64
	Txs     []string
}
