package phone

import "github.com/gobuz/publicspam/internal/domain/model/phone"

func ExportSpamPhone(db SpamPhone) phone.SpamPhone {
	return phone.SpamPhone{
		ID:            db.ID.Hex(),
		PhoneNumber:   db.PhoneNumber,
		Note:          db.Note,
		Type:          db.Type,
		ReportedTimes: db.ReportedTimes,
		Reports:       ExportReports(db.Reports...),
	}
}
func ExportReport(db Report) phone.Report {
	return phone.Report{
		UserId:      db.UserId.Hex(),
		Note:        db.Note,
		Type:        db.Type,
		CreatedDate: db.CreatedDate,
	}
}
func ExportReports(dbs ...Report) []phone.Report {
	var result []phone.Report
	for _, db := range dbs {
		result = append(result, ExportReport(db))
	}
	return result
}

// FieldPhoneNumber
func FieldPhoneNumber() string {
	return "phoneNumber"
}
