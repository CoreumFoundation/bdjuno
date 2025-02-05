package types

// DEXParamsRow represents a single row inside the dex_params table
type DEXParamsRow struct {
	OneRowID bool   `db:"one_row_id"`
	Params   string `db:"params"`
	Height   int64  `db:"height"`
}
