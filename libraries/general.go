package libraries

import (
	"go-jlpt-n5/models"
	"math"
)

func BuildPagination(total, page, perPage, count int) *models.Pagination {
	pagination := new(models.Pagination)
	pagination.Total = total
	pagination.Count = count
	pagination.PerPage = perPage
	pagination.CurrentPage = page
	pagination.TotalPages = int(math.Ceil(float64(total) / float64(perPage)))

	return pagination
}
