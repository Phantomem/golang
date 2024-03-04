package helpers

func ResolvePaginationFromArrayLength(length int, limit int, page int) (int, int) {
	var from int = page * limit
	var to int = page*limit + limit - 1
	if from > length {
		from = length
	}
	if to > length {
		to = length
	}
	return from, to
}
