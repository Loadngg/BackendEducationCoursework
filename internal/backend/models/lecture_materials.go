package models

//go:generate reform

// LectureMaterials represents a row in lecture_materials table.
//
//reform:lecture_materials
type LectureMaterials struct {
	ID        int32  `reform:"id,pk" json:"id"`
	Title     string `reform:"title" json:"title"`
	File      []byte `reform:"file" json:"file"`
	LectureID *int32 `reform:"lecture_id" json:"lecture_id"`
}