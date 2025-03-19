// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package widgets

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "edu-go-fiber/internal/vacancy"
import "edu-go-fiber/views/components"
import "fmt"

func VacancyList(vacancies []vacancy.Vacancy, pagesCount, page int) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = VacancyListStyle().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"vacancy-list__wrapper\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Title2("Последние вакансии", false).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<p>Найдите подходящую вакансию за пару минут!</p><div class=\"vacancy-list\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, vacancy := range vacancies {
			templ_7745c5c3_Err = components.VacancyCard(components.VacancyCardProps{
				Id:          vacancy.Id,
				Email:       vacancy.Email,
				Location:    vacancy.Location,
				Salary:      vacancy.Salary,
				CompanyType: vacancy.Type,
				Company:     vacancy.Company,
				Createdat:   vacancy.CreatedAt,
				Role:        vacancy.Role,
			}).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "</div><div class=\"vacancy-list__paginaton\"><div class=\"vacancy-list__paginaton-item\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if page != 1 {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<img src=\"/public/icons/chevron-left.svg\" alt=\"Стрелка влево\"> <a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 templ.SafeURL = templ.SafeURL(fmt.Sprintf("/?page=%d", page-1))
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var2)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "\">Предыдущие</a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "</div><div class=\"vacancy-list__paginaton-item\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if page != pagesCount {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "<a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 templ.SafeURL = templ.SafeURL(fmt.Sprintf("/?page=%d", page+1))
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var3)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "\">Следующие</a> <img src=\"/public/icons/chevron-right.svg\" alt=\"Стрелка справо\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "</div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func VacancyListStyle() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "<style>\n\t.vacancy-list__wrapper {\n\t\twidth: 100%;\n\t\tmax-width: 1296px;\n\t\tmargin: 0 auto;\n\t\tpadding: 60px 0;\n\t}\n\t.vacancy-list__wrapper p {\n\t\tmargin-bottom: 60px;\n\t}\n\t.vacancy-list {\n\t\tdisplay: flex;\n\t\talign-items: center;\n\t\tflex-direction: column;\n\t\tgap: 24px;\n\t\tmargin-bottom: 50px;\n\t}\n\t.vacancy-list__paginaton {\n\t\tdisplay: flex;\n\t\tjustify-content: space-between;\n\t\tgap: 24px;\n\t}\n\t.vacancy-list__paginaton-item {\n\t\tdisplay: flex;\n\t\tgap: 8px;\n\t\talign-items: center;\n\t}\n\t.vacancy-list__paginaton-item a {\n\t\ttext-decoration: none;\n\t\tcolor: #6C757D;\n\t}\n\t.vacancy-list__paginaton-item a:hover {\n\t\tcolor: #545a60;\n\t}\n</style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
