{{ .Username }} {{ .State }} the merge request
合入方向: {{ .Source }} -> {{ .Target }}
last_commit_msg:
{{ .CommitNumber }}: fixed readme
项目: {{ .Project }}
状态: {{ .MergeStatus }}
标题: {{ .Title }}
代码审阅人：@{{ .Assignee.Name }}(@{{ .Assignee.Username}})
创建时间: {{ .CreatedAt }}
更新时间: {{ .UpdatedAt }}
链接: {{ .Link }}