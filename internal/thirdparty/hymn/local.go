package hymn

type (
	Local struct{}
)

func (l *Local) LookupSheet(dateStr string) (HymnSheet, error) {
	// This is a placeholder implementation.
	// In a real implementation, you would read from a local file or database.
	return HymnSheet{
		Opening:      "1",
		Sacrament:    "89",
		Intermediate: "Intermediate Hymn",
		Closing:      "300",
	}, nil
}
