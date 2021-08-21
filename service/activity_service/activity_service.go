package activity_service

import "xms/models"

//GetCurrentNotes return the latest 6 notes
func GetCurrentNotes() ([]models.Note,error) {
	notes,err:=models.GetNotesDesc(6)
	if err!=nil{
		return nil,err
	}
	return notes,nil
}
