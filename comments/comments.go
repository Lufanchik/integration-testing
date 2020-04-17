package comments

import (
	"lab.dt.multicarta.ru/tp/integration-testing/test"
)

var CasesCommentsCRUD = test.Cases{
	{
		N: "1. CommentsCRUD",
		T: test.T{
			&test.CommentsCRUD{},
		},
	},
}
