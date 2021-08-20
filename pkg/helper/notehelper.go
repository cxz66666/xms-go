package helper

import (
	"fmt"
	"strconv"
	"xms/models"
)

//ToHumanReadableString convert models.Note to readable.
func ToHumanReadableString(note models.Note) string {
	switch note.Type {
	case models.Create:
		return "创建了工单"
	case models.Join:
		return "认领了工单"
	case models.ChangeState:
		// so dirty to read
		// rubbish go don't have enum
		status,err:= strconv.Atoi(note.Content)
		if err!=nil||status>models.Deleted.GetIndex()||status<models.Created.GetIndex(){
			return fmt.Sprintf("将工单的状态改变为%s",note.Content)
		}
		return fmt.Sprintf("将工单的状态改变为%s",models.TicketStatus(status).ToDisplayName())
	case models.Comment:
		return fmt.Sprintf("评论:%s",note.Content)
	case models.Edit:
		return fmt.Sprintf("编辑了工单。更改为:%s",note.Content)
	default:
		return ""
	}
}