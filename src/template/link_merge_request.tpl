{{ .Username }} {{ .State }} the merge request
合入方向: {{ .Source }} -> {{ .Target }}
last_commit_msg:
{{ .CommitNumber }}: fixed readme
状态: {{ .MergeStatus }}
标题: {{ .Title }}
项目: {{ .Project }}
代码审阅人：@{{ .Assignee }}
创建时间: {{ .CreatedAt }}
更新时间: {{ .UpdatedAt }}
链接: {{ .Link }}