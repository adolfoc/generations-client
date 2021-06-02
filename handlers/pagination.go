package handlers

import (
	"fmt"
	"html/template"
	"strings"
)

const (
	RecordsPerPage = 20
)

type Pagination struct {
	TotalRecords  int
	NumberOfPages int
	BaseURL       string
	CurrentPage   int
}

func MakePagination(totalRecords int, baseURL string, page int) *Pagination {
	numberOfPages := (totalRecords / RecordsPerPage) + 1
	p := &Pagination{
		TotalRecords:  totalRecords,
		NumberOfPages: numberOfPages,
		BaseURL:       baseURL,
		CurrentPage:   page,
	}

	return p
}

func (p *Pagination) HTML(currentPage int) template.HTML {
	if p.TotalRecords <= RecordsPerPage {
		return ""
	}

	var navigation []string
	startPaginationDiv := fmt.Sprintf("<nav aria-label=%q>", "Navegación")
	navigation = append(navigation, startPaginationDiv)
	startPaginationUl := fmt.Sprintf("<ul class=%q>", "pagination")
	navigation = append(navigation, startPaginationUl)

	prevLi := p.makePreviousIl(currentPage)
	navigation = append(navigation, prevLi)

	for page := 1; page <= p.NumberOfPages; page++ {
		cssClass := "page-item"
		if page == currentPage {
			cssClass += " active"
		}
		url := fmt.Sprintf(p.BaseURL, page)
		pageLi := fmt.Sprintf("<li class=%q><a class=%q href=%q>%d</a></li>", cssClass, "page-link", url, page)
		navigation = append(navigation, pageLi)
	}

	lastLi := p.makeNextIl(currentPage)
	navigation = append(navigation, lastLi)

	endPaginationUl := fmt.Sprintf("</ul>")
	navigation = append(navigation, endPaginationUl)
	endPaginationDiv := fmt.Sprintf("</nav>")
	navigation = append(navigation, endPaginationDiv)

	return template.HTML(strings.Join(navigation, "\n"))
}

func (p *Pagination) makePreviousIl(currentPage int) string {
	cssClass := "page-item"
	if currentPage == 1 {
		cssClass += " disabled"
	}

	url := fmt.Sprintf(p.BaseURL, currentPage - 1)
	prevLi := fmt.Sprintf("<li class=%q><a class=%q href=%q>%s</a></li>", cssClass, "page-link", url, "Anterior")
	return prevLi
}

func (p *Pagination) makeNextIl(currentPage int) string {
	cssClass := "page-item"
	if currentPage >= p.NumberOfPages {
		cssClass += " disabled"
	}

	url := fmt.Sprintf(p.BaseURL, currentPage + 1)
	prevLi := fmt.Sprintf("<li class=%q><a class=%q href=%q>%s</a></li>", cssClass, "page-link", url, "Próxima")
	return prevLi
}

